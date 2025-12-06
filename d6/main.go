package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "<file>")
		return
	}
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("could not read file", err)
		return
	}
	DATA := strings.Split(string(b), "\n")
	matrix := [][]int{}
	operators := []string{}
	multispace := regexp.MustCompile(`\s+`)
	for i, d := range DATA {
		d = multispace.ReplaceAllString(d, " ")
		d = strings.TrimLeft(strings.TrimRight(d, " "), " ")
		if i == len(DATA)-1 {
			operators = strings.Split(d, " ")
			continue
		}
		nums := strings.Split(d, " ")
		line := make([]int, len(nums))
		for i, n := range nums {
			num, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println("could not convert to number", err)
				return
			}
			line[i] = num
		}
		matrix = append(matrix, line)
	}
	sum := 0
	for i, o := range operators {
		colsum := 0
		switch o {
		case "+":
			for _, k := range matrix {
				colsum += k[i]
			}
		case "*":
			colsum = 1
			for _, k := range matrix {
				colsum *= k[i]
			}
		default:
			fmt.Println("unknown operand", o)
			return
		}
		sum += colsum
	}
	fmt.Println("Solution for Day 6 Part 1:", sum)
}
