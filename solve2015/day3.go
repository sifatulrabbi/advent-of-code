package solve2015

func genMatrix(l int) [][]int {
	mtx := make([][]int, l*2)
	for i := 0; i < l*2; i++ {
		x := make([]int, l*2)
		mtx[i] = x
	}
	return mtx
}

func SolveDay3Part1(moves string) int {
	var (
		l         = len(moves)
		mtx       = genMatrix(l)
		y         = l - 1
		x         = l - 1
		moveTable = map[string][2]int{
			"^": {0, -1},
			"v": {0, 1},
			">": {1, 0},
			"<": {-1, 0},
		}
	)

	mtx[y][x] += 1
	for i := 0; i < l; i++ {
		m := moveTable[string(moves[i])]
		x = x + m[0]
		y = y + m[1]
		mtx[y][x] += 1
	}

	housesReceived := 0
	for _, vy := range mtx {
		for _, vx := range vy {
			if vx > 0 {
				housesReceived++
			}
		}
	}
	return housesReceived
}

func SolveDay3Part2(moves string) int {
	var (
		l         = len(moves)
		mtx       = genMatrix(l)
		sy        = l - 1
		sx        = l - 1
		ry        = l - 1
		rx        = l - 1
		moveTable = map[string][2]int{
			"^": {0, -1},
			"v": {0, 1},
			">": {1, 0},
			"<": {-1, 0},
		}
	)

	mtx[sy][sx] += 1
	for i := 0; i < l; i++ {
		m := moveTable[string(moves[i])]
		if i%2 != 0 {
			// santa
			sx, sy = sx+m[0], sy+m[1]
			mtx[sy][sx] += 1
		} else {
			// robo-santa
			rx, ry = rx+m[0], ry+m[1]
			mtx[ry][rx] += 1
		}
	}

	housesReceived := 0
	for _, vy := range mtx {
		for _, vx := range vy {
			if vx > 0 {
				housesReceived++
			}
		}
	}
	return housesReceived
}
