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
}

func solvePart1(reports [][]int) {
	safe := 0
	for _, r := range reports {
		ok := true
		for i := 1; i < len(r)-1; i++ {
			if !isDiffWithinLimit(r[i-1], r[i]) || !isDiffWithinLimit(r[i], r[i+1]) {
				ok = false
				break
			}
			if (r[i-1] < r[i] && !(r[i] < r[i+1])) ||
				(r[i-1] > r[i] && !(r[i] > r[i+1])) {
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

func isDiffWithinLimit(v1, v2 int) bool {
	diff := v1 - v2
	if diff < 0 {
		diff = -diff
	}
	return diff < 4
}

func reportsToIntArr(reports []string) [][]int {
	intArr := [][]int{}
	for _, report := range reports {
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
