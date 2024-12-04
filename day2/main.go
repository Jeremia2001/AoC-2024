package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sliceOfSlices, err := readFileToSlice("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resultPartOne := partOne(sliceOfSlices)
	fmt.Println("Answer Part One:", resultPartOne)

	resultPartTwo := partTwo(sliceOfSlices)
	fmt.Println("Answer Part Two:", resultPartTwo)
}

func partOne(sliceOfSlices [][]int) int {
	var safeReports int
	for _, slice := range sliceOfSlices {
		isValid := isFullyValid(slice)

		if isValid {
			safeReports++
		}
	}

	return safeReports
}

func partTwo(sliceOfSlices [][]int) int {
	var safeReports int

	for _, slice := range sliceOfSlices {
		// Check if the slice is already safe without modification
		if isFullyValid(slice) {
			safeReports++
			continue
		}

		// Try removing each level one at a time and check if it becomes safe
		for i := 0; i < len(slice); i++ {
			modifiedSlice := append([]int{}, slice[:i]...)
			modifiedSlice = append(modifiedSlice, slice[i+1:]...)
			if isFullyValid(modifiedSlice) {
				safeReports++
				break
			}
		}
	}

	return safeReports
}

func isSorted(slice []int) (bool, bool) {
	isAscending := true
	isDescending := true

	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			isAscending = false
		}
		if slice[i] > slice[i-1] {
			isDescending = false
		}
	}

	return isAscending, isDescending
}

func isFullyValid(slice []int) bool {
	isAscending, isDescending := isSorted(slice)
	if !isAscending && !isDescending {
		return false
	}

	for i := 0; i < len(slice)-1; i++ {
		difference := slice[i] - slice[i+1]
		if difference < 0 {
			difference *= -1
		}
		if difference < 1 || difference > 3 {
			return false
		}
	}

	return true
}

func readFileToSlice(filename string) ([][]int, error) {
	var result [][]int

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into fields
		line := scanner.Text()
		parts := strings.Fields(line)

		// Convert fields to integers
		var lineSlice []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error parsing number '%s': %v", part, err)
			}
			lineSlice = append(lineSlice, num)
		}

		// Append the line slice to the result
		result = append(result, lineSlice)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
