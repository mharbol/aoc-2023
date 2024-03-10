package days

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day19Part1(lines []string) (string, error) {

	parts, workflows := parseMachinePartsAndWorkflows(lines)

	var accepted []*machinePart

	for _, part := range parts {
		currFlow := "in"
		for currFlow != "A" && currFlow != "R" {
			currFlow = workflows[currFlow].execute(part)
		}
		if currFlow == "A" {
			accepted = append(accepted, part)
		}
	}

	tot := 0
	for _, part := range accepted {
		tot += part.total()
	}

	return fmt.Sprint(tot), nil
}

func Day19Part2(lines []string) (string, error) {

	const (
		START int = 1
		END   int = 4000
	)

	_, workflows := parseMachinePartsAndWorkflows(lines)

	var tot int64 = 0

	var pairs []*workflowPartPair
	var accepted []*spanningMachinePart

	splittingPart := &spanningMachinePart{&incSpan{START, END}, &incSpan{START, END}, &incSpan{START, END}, &incSpan{START, END}}
	pairs = append(pairs, &workflowPartPair{nextWorkflow: "in", part: splittingPart})

	for len(pairs) > 0 {
		var nextPairs []*workflowPartPair
		for _, pair := range evalAllPartPairs(workflows, pairs) {
			if pair.nextWorkflow == "A" {
				accepted = append(accepted, pair.part)
			} else if pair.nextWorkflow == "R" {
				continue
			} else {
				nextPairs = append(nextPairs, pair)
			}
		}
		pairs = nextPairs
	}

	for _, acc := range accepted {
		tot += int64(acc.totalParts())
	}

	return fmt.Sprint(tot), nil
}

func parseMachinePartsAndWorkflows(lines []string) ([]*machinePart, map[string]*workflow) {
	split := slices.Index(lines, "")
	var parts []*machinePart
	workflows := make(map[string]*workflow)

	for idx := 0; idx < split; idx++ {
		flow := newWorkflowFromString(lines[idx])
		workflows[flow.name] = flow
	}

	for idx := split + 1; idx < len(lines); idx++ {
		parts = append(parts, newMachinePartFromString(lines[idx]))
	}

	return parts, workflows
}

type machinePart struct {
	x, m, a, s int
}

func newMachinePartFromString(line string) *machinePart {
	values := strings.Split(line[1:len(line)-1], ",")
	x, _ := strconv.Atoi(values[0][2:])
	m, _ := strconv.Atoi(values[1][2:])
	a, _ := strconv.Atoi(values[2][2:])
	s, _ := strconv.Atoi(values[3][2:])
	return &machinePart{x, m, a, s}
}

func (mp *machinePart) total() int {
	return mp.x + mp.m + mp.a + mp.s
}

type workflow struct {
	name string
	jobs []string
}

func newWorkflowFromString(line string) *workflow {
	flow := &workflow{}
	idxOpenBrace := strings.Index(line, "{")

	flow.name = line[:idxOpenBrace]
	flow.jobs = strings.Split(line[idxOpenBrace+1:len(line)-1], ",")

	return flow
}

// executes the workflow on the given machinePart
// returns the next workflow to execute if needed, otherwise
// "R" for reject or "A" for accept
func (w *workflow) execute(part *machinePart) string {

	for _, job := range w.jobs {
		colIdx := strings.Index(job, ":")
		// if this is an auto goto
		if colIdx == -1 {
			return job
		}

		// analyze and do the job
		// get the correct variable
		letter := 0
		switch job[0] {
		case 'x':
			letter = part.x
			break
		case 'm':
			letter = part.m
			break
		case 'a':
			letter = part.a
			break
		case 's':
			letter = part.s
			break
		}

		comparison, _ := strconv.Atoi(job[2:colIdx])

		// do everything as greater than so flip the sign if less than
		if job[1] == '<' {
			letter *= -1
			comparison *= -1
		}

		// do the comparison
		if letter > comparison {
			return job[colIdx+1:]
		}
	}

	return ""
}

func (w *workflow) evaluateOnSpannedPart(part *spanningMachinePart) []*workflowPartPair {
	following := part.copy()
	var matching *spanningMachinePart
	var pairs []*workflowPartPair
	for _, job := range w.jobs {
		matching, following = following.splitOnJob(job)
		if matching != nil {
			pairs = append(pairs, &workflowPartPair{nextWorkflow: getNextWorkflow(job), part: matching.copy()})
		}
		if following == nil {
			break
		}
	}
	return pairs
}

// wrapper for a destined workflow and the class of part which will enter said workflow
type workflowPartPair struct {
	nextWorkflow string               // name of the workflow this pairing is destined for
	part         *spanningMachinePart // the spanningMachinePart to enter the workflow
}

