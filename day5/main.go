package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func parseInput(input string) ([]Range, []int) {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\r\n", "\n")
	parts := strings.Split(input, "\n\n")

	if len(parts) < 2 {
		return nil, nil
	}

	rangePart := parts[0]
	idPart := parts[1]

	var ranges []Range
	// Ranges are space separated tokens like "3-5"
	rangeTokens := strings.Fields(rangePart)
	for _, t := range rangeTokens {
		nums := strings.Split(t, "-")
		if len(nums) == 2 {
			s, _ := strconv.Atoi(nums[0])
			e, _ := strconv.Atoi(nums[1])
			ranges = append(ranges, Range{Start: s, End: e})
		}
	}

	var ids []int
	idTokens := strings.Fields(idPart)
	for _, t := range idTokens {
		val, err := strconv.Atoi(t)
		if err == nil {
			ids = append(ids, val)
		}
	}

	return ranges, ids
}

func solvePart1(input string) int {
	ranges, ids := parseInput(input)
	count := 0

	for _, id := range ids {
		isFresh := false
		for _, r := range ranges {
			if id >= r.Start && id <= r.End {
				isFresh = true
				break
			}
		}
		if isFresh {
			count++
		}
	}
	return count
}

func solvePart2(input string) int {
	ranges, _ := parseInput(input)
	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by Start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	var merged []Range
	if len(ranges) > 0 {
		merged = append(merged, ranges[0])
	}

	for i := 1; i < len(ranges); i++ {
		cur := ranges[i]
		last := &merged[len(merged)-1]

		// Check overlap or contiguous (Start <= End + 1)
		// e.g. [1,5] and [6,10] can merge to [1,10] because 6 <= 5+1
		if cur.Start <= last.End+1 {
			if cur.End > last.End {
				last.End = cur.End
			}
		} else {
			merged = append(merged, cur)
		}
	}

	count := 0
	for _, r := range merged {
		count += (r.End - r.Start + 1)
	}
	return count
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
	exampleInput := `3-5 10-14 16-20 12-18

	1 5 8 11 17 32`
	// Note: Example input in problem description was slightly compressed,
	// Assuming standard \n\n separation for sections as described "a blank line".

	fmt.Printf("Example Result Part 1: %d\n", solvePart1(exampleInput))
	fmt.Printf("Example Result Part 2: %d\n", solvePart2(exampleInput))
}
