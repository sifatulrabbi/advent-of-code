package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// f, err := os.ReadFile("./solve2024/inputs/input-day2-test.txt")
	f, err := os.ReadFile("./solve2024/inputs/input-day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(f)
	reports := reportsToIntArr(strings.Split(input, "\n"))
	solvePart1(reports)
	solvePart2(reports)
}

func solvePart1(reports [][]int) {
	safe := 0
	for _, r := range reports {
		ok := true
		asc := 0
		for i := 1; i < len(r); i++ {
			if !isDiffWithinLimit(r[i-1], r[i]) {
				ok = false
				break
			}

			isIncreasing := r[i-1] < r[i]
			if asc == 0 {
				if isIncreasing {
					asc = 1
				} else {
					asc = -1
				}
				continue
			}
			if (isIncreasing && asc != 1) || (!isIncreasing && asc != -1) {
				ok = false
				break
			}
		}

		if ok {
			safe++
		}
	}
	fmt.Println("Part 1:", safe)
}

func solvePart2(reports [][]int) {
	safe := 0
	for _, r := range reports {
		skip := -1
		for i := 0; i < len(r); i++ {
			ok := isValidReport(r, skip)
			if ok {
				safe++
				break
			} else {
				skip = i
			}
		}
	}
	fmt.Println("Part 2:", safe)
}

func isValidReport(r []int, skip int) bool {
	ok, asc := true, 0
	for i := 1; i < len(r)-1; i++ {
		next := i + 1
		if skip > -1 && skip == next {
			next++
		}
		if !isDiffWithinLimit(r[i], r[next]) {
			ok = false
			break
		}

		isIncreasing := r[i] < r[next]
		if asc == 0 {
			if isIncreasing {
				asc = 1
			} else {
				asc = -1
			}
			continue
		}
		if (isIncreasing && asc != 1) || (!isIncreasing && asc != -1) {
			ok = false
			break
		}
	}
	return ok
}

func isDiffWithinLimit(v1, v2 int) bool {
	diff := v1 - v2
	if diff < 0 {
		diff = -diff
	}
	return diff > 0 && diff < 4
}

func reportsToIntArr(reports []string) [][]int {
	intArr := [][]int{}
	for _, report := range reports {
		if report == "" {
			continue
		}
		arrStr := strings.Split(report, " ")
		arr := []int{}
		for _, v := range arrStr {
			n, _ := strconv.Atoi(v)
			arr = append(arr, n)
		}
		intArr = append(intArr, arr)
	}
	return intArr
}
