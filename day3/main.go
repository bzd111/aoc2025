package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solveLine(line string) int {
	line = strings.TrimSpace(line)
	if len(line) < 2 {
		return 0
	}

	maxVal := 0

	for i := 0; i < len(line)-1; i++ {
		for j := i + 1; j < len(line); j++ {
			d1 := int(line[i] - '0')
			d2 := int(line[j] - '0')
			val := d1*10 + d2
			if val > maxVal {
				maxVal = val
			}
		}
	}
	return maxVal
}

func solveLinePart2(line string) int64 {
	line = strings.TrimSpace(line)
	L := len(line)
	required := 12
	if L < required {
		return 0
	}

	currentIdx := 0
	var resultStr strings.Builder

	for k := 1; k <= required; k++ {
		rem := required - k
		// We need to pick a digit from currentIdx such that there are at least 'rem' digits after it.
		// The last valid index is L - 1 - rem.
		maxIdx := L - 1 - rem

		bestDigit := byte('0' - 1) // smaller than '0'
		bestPos := -1

		// Greedy search for the largest digit in the valid range
		for i := currentIdx; i <= maxIdx; i++ {
			if line[i] > bestDigit {
				bestDigit = line[i]
				bestPos = i
				if bestDigit == '9' {
					break // Can't get better than 9
				}
			}
		}

		resultStr.WriteByte(bestDigit)
		currentIdx = bestPos + 1
	}

	val, _ := strconv.ParseInt(resultStr.String(), 10, 64)
	return val
}

func solvePart1(input string) int {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")
	totalSum := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		totalSum += solveLine(line)
	}
	return totalSum
}

func solvePart2(input string) int64 {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")
	var totalSum int64 = 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		totalSum += solveLinePart2(line)
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
	exampleInput := `987654321111111
811111111111119
234234234234278
818181911112111`
	fmt.Printf("Example Part 1: %d\n", solvePart1(exampleInput))
	fmt.Printf("Example Part 2: %d\n", solvePart2(exampleInput))
}
