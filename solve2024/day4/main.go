package main

import (
	"fmt"
	"os"
	"strings"
)

// y, x
var directions = [][2]int{
	{0, 1},   // right
	{1, 0},   // down
	{0, -1},  // left
	{-1, 0},  // up
	{-1, 1},  // up-right
	{-1, -1}, // up-left
	{1, 1},   // down-right
	{1, -1},  // down-left
}

func main() {
	f, _ := os.ReadFile("./solve2024/inputs/input-day4.txt")
	// f, _ := os.ReadFile("./solve2024/inputs/input-day4-test.txt")
	mtx := inputToMatrix(string(f))

	part1(mtx)
	part2(mtx)
}

func part1(mtx Matrix) {
	count := 0
	dim := [2]int{len(mtx), len(mtx[0])}

	for y := 0; y < dim[0]; y++ {
		for x := 0; x < dim[1]; x++ {
			char := mtx[y][x]
			if char != "X" {
				continue
			}
			for _, dir := range directions {
				if mtx.CheckValidSequence(y, x, dir) {
					count++
				}
			}
		}
	}

	fmt.Println("Part1:", count)
}

func part2(mtx Matrix) {
	count := 0
	dim := [2]int{len(mtx), len(mtx[0])}

	for y := 0; y < dim[0]; y++ {
		for x := 0; x < dim[1]; x++ {
			char := mtx[y][x]
			if char != "A" {
				continue
			}
			if mtx.CheckCrossSequence(y, x) {
				count++
			}
		}
	}

	fmt.Println("Part2:", count)
}

type Matrix [][]string

func (m Matrix) getByDirection(y, x int, direction [2]int) string {
	nextY := y + direction[0]
	nextX := x + direction[1]
	if nextY < 0 || nextY >= len(m) || nextX < 0 || nextX >= len(m[nextY]) {
		return ""
	}
	return m[nextY][nextX]
}

func (m Matrix) CheckValidSequence(y, x int, dir [2]int) bool {
	sequence := []string{"M", "A", "S"}
	currY, currX := y, x
	for _, s := range sequence {
		if s != m.getByDirection(currY, currX, dir) {
			return false
		}
		currY += dir[0]
		currX += dir[1]
	}
	return true
}

func (m Matrix) CheckCrossSequence(y, x int) bool {
	ul := [2]int{-1, -1} // upper-left
	br := [2]int{+1, +1} // bottom-right
	ur := [2]int{-1, +1} // upper-right
	bl := [2]int{+1, -1} // bottom-left
	matched := 0
	if (m.getByDirection(y, x, ul) == "M" && m.getByDirection(y, x, br) == "S") ||
		(m.getByDirection(y, x, ul) == "S" && m.getByDirection(y, x, br) == "M") {
		matched++
	}
	if (m.getByDirection(y, x, ur) == "M" && m.getByDirection(y, x, bl) == "S") ||
		(m.getByDirection(y, x, ur) == "S" && m.getByDirection(y, x, bl) == "M") {
		matched++
	}
	return matched == 2
}

func inputToMatrix(input string) Matrix {
	mtx := Matrix{}
	for _, line := range strings.Split(input, "\n") {
		if len(line) < 1 {
			continue
		}
		mtx = append(mtx, strings.Split(line, ""))
	}
	return mtx
}
