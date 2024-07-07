package solve2015

import (
	"fmt"
	"strconv"
	"strings"
)

type DataTable map[string]int

func (dt *DataTable) and(w string, values []string) {
	var (
		prevx int
		prevy int
	)
	if px, ok := (*dt)[values[0]]; ok {
		prevx = px
	} else {
		prevx = Int(values[0])
	}
	if py, ok := (*dt)[values[2]]; ok {
		prevy = py
	} else {
		prevy = Int(values[2])
	}
	(*dt)[w] = prevx & prevy
	fmt.Printf("%d AND %d -> %s\n", prevx, prevy, w)
}

func (dt *DataTable) or(w string, values []string) {
	var (
		prevx int
		prevy int
	)
	if px, ok := (*dt)[values[0]]; ok {
		prevx = px
	} else {
		prevx = Int(values[0])
	}
	if py, ok := (*dt)[values[2]]; ok {
		prevy = py
	} else {
		prevy = Int(values[2])
	}
	(*dt)[w] = prevx | prevy
	fmt.Printf("%d OR %d -> %s\n", prevx, prevy, w)
}

func (dt *DataTable) lshift(w string, values []string) {
	var (
		prevx int
		prevy int
	)
	if px, ok := (*dt)[values[0]]; ok {
		prevx = px
	} else {
		prevx = Int(values[0])
	}
	if py, ok := (*dt)[values[2]]; ok {
		prevy = py
	} else {
		prevy = Int(values[2])
	}
	(*dt)[w] = prevx << prevy
	fmt.Printf("%d LSHIFT %d -> %s\n", prevx, prevy, w)
}

func (dt *DataTable) rshift(w string, values []string) {
	var (
		prevx int
		prevy int
	)
	if px, ok := (*dt)[values[0]]; ok {
		prevx = px
	} else {
		prevx = Int(values[0])
	}
	if py, ok := (*dt)[values[2]]; ok {
		prevy = py
	} else {
		prevy = Int(values[2])
	}
	(*dt)[w] = prevx >> prevy
	fmt.Printf("%d RSHIFT %d -> %s\n", prevx, prevy, w)
}

func (dt *DataTable) not(w string, values []string) {
	var prev int
	if px, ok := (*dt)[values[1]]; ok {
		prev = px
	} else {
		px = Int(values[1])
	}
	(*dt)[w] = 65535 ^ prev
	fmt.Printf("NOT %d -> %s\n", prev, w)
}

func (dt *DataTable) assign(w string, values []string) {
	(*dt)[w] = Int(values[0])
	fmt.Printf("%s -> %s\n", values[0], w)
}

func SolveDay7Part1Bak(input string) DataTable {
	var (
		dataTable = DataTable{}
		lines     = strings.Split(input, "\n")
	)

	for _, l := range lines {
		segments := strings.Split(l, " -> ")
		wire := segments[1]
		operands := strings.Split(segments[0], " ")

		switch len(operands) {
		case 1:
			dataTable.assign(wire, operands)
			break
		case 3:
			switch operands[1] {
			case "AND":
				dataTable.and(wire, operands)
				break
			case "OR":
				dataTable.or(wire, operands)
				break
			case "RSHIFT":
				dataTable.rshift(wire, operands)
				break
			case "LSHIFT":
				dataTable.lshift(wire, operands)
				break
			}
			break
		case 2:
			switch operands[0] {
			case "NOT":
				dataTable.not(wire, operands)
				break
			}
			break
		}
	}
	return dataTable
}

func Int(str string) int {
	n, _ := strconv.ParseInt(str, 10, 32)
	return int(n)
}
