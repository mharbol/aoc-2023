package span

import "sort"

// represents a range of numbers from `start` to `end` (non-inclusively)
type Span struct {
	start int
	end   int
}

// Create a new *Span starting at `start` and ending at `end` non-inclusively.
// Returns `nil` if start >= end
func New(start, end int) *Span {
	if start >= end {
		return nil
	}
	return &Span{
		start: start,
		end:   end,
	}
}

// Effectively makes a Venn Diagram of the two spans.
// The first return value is what the two Spans have in common, nil if none.
// The second return value is what is exclusive to the calling Span, empty slice if none.
// The third return value is what is exclusive to the other Span, empty slice if none.
func (s *Span) Venn(o *Span) (*Span, []*Span, []*Span) {
	// nil check
	if nil == o {
		return nil, []*Span{New(s.start, s.end)}, []*Span{}
	}

	var common *Span = nil
	var self []*Span
	var other []*Span

	// work it left to right from other's perspective
	// starting at start, look for strict less than, followed by equal, then else is strictly greater than
	// follow the roughly the same logic for end
	if o.start < s.start {
		if o.end <= s.start {
			self = append(self, New(s.start, s.end))
			other = append(other, New(o.start, o.end))
		} else if o.end < s.end {
			self = append(self, New(o.end, s.end))
			other = append(other, New(o.start, s.start))
			common = New(s.start, o.end)
		} else if o.end == s.end {
			common = New(s.start, s.end)
			other = append(other, New(o.start, s.start))
		} else {
			common = New(s.start, s.end)
			other = append(other, New(o.start, s.start), New(s.end, o.end))
		}
	} else if o.start == s.start {
		if o.end < s.end {
			common = New(s.start, o.end)
			self = append(self, New(o.end, s.end))
		} else if o.end == s.end {
			common = New(s.start, s.end)
		} else {
			common = New(s.start, s.end)
			other = append(self, New(s.end, o.end))
		}
	} else if o.start < s.end {
		if o.end < s.end {
			common = New(o.start, o.end)
			self = append(self, New(s.start, o.start), New(o.end, s.end))
		} else if o.end == s.end {
			common = New(o.start, o.end)
			self = append(self, New(s.start, o.start))
		} else {
			common = New(o.start, s.end)
			self = append(self, New(s.start, o.start))
			other = append(other, New(s.end, o.end))
		}
	} else {
		self = append(self, New(s.start, s.end))
		other = append(other, New(o.start, o.end))
	}
	return common, self, other
}

// used to sort Span structs
// first compares based on starting value
// if starting value is the same then go off of ending value
func (s *Span) CompareTo(o *Span) int {
	if s.start != o.start {
		return s.start - o.start
	}
	return s.end - o.end
}

func (s *Span) IsIntInSpan(num int) bool {
	return num >= s.start && num < s.end
}

func (s *Span) Start() int {
	return s.start
}

func (s *Span) End() int {
	return s.end
}

func Combine(in []*Span) []*Span {

	var out []*Span

	if in == nil || len(in) == 0 {
		return out
	}

	sort.Slice(in, func(i, j int) bool { return in[i].CompareTo(in[j]) < 0 })

	var workingSpan = in[0]

	for idx := 1; idx < len(in); idx++ {
		if workingSpan.end >= in[idx].start {
			if in[idx].end > workingSpan.end {
				workingSpan = New(workingSpan.start, in[idx].end)
			}
		} else {
			out = append(out, workingSpan)
			workingSpan = in[idx]
		}
	}

	out = append(out, workingSpan)
	return out
}
