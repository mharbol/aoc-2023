package day15

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) string {

	entries := splitLineOnCommas(lines[0])

	tot := 0

	for _, str := range entries {
		tot += reindeerHash(str)
	}

	return fmt.Sprint(tot)
}

func Part2(lines []string) string {

	entries := splitLineOnCommas(lines[0])

	var boxes []*lensBox
	for i := 0; i < 256; i++ {
		boxes = append(boxes, newLensBox(i+1))
	}

	for _, entry := range entries {
		if strings.Contains(entry, "-") {
			label := entry[:len(entry)-1]
			boxNumber := reindeerHash(label)
			boxes[boxNumber].remove(label)
		} else {
			currLens := newLensFromString(entry)
			boxes[currLens.boxNumber()].replaceOrInsert(currLens)
		}
	}

	tot := 0

	for _, box := range boxes {
		tot += box.totalBox()
	}

	return fmt.Sprint(tot)
}

func reindeerHash(str string) int {
	hash := 0
	for _, char := range str {
		hash += int(char)
		hash *= 17
		hash %= 256
	}
	return hash
}

func splitLineOnCommas(line string) []string {
	return strings.Split(line, ",")
}

// struct to represent distinct lenses
type lens struct {
	label    string
	focalLen int
}

func newLens(label string, focalLen int) *lens {
	return &lens{label: label, focalLen: focalLen}
}

// for strings with equals
func newLensFromString(str string) *lens {
	strs := strings.Split(str, "=")
	focalLen, _ := strconv.Atoi(strs[1])
	return newLens(strs[0], focalLen)
}

func (l *lens) String() string {
	return fmt.Sprintf("[%s %d]", l.label, l.focalLen)
}

func (l *lens) boxNumber() int {
	return reindeerHash(l.label)
}

// struct to represent the boxes of lenses and facilitate fast insertions, removals, and replacements
type lensBox struct {

	// the key is the soft "index", this is max val when it was inserted and a way to keep order when they are manipulted
	lenses map[int]*lens

	// quick reference to which labels are in the lensBox and what their position (soft index) is
	lables map[string]int

	// a keeper for the next available "index"
	// at the end the boxes will be condensed so there is no need to worry about empty space
	nextIndex int

	// the value of the box when summing totals
	boxVal int
}

func (lb *lensBox) String() string {
	return fmt.Sprint(lb.asSlice())
}

// removes a lens via its label from this lensBox (provided it exists)
func (lb *lensBox) remove(label string) {
	idx, ok := lb.lables[label]

	// if exists, delete it
	if ok {
		delete(lb.lables, label)
		delete(lb.lenses, idx)
	}
	// otherwise do nothing
}

// replaces a lens (lensPtr) in this lensBox if already present or inserts lensPtr on the end if none with its label are present
func (lb *lensBox) replaceOrInsert(lensPtr *lens) {
	idx, ok := lb.lables[lensPtr.label]

	// if exists, replace
	if ok {
		lb.lenses[idx] = lensPtr
	} else { // otherwise tack on end
		lb.lables[lensPtr.label] = lb.nextIndex
		lb.lenses[lb.nextIndex] = lensPtr
		lb.nextIndex++
	}
}

// not an optimal way to iterate over the lensBox I am certain, but it comes after the computation of finding
// the final state of insertions so it's probably alright
func (lb *lensBox) asSlice() []*lens {
	var arr []*lens
	for i := 0; i < lb.nextIndex; i++ {
		if lensPtr, ok := lb.lenses[i]; ok {
			arr = append(arr, lensPtr)
		}
	}
	return arr
}

// returns the rule-defined internal value of this lensBox
func (lb *lensBox) getInternalValue() int {
	prod := 1

	for idx, lensPrt := range lb.asSlice() {
		prod *= (idx + 1) * lensPrt.focalLen
	}

	return prod
}

func (lb *lensBox) totalBox() int {
	tot := 0
	for idx, currLens := range lb.asSlice() {
		tot += lb.boxVal * (idx + 1) * currLens.focalLen
	}
	return tot
}

func newLensBox(boxVal int) *lensBox {
	return &lensBox{lenses: make(map[int]*lens), lables: make(map[string]int), nextIndex: 0, boxVal: boxVal}
}
