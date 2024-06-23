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
		if isNiceString(l) {
			totalNice++
		}
	}
	return totalNice
}

func isNiceString(str string) bool {
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
