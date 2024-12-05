package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// f, err := os.ReadFile("./solve2024/inputs/input-day4.txt")
	f, err := os.ReadFile("./solve2024/inputs/input-day4-test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	part1(string(f))
}

func part1(input string) {
	mtx := inputToMatrix(input)
	board := Board{
		mtx:           mtx,
		dimension:     [2]int{len(mtx) /* y */, len(mtx[0]) /* x */},
		moveDirection: nil,
		posY:          0,
		posX:          0,
	}
	count := 0

	for y := 0; y < board.dimension[0]; y++ {
		for x := 0; x < board.dimension[1]; x++ {
			char := board.mtx[y][x]
			fmt.Printf("x=%d y=%d v[%s] -> ", x, y, char)
			if char != "X" {
				fmt.Println("not [X]")
				continue
			}

			board.posY = y
			board.posX = x
			nextVal := getNextValidValue(char)
			for {
				if nextVal == nil {
					fmt.Println("fully matched!")
					count++
					board.moveDirection = nil
					break
				}
				fmt.Printf("next[%s] -> ", *nextVal)
				if !board.CheckNextValue(*nextVal) {
					fmt.Println("not matched!")
					board.moveDirection = nil
					break
				}
				nextVal = getNextValidValue(*nextVal)
			}
		}
	}
	fmt.Println("part1:", count)
}

func part2(board Board) {
}

var (
	R  = "right"
	D  = "down"
	L  = "left"
	U  = "up"
	UR = "up-right"
	UL = "up-left"
	DR = "down-right"
	DL = "down-left"
)

type Board struct {
	mtx           [][]string
	dimension     [2]int
	moveDirection *string
	posY          int
	posX          int
}

func (b *Board) CheckNextValue(next string) bool {
	if b.moveDirection != &L && next == b.GetRight(b.posY, b.posX) {
		b.moveDirection = &R
		b.ChangePos(0, 1)
		return true
	}
	if b.moveDirection != &U && next == b.GetDown(b.posY, b.posX) {
		b.moveDirection = &D
		b.ChangePos(1, 0)
		return true
	}
	if b.moveDirection != &R && next == b.GetLeft(b.posY, b.posX) {
		b.moveDirection = &L
		b.ChangePos(0, -1)
		return true
	}
	if b.moveDirection != &D && next == b.GetUp(b.posY, b.posX) {
		b.moveDirection = &U
		b.ChangePos(-1, 0)
		return true
	}
	if b.moveDirection != &DL && next == b.GetUpRight(b.posY, b.posX) {
		b.moveDirection = &UR
		b.ChangePos(-1, 1)
		return true
	}
	if b.moveDirection != &DR && next == b.GetUpLeft(b.posY, b.posX) {
		b.moveDirection = &UL
		b.ChangePos(-1, -1)
		return true
	}
	if b.moveDirection != &UL && next == b.GetDownRight(b.posY, b.posX) {
		b.moveDirection = &DR
		b.ChangePos(1, 1)
		return true
	}
	if b.moveDirection != &UR && next == b.GetDownLeft(b.posY, b.posX) {
		b.moveDirection = &DL
		b.ChangePos(1, -1)
		return true
	}
	return false
}

func (b Board) GetRight(y, x int) string {
	if x+1 >= b.dimension[1] {
		return ""
	}
	return b.mtx[y][x+1]
}

func (b Board) GetDown(y, x int) string {
	if y+1 >= b.dimension[0] {
		return ""
	}
	return b.mtx[y+1][x]
}

func (b Board) GetLeft(y, x int) string {
	if x-1 < 0 {
		return ""
	}
	return b.mtx[y][x-1]
}

func (b Board) GetUp(y, x int) string {
	if y-1 < 0 {
		return ""
	}
	return b.mtx[y-1][x]
}

func (b Board) GetUpLeft(y, x int) string {
	if y-1 < 0 || x-1 < 0 {
		return ""
	}
	return b.mtx[y-1][x-1]
}

func (b Board) GetUpRight(y, x int) string {
	if y-1 < 0 || x+1 >= b.dimension[1] {
		return ""
	}
	return b.mtx[y-1][x+1]
}

func (b Board) GetDownRight(y, x int) string {
	if y+1 >= b.dimension[0] || x+1 >= b.dimension[1] {
		return ""
	}
	return b.mtx[y+1][x+1]
}

func (b Board) GetDownLeft(y, x int) string {
	if y+1 >= b.dimension[0] || x-1 < 0 {
		return ""
	}
	return b.mtx[y+1][x-1]
}

func (b *Board) ChangePos(y, x int) {
	if b.posY+y >= 0 && b.posY+y < b.dimension[0] {
		b.posY += y
	}
	if b.posX+x >= 0 && b.posX+x < b.dimension[1] {
		b.posX += x
	}
}

func getNextValidValue(v string) *string {
	sequenceMap := map[string]string{
		"X": "M",
		"M": "A",
		"A": "S",
	}
	if next, ok := sequenceMap[v]; ok {
		return &next
	}
	return nil
}

func inputToMatrix(input string) [][]string {
	mtx := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		if len(line) < 1 {
			continue
		}
		mtx = append(mtx, strings.Split(line, ""))
	}
	return mtx
}
