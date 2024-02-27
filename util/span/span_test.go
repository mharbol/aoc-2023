package span

import (
	"testing"
)

type testCase struct {
	self           *Span
	other          *Span
	caseName       string
	expectedCommon *Span
	expectedSelf   []*Span
	expectedOther  []*Span
}

var baseSpan *Span = New(3, 14)
var emptySlice []*Span

func TestNew(t *testing.T) {
	span := &Span{start: 3, end: 14}

	if *span != *baseSpan {
		t.Fatal("Spans do not equal")
	}
}

var cases []*testCase = []*testCase{
    {baseSpan, New(3, 14), "same span (1)", New(3, 14), emptySlice, emptySlice},
	{baseSpan, New(0, 2), "totally out to the left (2)", nil, []*Span{New(3, 14)}, []*Span{New(0, 2)}},
	{baseSpan, New(1, 3), "barely out to the left (3)", nil, []*Span{New(3, 14)}, []*Span{New(1, 3)}},
    {baseSpan, New(15, 18), "totally off to right (4)", nil, []*Span{New(3, 14)}, []*Span{New(15, 18)}},
    {baseSpan, New(14, 17), "barely touches left and off to rigth (5)", nil, []*Span{New(3, 14)}, []*Span{New(14, 17)}},
    {baseSpan, New(5, 11), "totally covered by self (6)", New(5, 11), []*Span{New(3, 5), New(11, 14)}, emptySlice},
    {baseSpan, New(3, 10), "touching left and fully inside right (7)", New(3, 10), []*Span{New(10, 14)}, emptySlice},
    {baseSpan, New(8, 14), "totally covered by self and touches right (8)", New(8, 14), []*Span{New(3, 8)}, emptySlice},
	{baseSpan, New(2, 7), "out to the left and fully inside right (9)", New(3, 7), []*Span{New(7, 14)}, []*Span{New(2, 3)}},
    {baseSpan, New(11, 18), "totally covered by self and off to right (10)", New(11, 14), []*Span{New(3, 11)}, []*Span{New(14, 18)}},
	{baseSpan, New(0, 19), "totally covers self (11)", New(3, 14), emptySlice, []*Span{New(0, 3), New(14, 19)}},
	{baseSpan, New(1, 14), "totally covers self and barely touches right (12)", New(3, 14), emptySlice, []*Span{New(1, 3)}},
	{baseSpan, New(3, 16), "totally covered by self and off to right (13)", New(3, 14), emptySlice, []*Span{New(14, 16)}},
}

func TestVenn(t *testing.T) {
	for _, test := range cases {
		testVennUtil(test, t)
	}
}

func testVennUtil(tc *testCase, t *testing.T) {

	actualCommon, actualSelf, actualOther := tc.self.Venn(tc.other)

	if !spanEqualityCheck(actualCommon, tc.expectedCommon) {
		t.Fatalf("Failed case %s in common check. Expected %v got %v", tc.caseName, tc.expectedCommon, actualCommon)
	}

	if !spanSliceEqualityCheck(actualSelf, tc.expectedSelf) {
		t.Fatalf("Failed case %s in self check. Expected %v got %v", tc.caseName, tc.expectedSelf, actualSelf)
	}

	if !spanSliceEqualityCheck(actualOther, tc.expectedOther) {
		t.Fatalf("Failed case %s in other check. Expected %v got %v", tc.caseName, tc.expectedOther, actualOther)
	}
}

func spanEqualityCheck(s0, s1 *Span) bool {
	if s0 == nil && s1 != nil || s0 != nil && s1 == nil {
		return false
	}
	if s0 == nil && s1 == nil {
		return true
	}
	return *s0 == *s1
}

func spanSliceEqualityCheck(sl0, sl1 []*Span) bool {
	if sl0 == nil && sl1 != nil || sl0 != nil && sl1 == nil {
		return false
	}
	if sl0 == nil && sl1 == nil {
		return true
	}
	if len(sl0) != len(sl1) {
		return false
	}
	for idx := 0; idx < len(sl0); idx++ {
		if *sl0[idx] != *sl1[idx] {
			return false
		}
	}
	return true
}
