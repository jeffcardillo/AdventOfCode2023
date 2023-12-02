package main

/**
 * Advent of Code 2023
 * Day 2
 * https://adventofcode.com/2023/day/2
 */

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Id    int
	Red   int
	Blue  int
	Green int
}

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

	// for each line in the input text file
	for scanner.Scan() {
		line := scanner.Text()

		// process the line
		gameSlice := parseGameLine(line)

		if partA {
			if isValidGameSet(gameSlice) {
				integerList = append(integerList, gameSlice[0].Id)
			}
		} else {
			integerList = append(integerList, getPowerOfCubes(gameSlice))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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

/**
 * For Part A, find what games are valid if you only allow
 * only 12 red cubes, 14 blue cubes, 13 green cubes
 */
func isValidGameSet(gameSlice []Game) bool {
	for _, value := range gameSlice {
		if value.Red > 12 || value.Blue > 14 || value.Green > 13 {
			return false
		}
	}

	return true
}

/**
 * For Part B, find min number of each color stone
 * to allow the game to be played, and multiply those
 * results together.
 */
func getPowerOfCubes(gameSlice []Game) int {
	minRed := 0
	minGreen := 0
	minBlue := 0

	for _, value := range gameSlice {
		if minRed < value.Red {
			minRed = value.Red
		}
		if minGreen < value.Green {
			minGreen = value.Green
		}
		if minBlue < value.Blue {
			minBlue = value.Blue
		}
	}

	return minRed * minGreen * minBlue
}

func parseGameLine(line string) []Game {

	// split the line at the colon
	strs := strings.Split(line, ":")

	// parse out the id and store it
	gameIdPattern := `(\d+)`
	idRegEx, _ := regexp.Compile(gameIdPattern)
	gameId := idRegEx.FindString(strs[0])

	// break out the reset of the line for each grab
	grabs := strings.Split(strs[1], ";")

	// let's store this data as a list of objects once processed
	gameSlice := []Game{}

	// look at each "grab", a set of red, green, blue stones
	for _, grab := range grabs {
		// create an object for this grab
		var game Game

		// split the string for each color
		cubes := strings.Split(grab, ",")

		// iterate over the different color stones
		for _, pick := range cubes {
			// extract count and color
			countAndColorPattern := `(\d+) ([red|green|blue]+)`
			regex, _ := regexp.Compile(countAndColorPattern)
			match := regex.FindStringSubmatch(pick)
			count, _ := strconv.Atoi(match[1])
			color := match[2]

			// store the data on the game object
			game.Id, _ = strconv.Atoi(gameId)

			switch color {
			case "red":
				game.Red = count
			case "green":
				game.Green = count
			case "blue":
				game.Blue = count
			}
		}

		// add this grab to the list of grabs (games?)
		gameSlice = append(gameSlice, game)
	}

	return gameSlice
}
