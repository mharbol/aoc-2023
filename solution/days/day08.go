package days

import "fmt"

func Day08Part1(lines []string) (string, error) {

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

	return fmt.Sprint(count), nil
}

func Day08Part2(lines []string) (string, error) {

	return "", nil
}

type lrPair struct {
	left  string
	right string
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
