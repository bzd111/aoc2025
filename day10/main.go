package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	// Try to read input.txt, if not exists or empty, try example.txt
	inputFile := "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	} else if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		inputFile = "example.txt"
	} else {
		// check if empty
		info, _ := os.Stat(inputFile)
		if info.Size() == 0 {
			inputFile = "example.txt"
		}
	}

	fmt.Printf("Using input file: %s\n", inputFile)
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPresses := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		presses := solveMachine(line)
		if presses == -1 {
			fmt.Println("No solution found for machine!")
		} else {
			totalPresses += presses
		}
	}

	fmt.Println("Total fewest presses (Part 1):", totalPresses)

	// Reset file for Part 2
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	totalPressesPart2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		presses := solveMachinePart2(line)
		if presses >= 0 {
			totalPressesPart2 += presses
		} else {
			// fmt.Println("No solution Part 2")
		}
	}
	fmt.Println("Total fewest presses (Part 2):", totalPressesPart2)
}

func solveMachine(line string) int {
	// Parse line for Part 1 (GF2)
	// ... (Existing logic for extracting diagram/buttons/GF2 target)
	// Actually, careful: Parse logic in original code ignored {...}.
	// And 'target' was from [...]
	// I will just keep solveMachine as is for Part 1.
	// But duplicate parsing logic?
	// Refactor parsing?
	// I'll copy-paste parse logic for now to ensure I don't break Part 1 logic in the snippet I'm replacing?
	// Wait, I am replacing lines 48-248?
	// No, I'll modify main.go carefully.
	// The User provided file has solveMachine at line 51.
	// I will update main() AND add solveMachinePart2.
	// I will overwrite main() and append new functions.

	return solveMachineOriginal(line)
}

func solveMachineOriginal(line string) int {
	// Re-implementing Part 1 Logic since I'm replacing the block
	startBracket := strings.Index(line, "[")
	endBracket := strings.Index(line, "]")
	diagramStr := line[startBracket+1 : endBracket]
	numLights := len(diagramStr)
	target := make([]int, numLights)
	for i, c := range diagramStr {
		if c == '#' {
			target[i] = 1
		} else {
			target[i] = 0
		}
	}

	rest := line[endBracket+1:]
	curlyStart := strings.Index(rest, "{")
	if curlyStart != -1 {
		rest = rest[:curlyStart]
	}

	var buttons [][]int
	parts := strings.Split(rest, ")")
	for _, p := range parts {
		openParen := strings.Index(p, "(")
		if openParen == -1 {
			continue
		}
		content := p[openParen+1:]
		numsStr := strings.Split(content, ",")
		var btn []int
		for _, s := range numsStr {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			var val int
			fmt.Sscanf(s, "%d", &val)
			btn = append(btn, val)
		}
		btnVec := make([]int, numLights)
		for _, idx := range btn {
			if idx >= 0 && idx < numLights {
				btnVec[idx] = 1
			}
		}
		buttons = append(buttons, btnVec)
	}
	return solveSystem(numLights, buttons, target)
}

// Reuse solveSystem from original (GF2)
func solveSystem(n int, buttons [][]int, target []int) int {
	m := len(buttons)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, m+1)
		for j := 0; j < m; j++ {
			mat[i][j] = buttons[j][i]
		}
		mat[i][m] = target[i]
	}
	pivotRow := 0
	pivotCols := make([]int, 0)
	isPivot := make(map[int]bool)
	for col := 0; col < m && pivotRow < n; col++ {
		sel := -1
		for row := pivotRow; row < n; row++ {
			if mat[row][col] == 1 {
				sel = row
				break
			}
		}
		if sel == -1 {
			continue
		}
		mat[pivotRow], mat[sel] = mat[sel], mat[pivotRow]
		for row := 0; row < n; row++ {
			if row != pivotRow && mat[row][col] == 1 {
				for k := col; k <= m; k++ {
					mat[row][k] ^= mat[pivotRow][k]
				}
			}
		}
		pivotCols = append(pivotCols, col)
		isPivot[col] = true
		pivotRow++
	}
	for row := pivotRow; row < n; row++ {
		if mat[row][m] == 1 {
			return -1
		}
	}
	freeVars := make([]int, 0)
	for j := 0; j < m; j++ {
		if !isPivot[j] {
			freeVars = append(freeVars, j)
		}
	}
	minWeight := math.MaxInt32
	count := 1 << len(freeVars)
	for i := 0; i < count; i++ {
		x := make([]int, m)
		currentWeight := 0
		tempI := i
		for k := 0; k < len(freeVars); k++ {
			val := tempI & 1
			tempI >>= 1
			x[freeVars[k]] = val
			if val == 1 {
				currentWeight++
			}
		}
		for r := 0; r < len(pivotCols); r++ {
			col := pivotCols[r]
			val := mat[r][m]
			for _, free := range freeVars {
				if mat[r][free] == 1 {
					val ^= x[free]
				}
			}
			x[col] = val
			if val == 1 {
				currentWeight++
			}
		}
		if currentWeight < minWeight {
			minWeight = currentWeight
		}
	}
	if minWeight == math.MaxInt32 {
		return -1
	}
	return minWeight
}

