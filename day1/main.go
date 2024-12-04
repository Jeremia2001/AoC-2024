package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	leftData, rightData, err := readFileAndSplit("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	slices.Sort(leftData)
	slices.Sort(rightData)

	resultPartOne := partOne(leftData, rightData)
	fmt.Println("Answer Part One:", resultPartOne)

	resultPartTwo := partTwo(leftData, rightData)
	fmt.Println("Answer Part Two:", resultPartTwo)
}

func partTwo(leftData []int, rightData []int) int {
	var totalSimilarityScore int
	for _, value := range leftData {
		occurrences := findOccurrences(rightData, value)
		similarityScore := value * occurrences
		totalSimilarityScore += similarityScore
	}

	return totalSimilarityScore
}

func partOne(leftData []int, rightData []int) int {
	var totalDifference int
	for i, _ := range leftData {
		diff := leftData[i] - rightData[i]
		if diff < 0 {
			diff *= -1
		}

		totalDifference += diff
	}

	return totalDifference
}

func findOccurrences(data []int, valueToLookFor int) int {
	var occurrences int
	for _, value := range data {
		if value == valueToLookFor {
			occurrences++
		}
	}

	return occurrences
}

func readFileAndSplit(filename string) ([]int, []int, error) {
	var leftData []int
	var rightData []int

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split each line into parts
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		// Parse the numbers
		left, err1 := strconv.Atoi(parts[0])
		right, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("error parsing line: %s", line)
		}

		// Append to slices
		leftData = append(leftData, left)
		rightData = append(rightData, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftData, rightData, nil
}
