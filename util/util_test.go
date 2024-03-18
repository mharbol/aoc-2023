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

func TestReadDayInputFailure(t *testing.T) {
	const path = "inputs/day_50.txt"

	_, err := readDayInput(path)

	if err == nil {
		t.Fatalf("Expected funcion to error and did not.")
	}
}