func (p workflowPartPair) String() string {
	return fmt.Sprintf("{next: %v, part %v}", p.nextWorkflow, p.part)
}

// "inclusive span" [start, end]
type incSpan struct {
	start, end int
}

func (is *incSpan) copy() *incSpan {
	return &incSpan{is.start, is.end}
}

func (is *incSpan) span() int {
	return is.end - is.start + 1
}

// a machine part that represents the ranges of values a part can hold and pass through
type spanningMachinePart struct {
	x, m, a, s *incSpan
}

func (smp *spanningMachinePart) totalParts() int {
	return smp.x.span() * smp.m.span() * smp.a.span() * smp.s.span()
}

func (smp *spanningMachinePart) copy() *spanningMachinePart {
	return &spanningMachinePart{x: smp.x.copy(), m: smp.m.copy(), a: smp.a.copy(), s: smp.s.copy()}
}

func (smp spanningMachinePart) String() string {
	return fmt.Sprintf("{x: %v, m: %v, a: %v, s: %v}", smp.x, smp.m, smp.a, smp.s)
}

// first return is the part that matches the condition
// second return is the part that keeps going along the workflow
func (smp *spanningMachinePart) splitOnJob(job string) (matches, follows *spanningMachinePart) {

	// if is a default job, everything matches
	if !isComparisonJob(job) {
		return smp.copy(), nil
	}

	var matchingPart *spanningMachinePart = nil
	var followingPart *spanningMachinePart = nil

	var letter *incSpan

	switch job[0] {
	case 'x':
		letter = smp.x
		break
	case 'm':
		letter = smp.m
		break
	case 'a':
		letter = smp.a
		break
	case 's':
		letter = smp.s
		break
	}

	colIdx := strings.Index(job, ":")

	comparisonSign := job[1]
	comparisonNumber, _ := strconv.Atoi(job[2:colIdx])

	if comparisonSign == '>' {
		// if everything is greater
		if letter.start > comparisonNumber {
			matchingPart = smp.copy()
		} else if letter.end <= comparisonNumber {
			// if everything is less
			followingPart = smp.copy()
		} else {
			// split on comparisonNumber
			matchingPart = smp.copy()
			followingPart = smp.copy()
			switch job[0] {
			case 'x':
				matchingPart.x = &incSpan{start: comparisonNumber + 1, end: letter.end}
				followingPart.x = &incSpan{start: letter.start, end: comparisonNumber}
				break
			case 'm':
				matchingPart.m = &incSpan{start: comparisonNumber + 1, end: letter.end}
				followingPart.m = &incSpan{start: letter.start, end: comparisonNumber}
				break
			case 'a':
				matchingPart.a = &incSpan{start: comparisonNumber + 1, end: letter.end}
				followingPart.a = &incSpan{start: letter.start, end: comparisonNumber}
				break
			case 's':
				matchingPart.s = &incSpan{start: comparisonNumber + 1, end: letter.end}
				followingPart.s = &incSpan{start: letter.start, end: comparisonNumber}
				break
			}
		}
	} else { // comparisonSign == '<'
		// if everything is less than
		if letter.end < comparisonNumber {
			matchingPart = smp.copy()
		} else if letter.start >= comparisonNumber {
			// if everything is greater
			followingPart = smp.copy()
		} else {
			// split on comparisonNumber
			matchingPart = smp.copy()
			followingPart = smp.copy()
			switch job[0] {
			case 'x':
				matchingPart.x = &incSpan{start: letter.start, end: comparisonNumber - 1}
				followingPart.x = &incSpan{start: comparisonNumber, end: letter.end}
				break
			case 'm':
				matchingPart.m = &incSpan{start: letter.start, end: comparisonNumber - 1}
				followingPart.m = &incSpan{start: comparisonNumber, end: letter.end}
				break
			case 'a':
				matchingPart.a = &incSpan{start: letter.start, end: comparisonNumber - 1}
				followingPart.a = &incSpan{start: comparisonNumber, end: letter.end}
				break
			case 's':
				matchingPart.s = &incSpan{start: letter.start, end: comparisonNumber - 1}
				followingPart.s = &incSpan{start: comparisonNumber, end: letter.end}
				break
			}
		}
	}

	return matchingPart, followingPart
}

// returns whether or not the given job string is a comparison job
// depends on the existance of a ':' character for this determinaion
func isComparisonJob(job string) bool {
	return strings.Index(job, ":") != -1
}

func getNextWorkflow(job string) string {
	return job[strings.Index(job, ":")+1:]
}

func evalAllPartPairs(workflows map[string]*workflow, pairs []*workflowPartPair) []*workflowPartPair {
	var newPairs []*workflowPartPair

	for _, pair := range pairs {
		newPairs = append(newPairs, workflows[pair.nextWorkflow].evaluateOnSpannedPart(pair.part)...)
	}

	return newPairs
}
