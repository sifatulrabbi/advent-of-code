package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// f, err := os.ReadFile("./solve2024/inputs/input-day4.txt")
	// f, err := os.ReadFile("./solve2024/inputs/input-day4-test.txt")
	f, err := os.ReadFile("./solve2024/inputs/input-day4-test2.txt")
	if err != nil {
		log.Fatalln(err)
	}

	part1(string(f))
}

func part1(input string) {
	mtx := *inputToMatrix(input)
	for _, r := range mtx {
		fmt.Println(r)
	}
}

func part2(input string) {
}

func inputToMatrix(input string) *[][]string {
	mtx := [][]string{}

	for _, line := range strings.Split(input, "\n") {
		if len(line) < 1 {
			continue
		}
		buf := []string{}
		for i := 0; i < len(line); i++ {
			buf = append(buf, string(line[i]))
		}
		mtx = append(mtx, buf)
	}

	return &mtx
}
