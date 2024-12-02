package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("./solve2015/inputs/input-day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(f)

	part1 := ""
	part2 := ""
	for i := 0; ; i++ {
		hash := genMD5Hash(fmt.Sprintf("%s%d", input, i))
		if strings.HasPrefix(hash, "000000") {
			part2 = fmt.Sprintf("%d", i)
		} else if strings.HasPrefix(hash, "00000") {
			part1 = fmt.Sprintf("%d", i)
		}
		if part1 != "" && part2 != "" {
			break
		}
	}

	fmt.Printf("part1: %s\npart2: %s\n", part1, part2)
}

func genMD5Hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
