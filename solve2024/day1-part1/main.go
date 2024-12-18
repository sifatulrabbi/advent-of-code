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
	// inputFile, err := filepath.Abs(inputFile)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

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

	ascSort(l)
	ascSort(r)

	var distance int = 0
	for i, v := range l {
		d := v - r[i]
		if d < 0 {
			d = -d
		}
		distance += d
	}

	fmt.Println(distance)
}

func ascSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
