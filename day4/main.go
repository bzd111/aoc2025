package main

import (
	"fmt"
	"os"
	"strings"
)

// Directions for 8 neighbors
var dirs = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func solvePart1(input string) int {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return 0
	}

	rows := len(lines)
	cols := len(lines[0])
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if lines[r][c] != '@' {
				continue
			}

			neighborCount := 0
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
					if nc < len(lines[nr]) && lines[nr][nc] == '@' {
						neighborCount++
					}
				}
			}

			if neighborCount < 4 {
				count++
			}
		}
	}
	return count
}

func solvePart2(input string) int {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return 0
	}

	rows := len(lines)
	// Convert to mutable grid
	grid := make([][]byte, rows)
	for r := 0; r < rows; r++ {
		grid[r] = []byte(lines[r])
	}

	totalRemoved := 0

	for {
		// Identify rolls to remove in this round
		toRemove := [][2]int{}

		for r := 0; r < rows; r++ {
			cols := len(grid[r])
			for c := 0; c < cols; c++ {
				if grid[r][c] != '@' {
					continue
				}

				neighborCount := 0
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < rows && nc >= 0 {
						if nc < len(grid[nr]) && grid[nr][nc] == '@' {
							neighborCount++
						}
					}
				}

				if neighborCount < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		totalRemoved += len(toRemove)
		for _, p := range toRemove {
			grid[p[0]][p[1]] = '.'
		}
	}

	return totalRemoved
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
	exampleInput := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	fmt.Printf("Example Part 1: %d\n", solvePart1(exampleInput))
	fmt.Printf("Example Part 2: %d\n", solvePart2(exampleInput))
}
