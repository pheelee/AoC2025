package main

import (
	"fmt"
	"os"
	"strings"
)

func solve(input []string) (int, []string) {
	cols := len(input[0])
	rows := len(input)
	var updatedStorage = make([]string, len(input))
	removableRolls := 0
	for i := range rows {
		line := []byte(input[i])
		for j := range cols {
			if input[i][j] == '@' {
				rollsCount := 0
				if i > 0 {
					if input[i-1][j] == '@' {
						rollsCount++
					}
					if j > 0 && input[i-1][j-1] == '@' {
						rollsCount++
					}
					if j < rows-1 && input[i-1][j+1] == '@' {
						rollsCount++
					}
				}
				if i < rows-1 {
					if input[i+1][j] == '@' {
						rollsCount++
					}
					if j > 0 && input[i+1][j-1] == '@' {
						rollsCount++
					}
					if j < cols-1 && input[i+1][j+1] == '@' {
						rollsCount++
					}
				}
				if j > 0 && input[i][j-1] == '@' {
					rollsCount++
				}
				if j < cols-1 && input[i][j+1] == '@' {
					rollsCount++
				}
				if rollsCount < 4 {
					removableRolls++
					line[j] = 'X'
				}
			}
		}
		updatedStorage[i] = string(line)
	}
	return removableRolls, updatedStorage
}

func main() {
	b, err := os.ReadFile("input-go.txt")
	if err != nil {
		fmt.Println("could not read file", err)
		return
	}
	DATA := strings.Split(string(b), "\n")
	part1, _ := solve(DATA)

	part2 := 0
	input := DATA
	for {
		num, result := solve(input)
		if num == 0 {
			break
		}
		part2 += num
		input = result
	}

	fmt.Println("Solution for Day 4 Part 1:", part1)
	fmt.Println("Solution for Day 4 Part 2:", part2)
}
