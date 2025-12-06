package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type database struct {
	freshRanges [][]int
}

func (d *database) parseRange(s string) error {
	var (
		start, end int
		err        error
	)
	p := strings.Split(s, "-")
	if start, err = strconv.Atoi(p[0]); err != nil {
		return err
	}
	if end, err = strconv.Atoi(p[1]); err != nil {
		return err
	}
	d.freshRanges = append(d.freshRanges, []int{start, end})
	return nil
}

func (d *database) isFresh(id int) bool {
	for _, i := range d.freshRanges {
		if id >= i[0] && id <= i[1] {
			return true
		}
	}
	return false
}
func (d *database) getLowestId() int {
	lowest := d.freshRanges[0][0]
	for _, i := range d.freshRanges {
		if i[0] < lowest {
			lowest = i[0]
		}
	}
	return lowest
}
func (d *database) getHighestId() int {
	highest := d.freshRanges[0][1]
	for _, i := range d.freshRanges {
		if i[1] > highest {
			highest = i[1]
		}
	}
	return highest
}

func main() {
	b, err := os.ReadFile("input-go.txt")
	if err != nil {
		fmt.Println("cannot read file input", err)
		return
	}
	DATA := strings.Split(string(b), "\n")
	db := database{freshRanges: [][]int{}}
	samples := []int{}
	for _, i := range DATA {
		if strings.Contains(i, "-") {
			if err := db.parseRange(i); err != nil {
				fmt.Println("could not parse freshRange", err)
				return
			}
			continue
		}
		if i != "" {
			s, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println("could not parse sample", err)
				return
			}
			samples = append(samples, s)
		}
	}
	day1 := 0
	for _, s := range samples {
		if db.isFresh(s) {
			day1++
		}
	}
	fmt.Printf("Solution for Day 5 Part 1: %d\n", day1)
	low := db.getLowestId()
	high := db.getHighestId()
	fmt.Printf("lowest: %d, highest: %d\n", low, high)
	// fmt.Printf("Solution for Day 5 Part 2: %d\n", db.uniqueFreshCount())
}
