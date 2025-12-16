package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Candidate struct {
	area   int
	p1, p2 Point
	x1, x2 int
	y1, y2 int
}

func main() {
	inputFile := "input.txt"
	if len(os.Args) >= 2 {
		inputFile = os.Args[1]
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		// Fallback for running from root
		content, err = os.ReadFile("day9/" + inputFile)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	points := parseInput(string(content))
	fmt.Println("Part 1:", solve_part1(points))
	fmt.Println("Part 2:", solve_part2(points))
}

func parseInput(input string) []Point {
	var points []Point
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{x, y})
	}
	return points
}

func solve_part1(points []Point) int {
	maxArea := 0
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := points[i]
			p2 := points[j]

			width := int(math.Abs(float64(p1.x-p2.x))) + 1
			height := int(math.Abs(float64(p1.y-p2.y))) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func solve_part2(points []Point) int {
	var candidates []Candidate
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := points[i]
			p2 := points[j]

			// Normalize bounds
			x1, x2 := p1.x, p2.x
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			y1, y2 := p1.y, p2.y
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			width := int(math.Abs(float64(p1.x-p2.x))) + 1
			height := int(math.Abs(float64(p1.y-p2.y))) + 1
			area := width * height

			candidates = append(candidates, Candidate{area, p1, p2, x1, x2, y1, y2})
		}
	}

	// Sort Descending
	// We need 'sort' package imported in Part 1? Yes.
	// But slice sort requires sort.Slice.
	// Need to check imports.
	// Assuming sort is imported or I need to add it.
	// Part 1 snippet had "strings", "fmt", "math", "os", "strconv". Missing "sort".
	// I should verify imports area.

	// Helper for checking inside
	isInside := func(c Candidate) bool {
		// 1. Check for Vertex Intrusion
		// strictly inside: x1 < x < x2 AND y1 < y < y2
		for _, p := range points {
			if p.x > c.x1 && p.x < c.x2 && p.y > c.y1 && p.y < c.y2 {
				return false
			}
		}

		// 2. Check Center Inclusion (Ray Casting)
		// Midpoint
		midX := float64(c.x1+c.x2) / 2.0
		midY := float64(c.y1+c.y2) / 2.0

		intersections := 0
		for k := 0; k < n; k++ {
			pA := points[k]
			pB := points[(k+1)%n]

			// Check Edge Intersection with Rectangle Interior
			// Vertical Edge
			if pA.x == pB.x {
				// Edge X must be strictly between Rect X
				if pA.x > c.x1 && pA.x < c.x2 {
					// Edge Y range must overlap Rect Y strict range
					minEy, maxEy := pA.y, pB.y
					if minEy > maxEy {
						minEy, maxEy = maxEy, minEy
					}

					// Interval Overlap (minEy, maxEy) with (c.y1, c.y2)
					overlapStart := math.Max(float64(minEy), float64(c.y1))
					overlapEnd := math.Min(float64(maxEy), float64(c.y2))

					// Strictly overlapping logic:
					// The interior of the edge intersects the interior of the rect.
					// Since integer coord edges...
					// A vertical edge at X passes through y-range.
					// If the overlapping y-range > 0 length, it cuts the rect.
					if overlapStart < overlapEnd {
						return false
					}
				}

				// Ray Casting Logic (Vertical Edges Only)
				// Ray: y = midY, x > midX
				// Edge must straddle midY
				minEy, maxEy := float64(pA.y), float64(pB.y)
				if minEy > maxEy {
					minEy, maxEy = maxEy, minEy
				}

				if pA.x > int(midX) { // Strictly right of point
					if midY > minEy && midY < maxEy {
						intersections++
					}
				}
			} else {
				// Horizontal Edge
				// Check Intersection with Rect Interior
				minEx, maxEx := pA.x, pB.x
				if minEx > maxEx {
					minEx, maxEx = maxEx, minEx
				}

				if pA.y > c.y1 && pA.y < c.y2 {
					// Edge Y strictly inside Rect Y
					// Check X overlap
					overlapStart := math.Max(float64(minEx), float64(c.x1))
					overlapEnd := math.Min(float64(maxEx), float64(c.x2))
					if overlapStart < overlapEnd {
						return false
					}
				}
			}
		}

		return intersections%2 != 0
	}

	// Sort logic (Bubble sort if sort pkg missing? No, I'll update imports)
	// I'll manually implement simple sort or rely on replacement having imports.
	// Actually I will update imports in a separate block if needed.
	// But I can't do multiple discontinuous edits easily.
	// I will just use a simple selection sort since N^2 is 125000, sorting might be slow with O(M^2).
	// QuickSort is needed.
	// I will assume I can fix main.go fully.
	// Or I can use my own sort function.

	// Let's implement a inplace QuickSort to avoid import dependency if possible,
	// or assume I will fix imports. I'll fix imports.

	quicksort(candidates)

	for _, c := range candidates {
		if isInside(c) {
			return c.area
		}
	}

	return 0
}

func quicksort(arr []Candidate) {
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1
	pivot := arr[len(arr)/2].area
	i, j := left, right
	for i <= j {
		for arr[i].area > pivot {
			i++
		} // Descending
		for arr[j].area < pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if left < j {
		quicksort(arr[:j+1])
	}
	if i < right {
		quicksort(arr[i:])
	}
}
