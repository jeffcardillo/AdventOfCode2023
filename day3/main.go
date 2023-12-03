package main

/**
 * Advent of Code 2023
 * Day 3
 * https://adventofcode.com/2023/day/3
 */

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	partA()
	partB()
}

func partA() {
	integerList := readFile(true)
	fmt.Println("Part A: " + strconv.Itoa(sumList(integerList)))
}

func partB() {
	integerList := readFile(false)
	fmt.Println("Part B: " + strconv.Itoa(sumList(integerList)))
}

func readFile(partA bool) []int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lineList []string

	// for each line in the input text file
	for scanner.Scan() {
		lineList = append(lineList, scanner.Text())
	}

	lineSymbolMap := make(map[int][]int)
	var integerList []int

	if partA {
		// for each line in input file, discover all symbols, store in a map
		// (line number -> list of symbol locations)
		for index, line := range lineList {
			symbolIndexForLine := parseSymbolIndexsOnLine(line, `[^0-9.\s]`)
			lineSymbolMap[index] = symbolIndexForLine
		}

		// for each line, find numbers that are near the symbols
		for index, line := range lineList {
			res := findNumbersNearSymbols(line, index, lineSymbolMap)
			integerList = append(integerList, res...)
		}
	} else {
		// for each line in input file
		for index, line := range lineList {
			// discover all * symbols
			symbolIndexForLine := parseSymbolIndexsOnLine(line, `[*]`)

			// for each * symbol discovered on the current line
			for _, symbolIndex := range symbolIndexForLine {
				// find adjacent numbers and get multiplied value
				res := getGearRatioForSymbol(index, symbolIndex, lineList)

				if res != -1 {
					integerList = append(integerList, res)
				}
			}
		}
	}

	return integerList
}

/**
 * Sums all of the ints in a list
 */
func sumList(list []int) int {
	result := 0

	for i := 0; i < len(list); i++ {
		result += list[i]
	}

	return result
}

func parseSymbolIndexsOnLine(line string, pattern string) []int {
	var syms []int
	regEx, _ := regexp.Compile(pattern)
	matches := regEx.FindAllStringIndex(line, -1)

	for _, match := range matches {
		syms = append(syms, match[0]) // symbol is single character, just capture start index
	}

	return syms
}

func findNumbersNearSymbols(line string, lineIndex int, symbolIndexes map[int][]int) []int {
	// store all matched numbers
	var matchedNumbers []int

	// pattern to match numbers
	numPattern := `\d+`
	regEx, _ := regexp.Compile(numPattern)

	// extract all numbers from the line
	matches := regEx.FindAllStringIndex(line, -1)

	// for each number found
	for _, match := range matches {
		// get the start and end index for the number
		start := match[0]
		end := match[1] - 1
		number, _ := strconv.Atoi(line[start : end+1])

		// modify start and end bounds for symbol search, careful not to exceed valid boundaries
		start = int(math.Abs(float64(start) - 1))
		end = min(end+1, len(line)-1)

		matched := false

		// check line above and below the current line
		for i := lineIndex - 1; i <= lineIndex+1; i++ {
			// ensure valid line index
			if i >= 0 {
				// check if any symbols on the line are in range of this number
				if !matched && symbolsInRange(start, end, symbolIndexes[i]) {
					matched = true
				}
			}
		}

		if matched {
			// add this number to the results
			matchedNumbers = append(matchedNumbers, number)
		}
	}

	return matchedNumbers
}

func symbolsInRange(start int, end int, symbols []int) bool {
	// if a symbol in the list is within range then return true, else false
	for _, location := range symbols {
		if location >= start && location <= end {
			return true
		}
	}

	return false
}

func getGearRatioForSymbol(lineIndex int, symbolIndex int, lines []string) int {
	var matchedNumbers []int

	// extract numbers from the line
	numPattern := `\d+`
	regEx, _ := regexp.Compile(numPattern)

	// check line above and below the current line
	for i := lineIndex - 1; i <= lineIndex+1; i++ {
		// ensure valid index
		if i >= 0 {
			// get numbers on the current line
			matches := regEx.FindAllStringIndex(lines[i], -1)

			// for each number found
			for _, match := range matches {
				start := match[0]
				end := match[1] - 1
				number, _ := strconv.Atoi(lines[i][start : end+1])

				// stretch the start and end index by 1
				start = int(math.Abs(float64(start) - 1))
				end = min(end+1, len(lines[i])-1)

				// is the current symbol adjacent to this number?
				if symbolsInRange(start, end, []int{symbolIndex}) {
					matchedNumbers = append(matchedNumbers, number)
				}
			}
		}
	}

	// if we have 2 numbers adjacent to the symbol, return them multiplied together, else -1
	if len(matchedNumbers) == 2 {
		return matchedNumbers[0] * matchedNumbers[1]
	} else {
		return -1
	}
}

func min(left int, right int) int {
	if left < right {
		return left
	} else {
		return right
	}
}
