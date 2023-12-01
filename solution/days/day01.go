package days

import "fmt"

func Day01Part1(lines []string) (string, error) {

	acc := 0

	for _, line := range lines {
		if line == "" {
			break
		}
		first := firstDigit(line)
		last := lastDigit(line)
		acc += 10*first + last
	}

	return fmt.Sprint(acc), nil
}

func Day01Part2(lines []string) (string, error) {

	acc := 0

	for _, line := range lines {
		if line == "" {
			break
		}
		nums := detectWordsAndNums(line)
		acc += 10*nums[0] + nums[len(nums)-1]
	}

	return fmt.Sprint(acc), nil
}

func firstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		d := line[i]
		if d >= '0' && d <= '9' {
			return int(d - '0')
		}
	}
	return -1
}

func lastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		d := line[i]
		if d >= '0' && d <= '9' {
			return int(d - '0')
		}
	}
	return -1
}

func detectWordsAndNums(line string) []int {

	var ret []int

	for i := 0; i < len(line); i++ {
		d := line[i]
		if d >= '0' && d <= '9' {
			ret = append(ret, int(d-'0'))
			continue
		}
		for num, str := range numArray {
			if substring(line, i, i+len(str)) == str {
				ret = append(ret, num)
			}
		}
	}
	return ret
}

var numArray = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func substring(str string, start, end int) string {
	if end > len(str) {
		return ""
	}
	return str[start:end]
}
