package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("./inputs/input-day7-test.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(f)
	for _, c := range input {
	}
}
