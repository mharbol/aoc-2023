package day06

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Part1(lines []string) string {

	times, distances := parseDay6Part1(lines)

	acc := 1

	for idx, time := range times {
		count := 0
		dist := distances[idx]
		for x := 0; x <= time; x++ {
			if raceResult(time, x) > dist {
				count++
			}
		}
		acc *= count
	}

	return fmt.Sprint(acc)
}

func Part2(lines []string) string {

	time, distance := parseDay6Part2(lines)

	// quadratic formula
	var root float64 = math.Sqrt(float64(time*time - 4*distance))
	left := math.Ceil((float64(time) - root) / 2)
	right := math.Floor((float64(time) + root) / 2)

	return fmt.Sprint(int(right - left + 1))
}

func raceResult(timeTotal, timeTaken int) int {
	return (timeTotal - timeTaken) * timeTaken // timeTaken is also the speed
}

func parseDay6Part1(lines []string) ([]int, []int) {

	var times, distances []int

	re := regexp.MustCompile(" +")
	timesStrArr := re.Split(strings.TrimSpace(lines[0][6:]), -1)
	distStrArr := re.Split(strings.TrimSpace(lines[1][9:]), -1)

	for _, str := range timesStrArr {
		num, _ := strconv.Atoi(str)
		times = append(times, num)
	}
	for _, str := range distStrArr {
		num, _ := strconv.Atoi(str)
		distances = append(distances, num)
	}

	return times, distances
}

func parseDay6Part2(lines []string) (int, int) {

	timeStr := strings.ReplaceAll(lines[0][6:], " ", "")
	distStr := strings.ReplaceAll(lines[1][9:], " ", "")

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distStr)

	return time, distance
}
