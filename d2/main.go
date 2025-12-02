package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countOccurance(input string, num int, maxCluster int) int {
	count := 0
	for i := 1; i <= maxCluster; i++ {
		if strings.Count(input, string(input[0:i])) == len(input)/i && len(input)%i == 0 {
			count += num
			return count
		}
	}
	return count
}

func main() {
	var (
		err      error
		lower    int
		upper    int
		part1sum int
		part2sum int
	)
	b, err := os.ReadFile("input-go.txt")
	if err != nil {
		fmt.Println("could not read file input", err)
	}
	DATA := strings.SplitSeq(string(b), ",")
	for d := range DATA {
		parts := strings.Split(d, "-")
		if lower, err = strconv.Atoi(parts[0]); err != nil {
			panic(err)
		}
		if upper, err = strconv.Atoi(parts[1]); err != nil {
			panic(err)
		}
		for i := lower; i <= upper; i++ {
			snum := strconv.Itoa(i)
			half := len(snum) / 2
			if snum[0:half] == snum[half:] {
				part1sum += i
			}
			part2sum += countOccurance(snum, i, half)
		}
	}
	fmt.Printf("Solution for Day 2 Part 1: %d\n", part1sum)
	fmt.Printf("Solution for Day 2 Part 2: %d\n", part2sum)
}
