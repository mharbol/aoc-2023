package day20

// import "fmt"

func Part1(lines []string) string {

	// lowCount, highCount := 0, 0

	return ""
}

func Part2(lines []string) string {

	return ""
}

type module interface {
	getName() string
	moduleList() []*string
	receive(p pulse, lowCount *int, highCount *int) pulse
}

type state bool

const (
	ON  state = true
	OFF state = false
)

type pulse int

const (
	HIGH_PULSE pulse = 0
	LOW_PULSE  pulse = 1
	NO_PULSE   pulse = 2
)
