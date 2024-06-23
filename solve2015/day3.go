package solve2015

func genMatrix(l int) [][]int {
	mtx := make([][]int, l*2)
	for i := 0; i < l*2; i++ {
		x := make([]int, l*2)
		mtx[i] = x
	}
	return mtx
}

func SolveDay3(moves string) int {
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
