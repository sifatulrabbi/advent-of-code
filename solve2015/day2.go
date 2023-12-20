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
