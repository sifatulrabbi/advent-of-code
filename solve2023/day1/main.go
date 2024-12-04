package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var numberWordMap = map[string]int{
	"one":       1,
	"two":       2,
	"three":     3,
	"four":      4,
	"five":      5,
	"six":       6,
	"seven":     7,
	"eight":     8,
	"nine":      9,
	"eleven":    11,
	"twelve":    12,
	"thirteen":  13,
	"fourteen":  14,
	"fifteen":   15,
	"sixteen":   16,
	"seventeen": 17,
	"eighteen":  18,
	"nineteen":  19,
	"twenty":    20,
	"thirty":    30,
	"forty":     40,
	"fifty":     50,
	"sixty":     60,
	"seventy":   70,
	"eighty":    80,
	"ninety":    90,
}

func main() {
	// f, err := os.ReadFile("./solve2023/inputs/input-day1.txt")
	f, err := os.ReadFile("./solve2023/inputs/input-day1-test.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	// part1(string(f))
	part2(string(f))
}

func part1(input string) {
	sums := []int{}

	lines := strings.Split(input, "\n")
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

	fmt.Println("part1:", total)
}

func part2(input string) {
	nums := []int{}

	for _, line := range strings.Split(input, "\n") {
		for i := 0; i < len(line); i++ {
			char := string(line[i])

			if isNum(char) {
				j := i + 1
				for j < len(line) && isNum(string(line[j])) {
					j++
				}
				n, _ := strconv.Atoi(line[i:j])
				nums = append(nums, n)
				i = j
			}

			if isNumChar(char) {
				j := i + 1
				for j < len(line) && isNumChar(string(line[j])) {
					if isNumWord(string(line[i:j])) {
						break
					}
					j++
				}
				nums = append(nums, getNumFromWord(line[i:j]))
				i = j
			}

		}
	}
	// fmt.Println(nums)

	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println("part2:", total)
}

func isNum(v string) bool {
	return slices.Contains([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, v)
}

func isNumWord(v string) bool {
	return slices.Contains([]string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		// "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
		// "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
	}, v)
}

func isNumChar(v string) bool {
	return slices.Contains([]string{
		"o", "n", "e", "t", "w", "h", "r", "f", "u", "i", "v", "s", "g",
	}, v)
}

func getNumFromWord(w string) int {
	v, ok := numberWordMap[w]
	if !ok {
		return 0
	}
	return v
}
