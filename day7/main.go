package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Point struct for coordinates
type Point struct {
	x, y int
}

func solve_part1(input_text string) string {
	grid, start := parseInputString(input_text)

	// Set Logic: Count unique splitters hit
	// This yields 1546 for the input, as confirmed by the user.
	// This implies that beams merge fully when they meet at a cell.

	height := len(grid)
	if height == 0 {
		return "0"
	}
	width := len(grid[0])

	// Map of x-col -> exists
	beams := make(map[int]bool)
	beams[start.x] = true

	totalHits := 0

	for y := start.y; y < height; y++ {
		next := make(map[int]bool)
		for x := range beams {
			if x < 0 || x >= width {
				continue
			}

			if grid[y][x] == '^' {
				totalHits++
				next[x-1] = true
				next[x+1] = true
			} else {
				// conduit (.)
				next[x] = true
			}
		}
		beams = next
	}

	return strconv.Itoa(totalHits)
}

func solve_part2(input_text string) string {
	grid, start := parseInputString(input_text)

	height := len(grid)
	if height == 0 {
		return "0"
	}
	width := len(grid[0])

	// Map of x-col -> count of timelines
	beams := make(map[int]int)
	beams[start.x] = 1

	deadTimelines := 0

	for y := start.y; y < height; y++ {
		next := make(map[int]int)
		for x, count := range beams {
			// Check Bounds: if out of bounds, it's a "dead" timeline but valid path?
			// But loop iterates valid Y.
			// Beam position X was valid at Y.
			// It moves to Y+1.
			// If at Y, cell is ^. Targets X-1, X+1.
			// If X-1 < 0. That branch exits.
			// We must count it.

			if x < 0 || x >= width {
				// This shouldn't happen if we prune before adding to map?
				// My previous code: if x < 0 ... continue. (Pruned).
				// So beams map only contained valid.
				continue
			}

			if grid[y][x] == '^' {
				// Split into 2 timelines
				// Left Branch
				if x-1 < 0 || x-1 >= width {
					deadTimelines += count
				} else {
					next[x-1] += count
				}

				// Right Branch
				if x+1 < 0 || x+1 >= width {
					deadTimelines += count
				} else {
					next[x+1] += count
				}
			} else {
				// Continue timeline (.)
				// x stays same.
				next[x] += count
			}
		}
		beams = next
	}

	totalTimelines := 0
	for _, count := range beams {
		totalTimelines += count
	}
	// Add side (dead) timelines
	totalTimelines += deadTimelines

	return strconv.Itoa(totalTimelines)
}

func main() {
	inputFile := "input.txt"
	if len(os.Args) >= 2 {
		inputFile = os.Args[1]
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		// Fallback for running from root
		content, err = os.ReadFile("day7/" + inputFile)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}
	input_text := string(content)

	fmt.Println("Part 1:", solve_part1(input_text))
	fmt.Println("Part 2:", solve_part2(input_text))
}

func parseInputString(input string) ([][]rune, Point) {
	var grid [][]rune
	var start Point

	lines := strings.Split(input, "\n")
	y := 0
	for _, rawLine := range lines {
		text := strings.TrimSpace(rawLine)
		if len(text) == 0 {
			continue
		}
		line := []rune(text)
		grid = append(grid, line)
		for x, char := range line {
			if char == 'S' {
				start = Point{x, y}
			}
		}
		y++
	}
	return grid, start
}