func solveMachinePart2(line string) int {
	// Parse {...} part
	startBrace := strings.Index(line, "{")
	endBrace := strings.Index(line, "}")
	if startBrace == -1 || endBrace == -1 {
		return -1
	}
	targetStr := line[startBrace+1 : endBrace]
	tParts := strings.Split(targetStr, ",")
	var target []float64
	for _, s := range tParts {
		var v float64
		fmt.Sscanf(strings.TrimSpace(s), "%f", &v)
		target = append(target, v)
	}
	numCounters := len(target)

	// Parse buttons (...) same as before logic
	// Restrict to substring BEFORE {
	prefix := line[:startBrace]
	// Remove [...]
	endBracket := strings.Index(prefix, "]")
	if endBracket != -1 {
		prefix = prefix[endBracket+1:]
	} // buttons part

	var buttons [][]float64
	parts := strings.Split(prefix, ")")
	for _, p := range parts {
		openParen := strings.Index(p, "(")
		if openParen == -1 {
			continue
		}
		content := p[openParen+1:]
		numsStr := strings.Split(content, ",")
		var btn []int
		for _, s := range numsStr {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			var val int
			fmt.Sscanf(s, "%d", &val)
			btn = append(btn, val)
		}
		btnVec := make([]float64, numCounters)
		for _, idx := range btn {
			if idx >= 0 && idx < numCounters {
				btnVec[idx] = 1.0
			}
		}
		buttons = append(buttons, btnVec)
	}

	return solveSystemPart2(numCounters, buttons, target)
}

func solveSystemPart2(n int, buttons [][]float64, target []float64) int {
	m := len(buttons)
	// Matrix n x (m+1)
	mat := make([][]float64, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]float64, m+1)
		for j := 0; j < m; j++ {
			mat[i][j] = buttons[j][i]
		}
		mat[i][m] = target[i]
	}

	pivotRow := 0
	pivotCols := make([]int, 0)
	isPivot := make(map[int]bool)

	// Gaussian Elimination
	for col := 0; col < m && pivotRow < n; col++ {
		sel := -1
		for row := pivotRow; row < n; row++ {
			if math.Abs(mat[row][col]) > 1e-9 {
				sel = row
				break
			}
		}
		if sel == -1 {
			continue
		}

		mat[pivotRow], mat[sel] = mat[sel], mat[pivotRow]

		// Normalize pivot
		pivotVal := mat[pivotRow][col]
		for k := col; k <= m; k++ {
			mat[pivotRow][k] /= pivotVal
		}

		// Eliminate
		for row := 0; row < n; row++ {
			if row != pivotRow && math.Abs(mat[row][col]) > 1e-9 {
				factor := mat[row][col]
				for k := col; k <= m; k++ {
					mat[row][k] -= factor * mat[pivotRow][k]
				}
			}
		}

		pivotCols = append(pivotCols, col)
		isPivot[col] = true
		pivotRow++
	}

	// Consisteny Check
	for row := pivotRow; row < n; row++ {
		if math.Abs(mat[row][m]) > 1e-5 {
			return -1
		} // Inconsistent
	}

	// Identify Free Variables
	freeVars := make([]int, 0)
	for j := 0; j < m; j++ {
		if !isPivot[j] {
			freeVars = append(freeVars, j)
		}
	}

	// Recursive Search for Free Variables
	minTotal := -1

	// Range for free variables?
	// Given typical target values ~50-250 and buttons add >=1, x_i is bounded by Max(target).
	// We'll search up to 300.
	searchLimit := 300

	// Pre-calculate pivot rows and constants for faster checking?
	// x[pivotCol[r]] = mat[r][m] - sum(mat[r][free] * x[free])

	var search func(idx int, currentFree []int)
	search = func(idx int, currentFree []int) {
		if idx == len(freeVars) {
			// Check solution
			total := 0
			// Sum free vars
			for _, val := range currentFree {
				total += val
			}

			// Calculate Pivot Vars
			x := make([]float64, m)
			for k, fIdx := range freeVars {
				x[fIdx] = float64(currentFree[k])
			}

			possible := true
			for r := 0; r < len(pivotCols); r++ {
				col := pivotCols[r]
				val := mat[r][m]
				for _, fIdx := range freeVars {
					if math.Abs(mat[r][fIdx]) > 1e-9 {
						val -= mat[r][fIdx] * x[fIdx]
					}
				}

				// Check constraints
				if val < -1e-5 { // Non-negative constraint
					possible = false
					break
				}
				iv := int(math.Round(val))
				if math.Abs(float64(iv)-val) > 1e-5 { // Integer constraint
					possible = false
					break
				}
				total += iv
				x[col] = float64(iv)
			}

			if possible {
				if minTotal == -1 || total < minTotal {
					minTotal = total
				}
			}
			return
		}

		// Optimize: Pruning?
		// Iterate this free variable
		for val := 0; val <= searchLimit; val++ {
			// Optimization: If current sum of free vars > minTotal (and minTotal != -1), we can prune?
			// But pivot vars can correspond to 'negative' cost in sum if coefficients are weird?
			// But here coefficients are usually positive in A (buttons add).
			// In RREF, pivot = C - K*Free.
			// If K is positive, increasing Free reduces Pivot. Total sum could go down!
			// So we cannot prune based on current free sum strictly.
			search(idx+1, append(currentFree, val))
		}
	}

	search(0, []int{})

	return minTotal
}
