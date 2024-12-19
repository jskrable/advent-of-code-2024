package main

import (
	lib "advent-of-code/lib"
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readInput() [][]string {
	ex, _ := os.Executable()
	absPath, _ := filepath.Abs(filepath.Join(filepath.Base(ex), "./input.txt"))
	file, _ := os.Open(absPath)
	defer file.Close()

	records := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		records = append(records, strings.Split(scanner.Text(), " "))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return records
}

func evaluateReport(report []string) bool {
	volatile := false
	inconsistent := false
	for i, value := range report {
		dblPrevValue := -1
		previousValue := -1
		if i >= 2 {
			dblPrevValue, _ = strconv.Atoi(report[i-2])
		}
		if i >= 1 {
			previousValue, _ = strconv.Atoi(report[i-1])
		}
		currentValue, _ := strconv.Atoi(value)
		if previousValue == -1 {
			continue
		}
		if math.Abs(float64(previousValue)-float64(currentValue)) > 3 {
			volatile = true
			break
		}
		if previousValue == currentValue {
			inconsistent = true
			break
		}
		if dblPrevValue == -1 {
			continue
		}
		if dblPrevValue < previousValue && previousValue > currentValue {
			inconsistent = true
			break
		}
		if dblPrevValue > previousValue && previousValue < currentValue {
			inconsistent = true
			break
		}
	}
	fmt.Printf("%s - volatile: %t, inconsistent: %t, safe: %t\n", report, volatile, inconsistent, !volatile && !inconsistent)
	return !volatile && !inconsistent
}

func problemDampener(report []string) bool {
	saved := false
	for i := range report {
		canDampen := evaluateReport(lib.RemoveIndex(report, i))
		if canDampen {
			saved = true
			break
		}
	}
	return saved

}

func checkReports(records [][]string) int {
	safeCount := 0
	for _, report := range records {
		safe := evaluateReport(report)
		if safe {
			safeCount++
		} else {
			dampened := problemDampener(report)
			if dampened {
				safeCount++
			}
		}
	}
	return safeCount
}

func main() {
	records := readInput()
	fmt.Println(records)
	safety := checkReports(records)
	fmt.Println("safety: ", safety)
}
