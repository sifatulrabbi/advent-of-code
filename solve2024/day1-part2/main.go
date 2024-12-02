package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// inputFile := "./solve2024/input-day1-2024.test.txt"
	inputFile := "./solve2024/input-day1-2024.txt"

	f, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	input := string(f)
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	l, r := []int{}, []int{}
	for _, line := range lines {
		sides := strings.Split(line, "   ")
		leftSide, _ := strconv.Atoi(sides[0])
		rightSide, _ := strconv.Atoi(sides[1])
		l = append(l, leftSide)
		r = append(r, rightSide)
	}

	var totalSimilarity int = 0
	for _, v := range l {
		totalSimilarity += similarityScore(r, v)
	}

	fmt.Println(totalSimilarity)
}

func similarityScore(arr []int, n int) int {
	appearance := 0
	for _, v := range arr {
		if v == n {
			appearance++
		}
	}
	// fmt.Printf("similarity for %d is %d\n", n, appearance*n)
	return appearance * n
}
