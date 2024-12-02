package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	// Create slices to hold the numbers
	var firstNumbers []int
	var secondNumbers []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into two parts
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		if len(parts) != 2 {
			log.Println("Skipping line with invalid format:", line)
			continue
		}

		// Convert the strings to integers
		firstNum, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Println("Error converting first number:", err)
			continue
		}

		secondNum, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Println("Error converting second number:", err)
			continue
		}

		// Append the numbers to their respective lists
		firstNumbers = append(firstNumbers, firstNum)
		secondNumbers = append(secondNumbers, secondNum)
	}

	// Check for errors during reading
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Output the results
	fmt.Println("First numbers:", firstNumbers)
	fmt.Println("Second numbers:", secondNumbers)

	//sort the lists
	sort.Ints(firstNumbers)
	sort.Ints(secondNumbers)

	var distanceSum int = 0
	//Loop and compute distance sum, assumes array are same size.
	for i := 0; i < len(firstNumbers); i++ {
		distanceSum += absDiffInt(firstNumbers[i], secondNumbers[i])
	}

	fmt.Println("day0pt1:", distanceSum)

	//just do a nested loop, could do some cute stuff with a dict/set to count and check each number quicker but who cares
	//Leave them sorted I guess.
	var simiScore int = 0
	for i := 0; i < len(firstNumbers); i++ {
		var instanceCount int = 0
		for j := 0; j < len(secondNumbers); j++ {
			if secondNumbers[j] == firstNumbers[i] {
				instanceCount++
			}
		}
		simiScore += (firstNumbers[i] * instanceCount)
	}
	fmt.Println("day0pt2:", simiScore)

}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
