package solve2015

func Day1(path string) int {
	floor := 0
	arr := []rune(path)
	for i := 0; i < len(arr); i++ {
		v := string(arr[i])
		if v == "(" {
			floor++
		} else if v == ")" {
			floor--
		}
	}
	return floor
}

func Day1Part2(path string) int {
	floor := 0
	pos := -1
	arr := []rune(path)
	for i := 0; i < len(arr); i++ {
		v := string(arr[i])
		if v == "(" {
			floor++
		} else if v == ")" {
			floor--
		}
		if floor == -1 {
			pos = i
			break
		}
	}
	return pos + 1
}
