# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is an Advent of Code 2025 repository. Note that 2025 has **12 days of puzzles** (not the traditional 25 days).

## Repository Structure

Each day follows a consistent pattern:
- `dayx` - go solution directory (e.g., `day01`, `day02`, ..., `day12`)
- `dayx/input.txt` - input.txt with the quesiton(e.g., `day01/input.txt`)
- `dayx/go.mod` - go.mod file (e.g., `day01/go.mod`)
- `dayx/main.go` - main.go file (e.g., `day01/main.go`)

## Running Solutions

To run a solution:
```bash
go run dayx/main.go
```

This will print both Part 1 and Part 2 results.

## Code Structure

Each day's  go file follows this pattern:

```go
func solve_part1(input_text string) string {
    // Part 1 solution
    return result
}

func solve_part2(input_text string) string {
    // Part 2 solution
    return result
}

func main() {
    fmt.Println("Part 1:", solve_part1(input_text))
    fmt.Println("Part 2:", solve_part2(input_text))
}

## Fetching Puzzle Inputs

Inputs are user-specific and require authentication. Use the `web-browser` skill (available in `.claude/skills/web-browser`) to fetch them:

1. Launch the skill: `skill: "web-browser"`
2. Start Chrome: `./tools/start.js`
3. Navigate to the input URL: `./tools/nav.js https://adventofcode.com/2025/day/X/input`
4. Extract the text content: `./tools/eval.js 'document.body.innerText'`
5. Save to file using bash redirection or the Write tool

Example workflow:
```bash
cd .claude/skills/web-browser && ./tools/start.js
# Wait a moment for Chrome to start
cd .claude/skills/web-browser && ./tools/nav.js https://adventofcode.com/2025/day/2/input
cd .claude/skills/web-browser && ./tools/eval.js 'document.body.innerText' > inputs/day02.txt
```

Note: The user's session cookies are needed to fetch inputs, so the web-browser approach is required rather than simple curl commands.

## Puzzle URLs

- Main calendar: https://adventofcode.com/2025
- Day N problem: https://adventofcode.com/2025/day/N
- Day N input: https://adventofcode.com/2025/day/N/input

## Important Notes

- This year has only 12 days of puzzles (days 1-12)
- Inputs are personalized per user and cannot be shared
- Each puzzle has two parts (Part 1 and Part 2)
- Part 2 typically unlocks after completing Part 1