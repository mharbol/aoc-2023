package solution

type Solution interface {
	Day() uint8
	Part1() (string, error)
	Part2() (string, error)
}
