package day08

import "fmt"

func Part1(lines []string) string {

	instructions := lines[0]
	length := len(instructions)
	digraph := makeDigraph(lines[2:])
	currNode := "AAA"
	count := 0

	for currNode != "ZZZ" {
		instruction := instructions[count%length]
		count++
		if instruction == 'L' {
			currNode = digraph[currNode].left
		} else {
			currNode = digraph[currNode].right
		}
	}

	return fmt.Sprint(count)
}

func Part2(lines []string) string {

	instructions := lines[0]
	length := len(instructions)
	digraph := makeDigraph(lines[2:])

	aNodes := allNodesEndingInA(&digraph)
	var trackers []*nodeDepthTracker

	for _, node := range aNodes {
		trackers = append(trackers, &nodeDepthTracker{node: node, depth: 0})
	}

	// this was all I needed to do to find the cycles
	for _, tracker := range trackers {
		tracker.getNextZ(digraph, instructions, length)
	}

	prod := 1
	var commonDiv int = trackers[0].depth

	// find common divisor
	for i := 0; i < len(trackers); i++ {
		commonDiv = gcd(commonDiv, trackers[i].depth)
	}

	for _, tracker := range trackers {
		prod *= (tracker.depth) / commonDiv
	}

	return fmt.Sprint(prod * commonDiv)
}

type lrPair struct {
	left  string
	right string
}

// keeps track of how far down the node is in the
// search
// this ended up being overkill
type nodeDepthTracker struct {
	node  string // node name
	depth int    // step at which we arrived on this node
}

// updates the depth tracker to the next Z node
func (ndt *nodeDepthTracker) getNextZ(digraph map[string]lrPair, instructions string, length int) {

	// progress node by one
	var nextNode string
	instruction := instructions[ndt.depth%length]
	depth := ndt.depth + 1
	if instruction == 'L' {
		nextNode = digraph[ndt.node].left
	} else {
		nextNode = digraph[ndt.node].right
	}

	// loop until we find another Z
	for nextNode[2] != 'Z' {
		instruction = instructions[depth%length]
		depth++
		if instruction == 'L' {
			nextNode = digraph[nextNode].left
		} else {
			nextNode = digraph[nextNode].right
		}
	}

	ndt.depth = depth
	ndt.node = nextNode

}

func makeDigraph(lines []string) map[string]lrPair {

	digraph := make(map[string]lrPair)

	for _, line := range lines {
		key := line[:3]
		pair := lrPair{
			left:  line[7:10],
			right: line[12:15],
		}
		digraph[key] = pair
	}

	return digraph
}

func allNodesEndingInA(graph *map[string]lrPair) []string {

	var out []string

	for node := range *graph {
		if node[2] == 'A' {
			out = append(out, node)
		}
	}
	return out
}

// Euclidian gcd function
func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
