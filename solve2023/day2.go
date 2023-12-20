package solve2023

import (
	"strconv"
	"strings"
)

func Day2(input string) int {
	gameIds := []int{}

	for _, line := range strings.Split(input, "\n") {
		idAndSetArr := strings.Split(line, ": ")
		gameId := strings.Split(idAndSetArr[0], "Game ")[1]

		sets := idAndSetArr[1]
		gamePossible := true

		for _, set := range strings.Split(sets, "; ") {
			gamePossible = true
			requiredCubes := map[string]int{}

			for _, cube := range strings.Split(set, ", ") {
				info := strings.Split(cube, " ")
				n, _ := strconv.ParseInt(info[0], 10, 32)
				requiredCubes[info[1]] += int(n)
			}

			if requiredCubes["red"] > 12 || requiredCubes["blue"] > 14 || requiredCubes["green"] > 13 {
				gamePossible = false
				break
			}
		}

		if gamePossible {
			intGameId, _ := strconv.ParseInt(gameId, 10, 32)
			gameIds = append(gameIds, int(intGameId))
		}

	}

	total := 0
	for _, id := range gameIds {
		total += id
	}
	return total
}
