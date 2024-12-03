package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create slices to hold the reports
	var reports [][]int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into two parts
		line := scanner.Text()
		levels := strings.Split(line, " ")
		// Convert the array of strings to an array of integers
		var report []int
		for _, str := range levels {
			// Convert each string to an integer
			num, err := strconv.Atoi(str)
			if err != nil {
				// Handle error if the string cannot be converted to an integer
				fmt.Println("Error converting string to integer:", err)
				return
			}
			report = append(report, num)
		}
		reports = append(reports, report)

	}

	// Check for errors during reading
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//day 01 - loop through all reports,
	var validReportCount int = 0
	for i := 0; i < len(reports); i++ {
		if isReportValid(reports[i]) {
			validReportCount++
		}

	}

	fmt.Println("day0pt1:", validReportCount)

	var day2validreportcount int = 0
	//pt2 is tough, if a report is bad we need to retry with 1 entry removed
	//lets try the dumb way and brute force by just removing an entry if failed
	for i := 0; i < len(reports); i++ {
		if isReportValid(reports[i]) {
			day2validreportcount++
		} else {
			//remove every entry from the array, one at a time and re-check
			var startingReport []int = reports[i]
			for j := 0; j < len(startingReport); j++ {
				if isReportValid(removeAndShift(startingReport, j)) {
					day2validreportcount++
					break
				} else {
					continue
				}
			}
		}

	}

	fmt.Println("day0pt2:", day2validreportcount)

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
