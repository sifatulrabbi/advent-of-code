package solve2023

import (
	"regexp"
	"strconv"
	"strings"
)

var numberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func Day1Part2(text string) int {
	sums := []int{}
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		re := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|[0-9]`)
		matches := re.FindAllString(line, -1)

		if len(matches) < 1 {
			continue
		}

		sum := 10*numberMap[matches[0]] + numberMap[matches[len(matches)-1]]
		sums = append(sums, sum)
	}

	total := 0
	for _, sum := range sums {
		total += sum
	}
	return total
}

func Day1(text string) int {
	sums := []int{}

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		nl := []string{}

		for _, r := range strings.Split(line, "") {
			if isNum(r) {
				nl = append(nl, r)
			}
		}

		if len(nl) < 1 {
			continue
		}

		strSum := ""
		if len(nl) < 2 {
			strSum += nl[0]
			strSum += nl[0]
		} else {
			strSum += nl[0]
			strSum += nl[len(nl)-1]
		}

		sum, _ := strconv.ParseInt(strSum, 10, 32)
		sums = append(sums, int(sum))
	}

	total := 0
	for _, s := range sums {
		total += s
	}
	return total
}

func isNum(v string) bool {
	nums := strings.Split("1234567890", "")
	for _, n := range nums {
		if n == v {
			return true
		}
	}
	return false
}
