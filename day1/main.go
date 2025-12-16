package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solvePart1(input string) int {
	parts := strings.Fields(input)
	current := 50
	count := 0

	for _, part := range parts {
		if len(part) < 2 {
			continue
		}
		direction := part[0]
		valStr := part[1:]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			fmt.Printf("Error parsing value: %s\n", valStr)
			continue
		}

		if direction == 'R' {
			current = (current + val) % 100
		} else if direction == 'L' {
			current = (current - val) % 100
			if current < 0 {
				current += 100
			}
		}

		if current == 0 {
			count++
		}
	}
	return count
}

func solvePart2(input string) int {
	parts := strings.Fields(input)
	current := 50
	count := 0

	for _, part := range parts {
		if len(part) < 2 {
			continue
		}
		direction := part[0]
		valStr := part[1:]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			fmt.Printf("Error parsing value: %s\n", valStr)
			continue
		}

		for i := 0; i < val; i++ {
			if direction == 'R' {
				current = (current + 1) % 100
			} else if direction == 'L' {
				current = (current - 1) % 100
				if current < 0 {
					current += 100
				}
			}
			if current == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	// Try to read from input.txt first
	content, err := os.ReadFile("input.txt")
	if err != nil {
		// If fails, use the example input for demonstration
		fmt.Println("Could not read input.txt, using example input:")
		exampleInput := "L68 L30 R48 L5 R60 L55 L1 L99 R14 L82"

		fmt.Printf("Example Part 1: %d\n", solvePart1(exampleInput))
		fmt.Printf("Example Part 2: %d\n", solvePart2(exampleInput))
		return
	}

	input := string(content)
	if strings.TrimSpace(input) == "" {
		fmt.Println("input.txt is empty, using example input:")
		exampleInput := "L68 L30 R48 L5 R60 L55 L1 L99 R14 L82"

		fmt.Printf("Example Part 1: %d\n", solvePart1(exampleInput))
		fmt.Printf("Example Part 2: %d\n", solvePart2(exampleInput))
		return
	}

	fmt.Printf("Part 1 Result: %d\n", solvePart1(input))
	fmt.Printf("Part 2 Result: %d\n", solvePart2(input))
}
