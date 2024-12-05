package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Open the file

	b, err := os.ReadFile("input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	allMults := extractMulPattern(str)
	var sum int
	for _, value := range allMults {

		sum += processMult(value)
	}

	fmt.Println("day3pt01: ", sum)

	//pt2
	allMultsAndEnable := extractMultAndDoDont(str)
	var shouldProcess = true
	var pt2Sum = 0
	for _, value := range allMultsAndEnable {

		if value == "do()" {
			shouldProcess = true
		} else if value == "don't()" {
			shouldProcess = false
		} else if shouldProcess {
			pt2Sum += processMult(value)
		}
	}
	fmt.Println("day3pt02: ", pt2Sum)

}

func extractMulPattern(input string) []string {
	pattern := regexp.MustCompile("(?i)mul\\([0-9]+,[0-9]+\\)")
	return pattern.FindAllString(input, -1)
}

func extractMultAndDoDont(input string) []string {
	pattern := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
	return pattern.FindAllString(input, -1)

}
func processMult(input string) int {
	//string bullshit, split on comma, convert then  multiply
	halves := strings.Split(input, ",")

	leftNumStr := strings.Trim(strings.Trim((halves[0]), "mul"), "(")
	rightNumStr := strings.Trim(strings.Trim((halves[1]), "mul"), ")")

	leftNum, _ := strconv.Atoi(leftNumStr)
	rightNum, _ := strconv.Atoi(rightNumStr)

	return (leftNum * rightNum)
}

func removeAndShift(slice []int, index int) []int {
	// Create a new slice to avoid modifying the original
	newSlice := make([]int, len(slice)-1)

	// Copy elements before the index
	copy(newSlice, slice[:index])

	// Copy elements after the index
	copy(newSlice[index:], slice[index+1:])

	return newSlice
}

func isReportValid(report []int) bool {
	var prevLevel int = 0
	var prevAcend bool = false
	var validReport bool = false
	for j := 0; j < len(report); j++ {
		var currLevel int = report[j]
		//first level, just set and skip to next level
		if j == 0 {
			prevLevel = currLevel
			if report[j+1] > currLevel {
				prevAcend = true
			} else if report[j+1] < currLevel {
				prevAcend = false
			} else {
				break
			}

			continue
		}

		//equality is not ascending or descending, failure = skip report row
		if prevLevel == currLevel {
			validReport = false
			break
		}

		if isValidLevelDiff(prevLevel, currLevel, prevAcend) {
			validReport = true
			prevLevel = currLevel
			continue
		} else {
			validReport = false
			break
		}
	}

	return validReport
}

// if ascending, must be within 1-3
// if descending, must be within -1-> -3
func isValidLevelDiff(prevLevel int, currLevel int, priorAcend bool) bool {
	var diff int = (currLevel - prevLevel)
	if diff <= 3 && diff >= 1 && priorAcend {
		return true
	} else if diff >= -3 && diff <= -1 && !priorAcend {
		return true
	} else {
		return false
	}

}
