package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func readInput() string {
	ex, _ := os.Executable()
	absPath, _ := filepath.Abs(filepath.Join(filepath.Base(ex), "./input.txt"))
	file, _ := os.Open(absPath)
	defer file.Close()

	raw := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return raw
}

func trimNonNumeric(str string) int {
	re := regexp.MustCompile(`\D`)
	num, err := strconv.Atoi(re.ReplaceAllString(str, ""))
	if err != nil {
		panic(err)
	}
	return num
}

func parseMulStmts(raw string) []string {
	r := regexp.MustCompile(`mul[(][\d]+,[\d]+[)]`)
	matched := r.FindAllString(raw, -1)
	return matched
}

func filterDoBlocks(raw string) string {
	filtered := make([]string, 0)

	initialSplit := regexp.MustCompile(`don't[(][)]`).Split(raw, -1)
	for i, block := range initialSplit {
		if i == 0 {
			filtered = append(filtered, block)
			continue
		}
		doSplit := regexp.MustCompile(`do[(][)]`).Split(block, -1)
		_, do := doSplit[0], doSplit[1:]
		filtered = slices.Concat(filtered, do)
	}

	return strings.Join(filtered, "")
}

func calcTotal(stmts []string) int {
	total := 0
	for _, stmt := range stmts {
		split := strings.Split(stmt, ",")
		a := trimNonNumeric(split[0])
		b := trimNonNumeric(split[1])
		total += a * b
		// fmt.Printf("%s - a: %d, b: %d, sum: %d\n", stmt, a, b, a*b)
	}
	return total
}

func main() {
	input := readInput()
	fmt.Printf("raw input: %s\n", input)
	filteredStatements := parseMulStmts(input)
	total := calcTotal(filteredStatements)
	doBlockStatements := parseMulStmts(filterDoBlocks(input))
	doBlockTotal := calcTotal(doBlockStatements)
	fmt.Printf("unfiltered total: %d\n", total)
	fmt.Printf("DO block total: %d\n", doBlockTotal)
}
