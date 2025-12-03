package main

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

func findMax(input iter.Seq[string]) int {
	var result int
	for d := range input {
		highest := 0
		for i := 0; i < len(d); i++ {
			for j := 0; j < len(d); j++ {
				if i == j {
					continue
				}
				num, _ := strconv.Atoi(string([]byte{d[i], d[j]}))
				if num > highest && i < j {
					highest = num
				}
			}
		}
		fmt.Printf("Input: %s, Joltage: %d\n", d, highest)
		result += highest
	}
	return result
}

func main() {
	var (
		b   []byte
		err error
	)
	if b, err = os.ReadFile("input-go.txt"); err != nil {
		panic(err)
	}
	DATA := strings.SplitSeq(string(b), "\n")

	fmt.Println("Solution for Day 3 Part 1:", findMax(DATA))
}
