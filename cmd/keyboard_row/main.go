package main

import (
	"fmt"
	"io"
	"log"

	s "github.com/dantin/leetcode-go/keyboard-row"
)

func main() {
	var (
		line string
		err  error
	)
	input := make([]string, 0)

	for true {
		_, err = fmt.Scan(&line)
		if err == nil {
			input = append(input, line)
		} else if err == io.EOF {
			break
		} else {
			log.Println("Scan for line failed, due to ", err)
		}
	}
	fmt.Printf("output:\n%v\n", s.FindWords(input))
}
