package solve2023

import "testing"

func TestDay1(t *testing.T) {
	val := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	if result := Day1(val); result != 142 {
		t.Errorf("result must be 142, got: %d", result)
	}
}

func TestDay1part2(t *testing.T) {
	val := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	if result := Day1Part2(val); result != 281 {
		t.Errorf("result must be 281, got: %d", result)
	}
}
