package main

/**
 * Advent of Code 2023
 * Day 1
 * https://adventofcode.com/2023/day/1
 */

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var integerList []int

	for scanner.Scan() {
		line := scanner.Text()
		if partA {
			integerList = append(integerList, getNumberFromLinePartA(line))
		} else {
			integerList = append(integerList, getNumberFromLinePartB(line))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return integerList
}

func sumList(list []int) int {
	result := 0

	for i := 0; i < len(list); i++ {
		result += list[i]
	}

	return result
}

func getNumberFromLinePartA(line string) int {
	firstNum := "-1"
	secondNum := "-1"

	for _, char := range line {

		if unicode.IsDigit(char) {
			if firstNum == "-1" {
				firstNum = string(char)
			} else {
				secondNum = string(char)
			}
		}
	}

	resultStr := firstNum

	if secondNum != "-1" {
		resultStr = resultStr + secondNum
	} else {
		resultStr = resultStr + firstNum
	}

	num, _ := strconv.Atoi(resultStr)

	return num
}

func getNumberFromLinePartB(line string) int {

	firstNum := "-1"
	secondNum := "-1"

	for i, char := range line {

		foundNumber := ""

		if unicode.IsDigit(char) {
			foundNumber = string(char)
		} else {
			subString := line[i:]
			foundNumber = getWordNumber(subString)
		}

		if foundNumber != "" {
			if firstNum == "-1" {
				firstNum = string(foundNumber)
			} else {
				secondNum = string(foundNumber)
			}
		}
	}

	resultStr := firstNum

	if secondNum != "-1" {
		resultStr = resultStr + secondNum
	} else {
		resultStr = resultStr + firstNum
	}

	num, _ := strconv.Atoi(resultStr)

	return num
}

func getWordNumber(line string) string {
	// Go doesn't support enums, let's use a map
	numberMap := make(map[int]string)
	numberMap[1] = "one"
	numberMap[2] = "two"
	numberMap[3] = "three"
	numberMap[4] = "four"
	numberMap[5] = "five"
	numberMap[6] = "six"
	numberMap[7] = "seven"
	numberMap[8] = "eight"
	numberMap[9] = "nine"

	for i := 1; i <= 9; i++ {
		if strings.HasPrefix(line, numberMap[i]) {
			return strconv.Itoa(i)
		}
	}

	return ""
}
