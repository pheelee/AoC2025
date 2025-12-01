package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Safe struct {
	CurrentState int
	ZeroCounter  int
}

func (s *Safe) rotate(instruction string, part int) error {
	num, err := strconv.Atoi(string(instruction[1:]))
	if err != nil {
		return err
	}
	// Records previous state which is needed to avoid duplicate counts when there was a 0 last time but only for the first occurance
	previousState := s.CurrentState
	switch instruction[0] {
	case 'L':
		for range num {
			s.CurrentState -= 1
			if s.CurrentState == -1 {
				s.CurrentState = 99
				if part == 2 && previousState != 0 {
					s.ZeroCounter++
				} else {
					previousState = s.CurrentState
				}
			}
		}
	case 'R':
		i := 0
		for range num {
			i++
			s.CurrentState += 1
			if s.CurrentState == 100 {
				s.CurrentState = 0
				// Only increase zero position counter when not on last rotation
				if part == 2 && i < num {
					s.ZeroCounter++
				}
			}
		}
	default:
		return fmt.Errorf("invalid input: %b", instruction[0])
	}
	if s.CurrentState == 0 {
		s.ZeroCounter++
	}
	return nil
}

func solve(data []string, part int) (int, error) {
	s := Safe{CurrentState: 50, ZeroCounter: 0}
	for _, d := range data {
		if err := s.rotate(d, part); err != nil {
			return 0, err
		}
	}
	return s.ZeroCounter, nil
}

func main() {
	b, err := os.ReadFile("input-go.txt")
	if err != nil {
		fmt.Println("Could not load data", err)
		return
	}
	DATA := strings.Split(string(b), "\n")
	// Part 1
	solution, err := solve(DATA, 1)
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
	fmt.Printf("Solution for Day 1 Part 1: %d\n", solution)
	// Part 2
	solution, err = solve(DATA, 2)
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
	fmt.Printf("Solution for Day 1 Part 2: %d\n", solution)
}
