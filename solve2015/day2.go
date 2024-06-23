package solve2015

import (
	"strconv"
	"strings"
)

func Day2(dimentions string) int {
	totalSq := 0

	for _, line := range strings.Split(dimentions, "\n") {
		dims := []int{}
		smallestSide := 0

		for _, v := range strings.Split(line, "x") {
			d, _ := strconv.ParseInt(v, 10, 32)
			dims = append(dims, int(d))
		}

		l := len(dims)
		for i := 0; i < l; i++ {
			// calc the area
			area := 0
			if (i + 1) < l {
				area += dims[i] * dims[i+1]
			} else {
				area += dims[i] * dims[0]
			}

			// find the smallest area
			if smallestSide == 0 {
				smallestSide = area
			}
			if smallestSide > area {
				smallestSide = area
			}

			totalSq += (2 * area)
		}

		totalSq += smallestSide
	}

	return totalSq
}

func Day2Part2(boxes string) int {
	totalRibbonLn := 0
	boxList := strings.Split(boxes, "\n")
	for _, box := range boxList {
		dimStr := strings.Split(box, "x")
		dims := []int{}
		for _, d := range dimStr {
			n, _ := strconv.ParseInt(d, 10, 32)
			dims = append(dims, int(n))
		}

		for i := 0; i < len(dims); i++ {
			j := i + 1
			if j >= len(dims) {
				continue
			}
			if dims[i] > dims[j] {
				dims[i], dims[j] = dims[j], dims[i]
			}
		}

		ribbonLn := 2*(dims[0]+dims[1]) + (dims[0] * dims[1] * dims[2])
		totalRibbonLn += ribbonLn
	}
	return totalRibbonLn
}
