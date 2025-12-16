package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solvePart1(input string) int {
	adj := make(map[string][]string)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		from := strings.TrimSpace(parts[0])
		destinations := strings.Fields(parts[1])
		adj[from] = destinations
	}

	memo := make(map[string]int)
	return countPaths("you", "out", adj, memo)
}

func countPaths(curr, target string, adj map[string][]string, memo map[string]int) int {
	if curr == target {
		return 1
	}
	if val, ok := memo[curr]; ok {
		return val
	}

	count := 0
	if neighbors, ok := adj[curr]; ok {
		for _, neighbor := range neighbors {
			count += countPaths(neighbor, target, adj, memo)
		}
	}

	memo[curr] = count
	return count
}

func solvePart2(input string) int {
	adj := make(map[string][]string)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		from := strings.TrimSpace(parts[0])
		destinations := strings.Fields(parts[1])
		adj[from] = destinations
	}

	memo := make(map[string]map[int]int)
	return countPathsPart2("svr", "out", adj, 0, memo)
}

// countPathsPart2 counts paths from curr to target that visit both "dac" and "fft".
// visitedMask uses bitmask: 0 (none), 1 (dac), 2 (fft), 3 (both)
func countPathsPart2(curr, target string, adj map[string][]string, visitedMask int, memo map[string]map[int]int) int {
	// Update visitedMask for current node
	if curr == "dac" {
		visitedMask |= 1
	}
	if curr == "fft" {
		visitedMask |= 2
	}

	// If already computed for this (curr, visitedMask) state, return memoized value
	if _, ok := memo[curr]; ok {
		if val, ok := memo[curr][visitedMask]; ok {
			return val
		}
	} else {
		memo[curr] = make(map[int]int)
	}

	// Base case: Reached target
	if curr == target {
		if visitedMask == 3 { // Both dac and fft must have been visited
			return 1
		}
		return 0
	}

	count := 0
	if neighbors, ok := adj[curr]; ok {
		for _, neighbor := range neighbors {
			count += countPathsPart2(neighbor, target, adj, visitedMask, memo)
		}
	}

	memo[curr][visitedMask] = count
	return count
}


func main() {
	filePath := "input.txt"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Could not read %s: %v\n", filePath, err)
		// Fallback to example if standard file missing
		if filePath == "input.txt" {
			fmt.Println("Trying example.txt...")
			content, err = os.ReadFile("example.txt")
			if err != nil {
                 // Try relative path for example if run from root
                 content, err = os.ReadFile("day11/example.txt")
                 if err != nil {
                     fmt.Println("Could not read example input.")
                     return
                 }
			}
		} else {
            return
        }
	}

	input := string(content)
	fmt.Printf("Part 1 Result: %d\n", solvePart1(input))
	fmt.Printf("Part 2 Result: %d\n", solvePart2(input))
}
