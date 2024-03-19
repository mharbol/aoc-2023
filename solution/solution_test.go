package solution

import (
	"os"
	"path"
	"runtime"
	"sort"
	"sync"
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

	var wg sync.WaitGroup
	dayChan := make(chan *testSolutionStruct)

	for key := range allExpected {
		wg.Add(1)
		go runDayX(key, dayChan, &wg)
	}

	go func() {
		defer close(dayChan)
		wg.Wait()
	}()

	dayNums := make([]uint8, 0)
	solutions := make(map[uint8]*testSolutionStruct)
	for sol := range dayChan {
		dayNums = append(dayNums, sol.day)
		solutions[sol.day] = sol
	}

	sort.Slice(dayNums, func(i, j int) bool {
		return dayNums[i] < dayNums[j]
	})

	for _, day := range dayNums {
		expectedStruct, ok := allExpected[day]
		if !ok {
			t.Fatalf("Error for day %d: No expected value in registry.", day)
		}
		sol := solutions[day]
		if sol.err1 != nil {
			t.Fatalf("Error in day %d part 1: %s", day, sol.err1.Error())
		}
		if sol.err2 != nil {
			t.Fatalf("Error in day %d part 2: %s", day, sol.err2.Error())
		}
		if solutions[day].part1 != allExpected[day].part1 {
			t.Fatalf("Mismatch in day %d. Expected %s, actual: %s", day, solutions[day].part1, expectedStruct.part1)
		}
		if solutions[day].part2 != allExpected[day].part2 {
			t.Fatalf("Mismatch in day %d. Expected %s, actual: %s", day, solutions[day].part2, expectedStruct.part2)
		}
	}
}

type testSolutionStruct struct {
	day          uint8
	part1, part2 string
	err1, err2   error
}

func runDayX(day uint8, ch chan<- *testSolutionStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	p1Actual, e1 := Solve(day, 1)
	p2Actual, e2 := Solve(day, 2)
	ch <- &testSolutionStruct{day: day, part1: p1Actual, part2: p2Actual, err1: e1, err2: e2}
}
