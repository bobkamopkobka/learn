package main

import (
	"fmt"
	"log"
	"os"
	"github.com/bobkamopkobka/learn/readfromfile"
)

func main() {

	file, err := os.Open("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines, err := readfromfile.GetLines(file)
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}
