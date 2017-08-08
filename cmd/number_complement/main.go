package main

import (
	"fmt"
	"log"

	s "github.com/dantin/leetcode-go/number-complement"
)

func main() {
	var num int
	if _, err := fmt.Scan(&num); err != nil {
		log.Println("Scan for num failed, due to ", err)
	}
	fmt.Println(s.FindComplement(num))
}
