package main

import (
	"fmt"
	"github.com/majustman/word-count/service"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(fmt.Errorf("number of arguments equal %v, expected 1", len(os.Args)-1))
		return
	}
	s := os.Args[1]
	m := service.WordCount(s)
	for word, count := range m {
		fmt.Printf("\"%v\" - %v\n", word, count)
	}
}
