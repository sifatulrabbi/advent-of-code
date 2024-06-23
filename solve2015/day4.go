package solve2015

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
)

func SolveDay4Part1(secret string) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(secret)); err != nil {
		log.Fatalln(err)
	}
	tries := 100_000_000_000
	for i := 0; i < tries; i++ {
		md5Sig := hex.EncodeToString(hasher.Sum(nil))
		fmt.Println(md5Sig)
	}
}
