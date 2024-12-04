package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	// f, err := os.ReadFile("./solve2023/inputs/input-day3.txt")
	f, err := os.ReadFile("./solve2023/inputs/input-day3-test.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	part1(string(f))
	part2(string(f))
}

func part1(input string) {
}

func part2(input string) {
}

func parseInput(input string) []string {
	arr := []string{}
	for i := 0; i < len(input); i++ {
		c := string(input[i])
		if c == "." {
			continue
		}
		n := ""
		for j := 0; j < len(c); j++ {
			if isNumber(string(c[j])) {
				n += string(c[j])
			}
		}
		arr = append(arr, n)
	}
	return arr
}

func isNumber(c string) bool {
	_, err := strconv.Atoi(c)
	return err == nil
}
