package util

import "testing"

func TestMakeFilePathSingleDigit(t *testing.T) {
	const day uint8 = 1

	expected := "inputs/day_01.txt"
	actual := makeFilePath(day)

	if expected != actual {
		t.Fatalf("For day %d, expected: %s, returned: %s", day, expected, actual)
	}
}

func TestMakeFilePathDoubleDigit(t *testing.T) {
	const day uint8 = 23

	expected := "inputs/day_23.txt"
	actual := makeFilePath(day)

	if expected != actual {
		t.Fatalf("For day %d, expected: %s, returned: %s", day, expected, actual)
	}
}

func TestReadDayInputSuccess(t *testing.T) {
	const path = "../inputs/day_00.txt"
	expected := []string{"line 1", "line 2", "line 3", ""}

	actual, err := readDayInput(path)

	if err != nil {
		t.Fatalf("Got error %s", err.Error())
	}
	if expected[0] != actual[0] && expected[1] != actual[1] && expected[2] != actual[2] && expected[3] != actual[3] {
		t.Fatalf("Expected did not match actual")
	}
}

func TestReadDayInputFailure(t *testing.T) {
	const path = "inputs/day_50.txt"

	_, err := readDayInput(path)

	if err == nil {
		t.Fatalf("Expected funcion to error and did not.")
	}
}

func TestReadProblemInfoNewEntry(t *testing.T) {

}

func TestReadProblemInfoExistingEntry(t *testing.T) {

}

func TestReadProblemNoEntry(t *testing.T) {

}
