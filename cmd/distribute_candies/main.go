package main

import (
	"fmt"
	"io"
	"log"

	s "github.com/dantin/leetcode-go/distribute-candies"
)

func main() {
	nums := make([]int, 0)
	var i int
	var err error
	for true {
		_, err = fmt.Scan(&i)
		if err == nil {
			nums = append(nums, i)
		} else if err == io.EOF {
			break
		} else {
			log.Println("scan for i failed, due to ", err)
		}
	}
	fmt.Printf("input: %v\nanswer: %d\n", nums, s.DistributeCandies(nums))
}
