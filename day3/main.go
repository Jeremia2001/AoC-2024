package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	patternMul    = `mul\(\d+,\d+\)`
	patternDo     = `do\(\)`
	patternDont   = `don't\(\)`
	patternDigits = `\d+,\d+`
)

func main() {
	stringContent, err := readFileToString("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	resultPartOne, err := partOne(stringContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Answer Part One:", resultPartOne)

	resultPartTwo, err := partTwo(stringContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Answer Part Two:", resultPartTwo)
}

func partTwo(input string) (int, error) {
	mulRegex := regexp.MustCompile(patternMul)
	doRegex := regexp.MustCompile(patternDo)
	dontRegex := regexp.MustCompile(patternDont)

	// Extract all matches for `do()`, `don't()`, and `mul()`
	instructions := regexp.MustCompile(fmt.Sprintf(`%s|%s|%s`, patternMul, patternDo, patternDont)).FindAllString(input, -1)

	isEnabled := true
	resultOfAllMultiplications := 0

	for _, instruction := range instructions {
		if doRegex.MatchString(instruction) {
			isEnabled = true
			continue
		}
		if dontRegex.MatchString(instruction) {
			isEnabled = false
			continue
		}

		// Process `mul()` instructions only if enabled
		if isEnabled && mulRegex.MatchString(instruction) {
			// Extract the numbers within the `mul()`
			digitRegex := regexp.MustCompile(patternDigits)
			digitMatch := digitRegex.FindString(instruction)
			if digitMatch == "" {
				return 0, fmt.Errorf("failed to extract digits from instruction: %s", instruction)
			}

			result, err := processSingleMultiplication(digitMatch)
			if err != nil {
				return 0, err
			}

			resultOfAllMultiplications += result
		}
	}

	return resultOfAllMultiplications, nil
}

func partOne(input string) (int, error) {
	mulRegex := regexp.MustCompile(patternMul)

	// Find all matches using regex
	matches := mulRegex.FindAllString(input, -1)

	var resultOfAllMultiplications int
	for _, m := range matches {
		digitRegex := regexp.MustCompile(patternDigits)
		digitMatches := digitRegex.FindAllString(m, 1)
		for _, d := range digitMatches {
			resultOfMultiplication, err := processSingleMultiplication(d)
			if err != nil {
				return 0, err
			}

			resultOfAllMultiplications += resultOfMultiplication
		}
	}

	return resultOfAllMultiplications, nil
}

func processSingleMultiplication(digitMatch string) (int, error) {
	stringDigits := strings.Split(digitMatch, ",")
	var digits []int
	for _, stringDigit := range stringDigits {
		digit, err := strconv.Atoi(stringDigit)
		if err != nil {
			return 0, err
		}

		digits = append(digits, digit)
	}

	if len(digits) > 2 {
		return 0, fmt.Errorf("found too many digits in %s", digitMatch)
	}

	return digits[0] * digits[1], nil
}

func readFileToString(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
