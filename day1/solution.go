package main

import (
	lib "advent-of-code/lib"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func getLists() ([]int64, []int64) {
	ex, _ := os.Executable()
	absPath, _ := filepath.Abs(filepath.Join(filepath.Base(ex), "./input.txt"))
	file, _ := os.Open(absPath)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	records, _ := reader.ReadAll()

	var a, b []int64
	for _, record := range records {
		if len(record) != 2 {
			break
		}
		forA, _ := strconv.ParseInt(record[0], 0, 64)
		forB, _ := strconv.ParseInt(record[1], 0, 64)

		a = append(a, forA)
		b = append(b, forB)
	}

	return a, b
}

func findMinIndex(s []int64) int {
	index := 0
	min := s[index]
	for i, num := range s {
		if num < min {
			index = i
			min = s[index]
		}
	}
	return index
}

func calcTotalDistance(a []int64, b []int64) int64 {
	var dist int64
	size := len(a)
	for i := 0; i < size; i++ {
		minA := findMinIndex(a)
		minB := findMinIndex(b)
		dist += int64(math.Abs(float64(a[minA] - b[minB])))
		a = lib.RemoveIndex(a, minA)
		b = lib.RemoveIndex(b, minB)
	}
	return dist
}

func getAppearanceCount(s []int64, val int64) int {
	count := 0
	for _, num := range s {
		if num == val {
			count++
		}
	}
	return count
}

func getSimilarity(a []int64, b []int64) int64 {
	var similarity int64
	size := len(a)
	for i := 0; i < size; i++ {
		value := a[i]
		appearances := getAppearanceCount(b, value)
		similarity += value * int64(appearances)
	}
	return similarity
}

func main() {
	a, b := getLists()
	distance := calcTotalDistance(a, b)
	similarity := getSimilarity(a, b)
	fmt.Println("total distance: ", distance)
	fmt.Println("similarity: ", similarity)
}
