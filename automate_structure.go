package main

import (
	"fmt"
	"os"
)

var years = []string{
	"2015",
	"2016",
	"2017",
	"2018",
	"2019",
	"2020",
	"2021",
	"2022",
	"2023",
}

var days = []string{
	"day1",
	"day2",
	"day3",
	"day4",
	"day5",
	"day6",
	"day7",
	"day8",
	"day9",
	"day10",
	"day11",
	"day12",
	"day13",
	"day14",
	"day15",
	"day16",
	"day17",
	"day18",
	"day19",
	"day20",
	"day21",
	"day22",
	"day23",
	"day24",
	"day25",
}

// automate the process of creating files and test files.
func automateStructure() {
	prepareYear()
}

// this will create directories where the solve files will be stored.
func prepareYear() {
	for _, y := range years {
		dirName := fmt.Sprintf("solve%s", y)
		_, err := os.Stat(dirName)
		if os.IsNotExist(err) {
			if err := os.Mkdir(dirName, 0o755); err != nil {
				panic(err)
			}
		}

		prepareDays(dirName)
	}
}

func prepareDays(dir string) {
	for _, day := range days {
		fileName := fmt.Sprintf("%s/%s.go", dir, day)
		_, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			content := fmt.Sprintf("package %s", dir)
			os.WriteFile(fileName, []byte(content), 0o644)
		}
	}
}
