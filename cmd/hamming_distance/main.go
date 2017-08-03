package main

import (
	"fmt"
	"log"

	s "github.com/dantin/leetcode-go/hamming-distance"
)

func main() {
	var i, j int
	if _, err := fmt.Scan(&i, &j); err != nil {
		log.Println("Scan for i, j failed, due to ", err)
	}
	fmt.Println(s.HammingDistance(i, j))
}
