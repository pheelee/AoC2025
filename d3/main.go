package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMax(batteryCount int, input []string) int {
	var result int
	findNextNumber := func(input string, start, position int) (byte, int) {
		var h byte = 0
		var pos int = start
		for i := start; i <= len(input)-(batteryCount-position); i++ {
			if input[i] > h {
				h = input[i]
				pos = i + 1
			}
		}
		return h, pos
	}
	for _, row := range input {
		joltage := make([]byte, batteryCount)
		startPos := 0
		for i := range joltage {
			joltage[i], startPos = findNextNumber(row, startPos, i)
		}
		sum, err := strconv.Atoi(string(joltage))
		if err != nil {
			panic(err)
		}
		result += sum
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
	DATA := strings.Split(string(b), "\n")

	fmt.Println("Solution for Day 3 Part 1:", findMax(2, DATA))
	fmt.Println("Solution for Day 3 Part 2:", findMax(12, DATA))
}
