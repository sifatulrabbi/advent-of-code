package solve2015

import (
	"slices"
	"strings"
)

var (
	naughtyCombos = []string{"ab", "cd", "pq", "xy"}
	vowelsTable   = map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
)

func SolveDay5Part1(input string) int {
	lines := strings.Split(input, "\n")
	totalNice := 0
	for _, l := range lines {
		if isNiceStringP1(l) {
			totalNice++
		}
	}
	return totalNice
}

func isNiceStringP1(str string) bool {
	arr := strings.Split(str, "")
	hasDouble := false
	vowels := 0
	prev := ""
	for _, s := range arr {
		if slices.Contains(naughtyCombos, prev+s) {
			hasDouble = false
			vowels = 0
			break
		}
		if prev == s {
			hasDouble = true
		}
		if _, ok := vowelsTable[s]; ok {
			vowels++
		}
		prev = s
	}
	if hasDouble && vowels >= 3 {
		return true
	}
	return false
}

func SolveDay5Part2(input string) []string {
	lines := strings.Split(input, "\n")
	totalNice := []string{}
	for _, l := range lines {
		if isNiceStringP2(l) {
			totalNice = append(totalNice, l)
		}
	}
	return totalNice
}

func isNiceStringP2(str string) bool {
	combo2 := map[string]int{}
	combo3 := 0
	arr := strings.Split(str, "")

	for i := 0; i < len(arr); i++ {
		if k := i + 2; k < len(arr) && arr[i] == arr[k] {
			combo3++
		}

		if j := i + 1; j < len(arr) {
			combo := arr[i] + arr[j]
			valid := true
			if arr[i] == arr[j] &&
				!(j+1 < len(arr) && arr[j] == arr[j+1]) &&
				i-1 >= 0 && arr[i] == arr[i-1] {
				valid = false
			}
			if valid {
				if _, ok := combo2[combo]; ok {
					combo2[combo]++
				} else {
					combo2[combo] = 1
				}
			}
		}
	}

	combo2count := 0
	for _, v := range combo2 {
		if v >= 2 {
			combo2count += v
		}
	}
	if combo3 > 0 && combo2count >= 2 {
		return true
	}
	return false
}
