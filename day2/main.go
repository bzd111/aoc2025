package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalid(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
}

func isInvalidPart2(n int) bool {
	s := strconv.Itoa(n)
	L := len(s)
	for length := 1; length <= L/2; length++ {
		if L%length != 0 {
			continue
		}
		pattern := s[:length]
		repeated := strings.Repeat(pattern, L/length)
		if s == repeated {
			return true
		}
	}
	return false
}

func solvePart1(input string) int {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\n", "")
	ranges := strings.Split(input, ",")
	totalSum := 0

	for _, r := range ranges {
		if r == "" {
			continue
		}
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			if isInvalid(i) {
				totalSum += i
			}
		}
	}
	return totalSum
}

func solvePart2(input string) int {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\n", "")
	ranges := strings.Split(input, ",")
	totalSum := 0

	for _, r := range ranges {
		if r == "" {
			continue
		}
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			if isInvalidPart2(i) {
				totalSum += i
			}
		}
	}
	return totalSum
}

func main() {
	// Try to read from input.txt first
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Could not read input.txt, using example input:")
		runExample()
		return
	}

	input := string(content)
	if strings.TrimSpace(input) == "" {
		fmt.Println("input.txt is empty, using example input:")
		runExample()
		return
	}

	fmt.Printf("Part 1 Result: %d\n", solvePart1(input))
	fmt.Printf("Part 2 Result: %d\n", solvePart2(input))
}

func runExample() {
	exampleInput := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	fmt.Printf("Example Part 1: %d\n", solvePart1(exampleInput))
	fmt.Printf("Example Part 2: %d\n", solvePart2(exampleInput))
}
