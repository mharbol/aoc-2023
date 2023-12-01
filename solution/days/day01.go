package days

import "fmt"

func Day01Part1(lines []string) (string, error) {

	acc := 0

	for _, line := range lines {
		first := firstDigit(line)
		last := lastDigit(line)
		acc += 10*first + last
	}

	return fmt.Sprint(acc), nil
}

func Day01Part2(lines []string) (string, error) {

	acc := 0

	for _, line := range lines {
		nums := detectWordsAndNums(line)
		acc += 10*nums[0] + nums[len(nums)-1]
	}

	return fmt.Sprint(acc), nil
}

func firstDigit(line string) int {
	for index := 0; index < len(line); index++ {
		digit := line[index]
		if digit >= '0' && digit <= '9' {
			return int(digit - '0')
		}
	}
	return -1
}

func lastDigit(line string) int {
	for index := len(line) - 1; index >= 0; index-- {
		digit := line[index]
		if digit >= '0' && digit <= '9' {
			return int(digit - '0')
		}
	}
	return -1
}

func detectWordsAndNums(line string) []int {

	var ret []int

	for index := 0; index < len(line); index++ {
		digit := line[index]
		if digit >= '0' && digit <= '9' {
			ret = append(ret, int(digit-'0'))
			continue
		}
		for num, str := range numArray {
			if substring(line, index, index+len(str)) == str {
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
