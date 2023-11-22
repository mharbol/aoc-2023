package solution

import (
	"os"
	"path"
	"runtime"
	"testing"
)

// Sets the filepath to the correct place for the purposes of these tests.
func init() {
    _, filename, _, _ := runtime.Caller(0)
    dir := path.Join(path.Dir(filename), "..")
    err := os.Chdir(dir)
    if err != nil {
        panic(err)
    }
}

func TestAllSolutions(t *testing.T) {

	for key := range allExpected {
		testPartX(key, 1, allExpected[key].part1, t)
		testPartX(key, 2, allExpected[key].part2, t)
	}
}

func testPartX(day, part uint8, expected string, t *testing.T) {
	var (
		actual string
		err    error
	)

	switch part {
	case 1:
		actual, err = Solve(day, 1)
	case 2:
		actual, err = Solve(day, 2)
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
