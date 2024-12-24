package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func readInput() (map[string]string, int, int) {
	ex, _ := os.Executable()
	absPath, _ := filepath.Abs(filepath.Join(filepath.Base(ex), "./input.txt"))
	file, _ := os.Open(absPath)
	defer file.Close()

	puzzle := make(map[string]string)
	scanner := bufio.NewScanner(file)
	i := 0
	height, width := 0, 0

	for scanner.Scan() {
		row := scanner.Text()
		width = len(row)
		for j, char := range row {
			puzzle[fmt.Sprintf("%d,%d", i, j)] = string(char)
		}
		i++
	}
	height = i

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return puzzle, height, width
}

type coordFn func(int) int

func searchVector(puzzle map[string]string, term string, yCoord coordFn, xCoord coordFn) bool {
	for k := range term {
		y := yCoord(k)
		x := xCoord(k)
		if puzzle[fmt.Sprintf("%d,%d", y, x)] != string(term[k]) {
			return false
		}
	}
	return true
}

func search(puzzle map[string]string, term string, i, j int) int {
	count := 0
	if puzzle[fmt.Sprintf("%d,%d", i, j)] == string(term[0]) {
		// search right
		if searchVector(puzzle, term, func(k int) int { return i }, func(k int) int { return j + k }) {
			count++
		}
		// search down-right
		if searchVector(puzzle, term, func(k int) int { return i + k }, func(k int) int { return j + k }) {
			count++
		}
		// search down
		if searchVector(puzzle, term, func(k int) int { return i + k }, func(k int) int { return j }) {
			count++
		}
		// search down-left
		if searchVector(puzzle, term, func(k int) int { return i + k }, func(k int) int { return j - k }) {
			count++
		}
		// search left
		if searchVector(puzzle, term, func(k int) int { return i }, func(k int) int { return j - k }) {
			count++
		}
		// search up-left
		if searchVector(puzzle, term, func(k int) int { return i - k }, func(k int) int { return j - k }) {
			count++
		}
		// search up
		if searchVector(puzzle, term, func(k int) int { return i - k }, func(k int) int { return j }) {
			count++
		}
		// search up-right
		if searchVector(puzzle, term, func(k int) int { return i - k }, func(k int) int { return j + k }) {
			count++
		}
	}
	return count
}

func getCorners(puzzle map[string]string, i, j int) []string {
	return []string{
		puzzle[fmt.Sprintf("%d,%d", i-1, j-1)],
		puzzle[fmt.Sprintf("%d,%d", i-1, j+1)],
		puzzle[fmt.Sprintf("%d,%d", i+1, j+1)],
		puzzle[fmt.Sprintf("%d,%d", i+1, j-1)],
	}
}

func xSearch(puzzle map[string]string, i, j int) int {
	count := 0
	if puzzle[fmt.Sprintf("%d,%d", i, j)] == "A" {
		corners := getCorners(puzzle, i, j)
		for _, v := range corners {
			if v == "X" {
				return count
			}
		}

		if (corners[0] == "S" && corners[2] == "M" ||
			corners[0] == "M" && corners[2] == "S") && (corners[1] == "S" && corners[3] == "M" ||
			corners[1] == "M" && corners[3] == "S") {
			count++
		}
	}
	return count
}

func wordCount(puzzle map[string]string, height, width int) int {
	term := "XMAS"
	count := 0
	for i := range height {
		for j := range width {
			count += search(puzzle, term, i, j)
		}
	}
	return count
}

func xCount(puzzle map[string]string, height, width int) int {
	count := 0
	for i := range height {
		for j := range width {
			count += xSearch(puzzle, i, j)
		}
	}
	return count
}

func main() {
	puzzle, height, width := readInput()
	fmt.Printf("puzzle size: %dx%d\n", height, width)
	count := wordCount(puzzle, height, width)
	fmt.Printf("found count: %d\n", count)
	xCount := xCount(puzzle, height, width)
	fmt.Printf("found X count: %d\n", xCount)
}
