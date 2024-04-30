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

		sol := runDayX(key)

		expectedStruct, ok := allExpected[sol.day]
		if !ok {
			t.Fatalf("Error for day %d: No expected value in registry.", sol.day)
		}
		if sol.err1 != nil {
			t.Fatalf("Error in day %d part 1: %s", sol.day, sol.err1.Error())
		}
		if sol.err2 != nil {
			t.Fatalf("Error in day %d part 2: %s", sol.day, sol.err2.Error())
		}
		if sol.part1 != allExpected[sol.day].part1 {
			t.Fatalf("Mismatch in day %d. Expected %s, actual: %s", sol.day, sol.part1, expectedStruct.part1)
		}
		if sol.part2 != allExpected[sol.day].part2 {
			t.Fatalf("Mismatch in day %d. Expected %s, actual: %s", sol.day, sol.part2, expectedStruct.part2)
		}
	}
}

type testSolutionStruct struct {
	day          uint8
	part1, part2 string
	err1, err2   error
}

func runDayX(day uint8) *testSolutionStruct {
	p1Actual, e1 := Solve(day, 1)
	p2Actual, e2 := Solve(day, 2)
	return &testSolutionStruct{day: day, part1: p1Actual, part2: p2Actual, err1: e1, err2: e2}
}
