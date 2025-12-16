package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solvePart1(input string) int {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	if len(lines) == 0 {
		return 0
	}

	// Determine max line length
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// Identify empty columns
	isEmptyCol := make([]bool, maxLen)
	for i := 0; i < maxLen; i++ {
		isEmptyCol[i] = true
	}

	for _, line := range lines {
		for i, ch := range line {
			if ch != ' ' {
				isEmptyCol[i] = false
			}
		}
	}

	// Define problem ranges (contiguous non-empty columns)
	type Range struct {
		start, end int
		tokens     []string
	}
	var ranges []*Range

	i := 0
	for i < maxLen {
		if isEmptyCol[i] {
			i++
			continue
		}
		start := i
		for i < maxLen && !isEmptyCol[i] {
			i++
		}
		ranges = append(ranges, &Range{start: start, end: i})
	}

	// Collect tokens and assign to ranges
	for _, line := range lines {
		inToken := false
		tokenStart := 0
		tokenValue := ""

		for i, ch := range line {
			if ch != ' ' {
				if !inToken {
					inToken = true
					tokenStart = i
					tokenValue = string(ch)
				} else {
					tokenValue += string(ch)
				}
			} else {
				if inToken {
					// Find range for this token
					for _, r := range ranges {
						if tokenStart >= r.start && tokenStart < r.end {
							r.tokens = append(r.tokens, tokenValue)
							break
						}
					}
					inToken = false
				}
			}
		}
		if inToken {
			for _, r := range ranges {
				if tokenStart >= r.start && tokenStart < r.end {
					r.tokens = append(r.tokens, tokenValue)
					break
				}
			}
		}
	}

	grandTotal := 0
	for _, r := range ranges {
		if len(r.tokens) == 0 {
			continue
		}

		operator := r.tokens[len(r.tokens)-1]
		numbers := r.tokens[:len(r.tokens)-1]

		var nums []int
		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err == nil {
				nums = append(nums, num)
			}
		}

		if len(nums) == 0 {
			continue
		}

		result := nums[0]
		if operator == "+" {
			for i := 1; i < len(nums); i++ {
				result += nums[i]
			}
		} else if operator == "*" {
			for i := 1; i < len(nums); i++ {
				result *= nums[i]
			}
		}
		grandTotal += result
	}

	return grandTotal
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input.txt:", err)
		return
	}

	input := string(content)
	fmt.Printf("Part 1 Result: %d\n", solvePart1(input))
}
