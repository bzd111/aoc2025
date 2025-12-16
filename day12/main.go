package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Shape struct {
	id    int
	w, h  int
	cells [][]bool
	area  int
}

type Task struct {
	w, h   int
	counts []int
}

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	shapes := []Shape{}
	tasks := []Task{}

	// Parse Shapes
	// Shapes 0-5.
	// Format: "check for ID:"

	// Quick parse
	idx := 0
	for idx < len(lines) {
		line := strings.TrimSpace(lines[idx])
		if line == "" {
			idx++
			continue
		}

		if strings.HasSuffix(line, ":") {
			// ID check
			// idStr := strings.TrimSuffix(line, ":")
			// Read grid
			idx++
			var grid []string
			for idx < len(lines) {
				l := strings.TrimSpace(lines[idx])
				if l == "" {
					break
				}
				if strings.Contains(l, "x") {
					break
				} // Task line start?
				// Actually tasks start with WxH:
				// So if it contains ':', it's likely a task line or Shape ID.
				// Shape grids don't have ':'.
				// Tasks lines have "40x40: ..."
				if strings.Contains(l, ":") && !strings.HasSuffix(l, ":") {
					break
				} // Task line
				if strings.HasSuffix(l, ":") {
					break
				} // Next shape
				grid = append(grid, l)
				idx++
			}
			// Create shape
			h := len(grid)
			if h > 0 {
				w := len(grid[0])
				cells := make([][]bool, h)
				area := 0
				for r := 0; r < h; r++ {
					cells[r] = make([]bool, w)
					for c := 0; c < w; c++ {
						if c < len(grid[r]) && grid[r][c] == '#' {
							cells[r][c] = true
							area++
						}
					}
				}
				shapes = append(shapes, Shape{len(shapes), w, h, cells, area})
			}
		} else if strings.Contains(line, "x") && strings.Contains(line, ":") {
			// Task line "50x44: 49 45 ..."
			parts := strings.Split(line, ":")
			dims := strings.Split(parts[0], "x")
			W, _ := strconv.Atoi(dims[0])
			H, _ := strconv.Atoi(dims[1])

			cntsStr := strings.Fields(parts[1])
			var cnts []int
			for _, s := range cntsStr {
				v, _ := strconv.Atoi(s)
				cnts = append(cnts, v)
			}
			tasks = append(tasks, Task{W, H, cnts})
			idx++
		} else {
			idx++
		}
	}

	fmt.Printf("Parsed %d shapes\n", len(shapes))
	for _, s := range shapes {
		fmt.Printf("Shape %d: %dx%d Area %d\n", s.id, s.w, s.h, s.area)
	}

	validArea := 0
	total := 0
	minUtil := 1.0
	maxUtil := 0.0

	for _, t := range tasks {
		reqArea := 0
		for i, c := range t.counts {
			if i < len(shapes) {
				reqArea += c * shapes[i].area
			}
		}

		regArea := t.w * t.h
		util := float64(reqArea) / float64(regArea)

		if util < minUtil {
			minUtil = util
		}
		if util > maxUtil {
			maxUtil = util
		}

		fits := reqArea <= regArea
		if fits {
			validArea++
		}
		total++
		fmt.Printf("Task %dx%d: Req %d / %d (%.2f%%) Fits? %v\n", t.w, t.h, reqArea, regArea, util*100, fits)
	}

	fmt.Printf("Total Tasks: %d\n", total)
	fmt.Printf("Valid Area: %d\n", validArea)

	maxValidUtil := 0.0
	for _, t := range tasks {
		reqArea := 0
		for i, c := range t.counts {
			if i < len(shapes) {
				reqArea += c * shapes[i].area
			}
		}
		regArea := t.w * t.h
		util := float64(reqArea) / float64(regArea)
		if reqArea <= regArea {
			if util > maxValidUtil {
				maxValidUtil = util
			}
		}
	}
	fmt.Printf("Max Utilization of Valid Area Tasks: %.2f%%\n", maxValidUtil*100)
	fmt.Println("Day 12 Result")
	fmt.Printf("Part 1: %d\n", validArea)
}
