package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
	id      int
}

type Pair struct {
	u, v   int
	distSq int
}

func main() {
	inputFile := "input.txt"
	if len(os.Args) >= 2 {
		inputFile = os.Args[1]
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		content, err = os.ReadFile("day8/" + inputFile)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	points := parseInput(string(content))

	connections := 1000
	if len(points) < 100 {
		connections = 10
	}

	fmt.Println("Part 1:", solve_part1(points, connections))
	fmt.Println("Part 2:", solve_part2(points))
}

func parseInput(input string) []Point {
	var points []Point
	lines := strings.Split(input, "\n")
	id := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, Point{x, y, z, id})
		id++
	}
	return points
}

func solve_part1(points []Point, connections int) int {
	// 1. Generate Pairs
	var pairs []Pair
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z
			distSq := dx*dx + dy*dy + dz*dz
			pairs = append(pairs, Pair{points[i].id, points[j].id, distSq})
		}
	}

	// 2. Sort Pairs
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distSq < pairs[j].distSq
	})

	// 3. Union-Find
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	var find func(int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	union := func(i, j int) {
		rootI := find(i)
		rootJ := find(j)
		if rootI != rootJ {
			parent[rootJ] = rootI
			size[rootI] += size[rootJ]
		}
	}

	// Connect top K pairs
	count := 0
	for _, p := range pairs {
		if count >= connections {
			break
		}
		// "Connect" implies adding the edge.
		// "After making the ten shortest connections"
		// Does this mean 10 *successful* connections (merges)?
		// Or 10 wires used?
		// "Process continues... connect together the 1000 pairs... which are closest"
		// "Afterward..."
		// "Because these two... were already in the same circuit, nothing happens!"
		// This implies we iterate through the top K distinct pairs.
		// Even if they are already connected, we "use" a cable.
		// So we strictly process the first K items in sorted list.

		union(p.u, p.v)
		count++
	}

	// 4. Analyze Components
	// Find root for all valid nodes
	// Since parent array tracks roots.
	// But we only care about component sizes.
	// size array at roots is valid.

	// Collect valid sizes
	var sizes []int
	seenRoots := make(map[int]bool)
	for i := 0; i < n; i++ {
		root := find(i)
		if !seenRoots[root] {
			sizes = append(sizes, size[root])
			seenRoots[root] = true
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j] // Descending
	})

	if len(sizes) < 3 {
		return 0 // Should not happen
	}

	return sizes[0] * sizes[1] * sizes[2]
}

func solve_part2(points []Point) int {
	// 1. Generate Pairs
	var pairs []Pair
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z
			distSq := dx*dx + dy*dy + dz*dz
			pairs = append(pairs, Pair{points[i].id, points[j].id, distSq})
		}
	}

	// 2. Sort Pairs
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distSq < pairs[j].distSq
	})

	// 3. Union-Find until 1 component remains
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	components := n

	var find func(int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	for _, p := range pairs {
		rootU := find(p.u)
		rootV := find(p.v)

		if rootU != rootV {
			parent[rootV] = rootU
			components--

			// Check if this was the last merge needed
			if components == 1 {
				// Result = Product of X coordinates
				return points[p.u].x * points[p.v].x
			}
		}
	}

	return 0
}
