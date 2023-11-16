package solution

import "testing"

func TestAllSolutions(t *testing.T) {

	for key := range allExpected {
		testPartX(key, 1, allExpected[key].part1, t)
		testPartX(key, 2, allExpected[key].part2, t)
	}
}

func testPartX(day, part uint8, expected string, t *testing.T) {
	s, ok := allSolutions[day]
	if !ok {
		t.Fatalf("Solution for day %d not found", day)
	}

	var actual string
	var err error

	switch part {
	case 1:
		actual, err = Solve(s, 1)
	case 2:
		actual, err = Solve(s, 2)
    default:
		t.Fatalf("Part must be 1 or 2.")
	}

	if err != nil {
		t.Fatalf("Error in day %d part %d: %s", day, part, err.Error())
	}

	if expected != actual {
		t.Fatalf("Fail for day %d part %d:\nExpected: %s\nActual: %s", day, part, expected, actual)
	}
}
