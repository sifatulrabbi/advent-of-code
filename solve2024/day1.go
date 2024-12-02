package solve2024

import (
	"strconv"
	"strings"
)

func Day1Part1(input string) int {
	// input := utils.ReadInputFile(1, 2024)
	// fmt.Println(input)

	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	l, r := []int64{}, []int64{}
	for _, line := range lines {
		sides := strings.Split(line, "   ")
		leftSide, _ := strconv.ParseInt(sides[0], 10, 32)
		rightSide, _ := strconv.ParseInt(sides[1], 10, 32)
		l = append(l, leftSide)
		r = append(r, rightSide)
	}

	ascSort(l)
	ascSort(r)

	var distance int64 = 0
	for i, v := range l {
		d := v - r[i]
		if d < 0 {
			d = -d
		}
		// fmt.Println(v, r[i], "=", d)
		distance += d
	}

	return int(distance)
}

func ascSort(arr []int64) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
