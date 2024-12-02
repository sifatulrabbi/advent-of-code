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
		if isValidReport(r, -1) {
			safe++
		}
	}
	fmt.Println("Part 1:", safe)
}

func solvePart2(reports [][]int) {
	safe := 0
	for _, r := range reports {
		for i := -1; i < len(r); i++ {
			if isValidReport(r, i) {
				safe++
				break
			}
		}
	}
	fmt.Println("Part 2:", safe)
}

func isValidReport(r []int, skip int) bool {
	ok, asc := true, 0
	for i := 0; i < len(r)-1; i++ {
		if i == skip {
			continue
		}

		next := i + 1
		if next == skip {
			if next+1 < len(r) {
				next++
			} else {
				break
			}
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
