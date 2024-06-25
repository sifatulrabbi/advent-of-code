package solve2015

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	OFF int = iota
	ON
	TOGGLE
)

func SolveDay6Part1(input string) int {
	instructions, actions := parseInstructions(input)
	grid := getLightsGrid()

	for i, ins := range instructions {
		yDist := ins[1][0] - ins[0][0]
		xDist := ins[1][1] - ins[0][1]
		for y := 0; y <= yDist; y++ {
			for x := 0; x <= xDist; x++ {
				yi := ins[0][0] + y
				xi := ins[0][1] + x
				switch actions[i] {
				case OFF:
					grid[yi][xi] = false
					break
				case ON:
					grid[yi][xi] = true
					break
				case TOGGLE:
					grid[yi][xi] = !grid[yi][xi]
					break
				}
			}
		}
	}

	totalLit := 0
	for _, row := range grid {
		for _, v := range row {
			if v {
				totalLit++
			}
		}
	}

	return totalLit
}

func parseInstructions(input string) ([][][]int, []int) {
	var (
		instructions [][][]int
		actions      []int
	)
	r := regexp.MustCompile("([0-9]*,[0-9]*)*")
	for _, v := range strings.Split(input, "\n") {
		matches := r.FindAllString(v, -1)
		set := [][]int{}
		for _, m := range matches {
			if m == "" {
				continue
			}

			loc := []int{}
			locStr := strings.Split(m, ",")

			if len(locStr) != 2 {
				log.Panicf("need x and y value pairs to find a light. found '%v' from '%s'", locStr, m)
			}

			for _, p := range locStr {
				if n, err := strconv.ParseInt(p, 10, 32); err != nil {
					log.Panicf("unable to convert %s into a integer", p)
				} else {
					loc = append(loc, int(n))
				}
			}
			set = append(set, loc)
		}

		instructions = append(instructions, set)

		if strings.Contains(v, "turn on") {
			actions = append(actions, ON)
		} else if strings.Contains(v, "turn off") {
			actions = append(actions, OFF)
		} else {
			actions = append(actions, TOGGLE)
		}

	}
	return instructions, actions
}

func getLightsGrid() [][]bool {
	mtx := make([][]bool, 1000)
	for i := 0; i < len(mtx); i++ {
		row := make([]bool, 1000)
		mtx[i] = row
	}
	return mtx
}
