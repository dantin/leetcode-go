package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	s "github.com/dantin/leetcode-go/reverse-words-in-a-string-iii"
)

func main() {
	// read whole line
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Scan for string failed, due to ", err)
	}
	// ignore carriage return key
	input = strings.TrimSpace(input)
	fmt.Printf("%s\n", s.ReverseWords(input))
}
