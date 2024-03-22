package main

import (
	"fmt"
	"log"
	"readfromfile"
)

func main() {
	lines, err := readfromfile.GetLines("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}
