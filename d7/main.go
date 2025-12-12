package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "<file>")
		return
	}
	b, err := os.ReadFile(os.Args[1])
	// b, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("could not read file", err)
		return
	}
	DATA := bytes.Split(b, []byte("\n"))
	// get index of start
	startIndex := bytes.Index(DATA[0], []byte("S"))
	part1 := 0
	// write first tachyon beam
	DATA[1][startIndex] = '|'
	for i := 1; i < len(DATA); i++ {
		for j := 0; j < len(DATA[i]); j++ {
			if DATA[i][j] == '^' {
				if j > 0 {
					DATA[i][j-1] = '|'
				}
				if j < len(DATA[i])-1 {
					DATA[i][j+1] = '|'
				}
			}
			if DATA[i-1][j] == '|' && DATA[i][j] != '^' {
				DATA[i][j] = '|'
			}
		}
	}
	// count beam pre split
	for i := 1; i < len(DATA); i++ {
		for j := 0; j < len(DATA[i]); j++ {
			if DATA[i][j] == '^' && DATA[i-1][j] == '|' {
				part1++
			}
		}
	}
	fmt.Print(string(bytes.Join(DATA, []byte("\n"))))
	fmt.Println()
	fmt.Println("Solution for Day 7 Part 1:", part1)
}
