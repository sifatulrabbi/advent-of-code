package utils

import (
	"log"
	"os"
	"strconv"
)

// ReadInputFile(day, year, part??)
func ReadInputFile(args ...int) string {
	if len(args) < 2 {
		log.Fatal("Not enough args")
	}
	day, year, part := args[0], args[1], 0
	if len(args) > 2 {
		part = args[2]
	}
	fileName := "input-" + "day" + strconv.Itoa(day) + "-"
	if part > 0 {
		fileName += "part" + strconv.Itoa(part) + "-"
	}
	fileName += strconv.Itoa(year) + ".txt"
	f, err := os.ReadFile(fileName)
	log.Fatal(err)
	return string(f)
}
