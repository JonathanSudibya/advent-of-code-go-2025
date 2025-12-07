package main

import (
	"fmt"
	"os"

	"github.com/JonathanSudibya/advent-of-code-go-2025/lib"
)

const FILENAME = "./inputs/1.txt"

func main() {
	data, err := os.ReadFile(FILENAME)
	if err != nil {
		fmt.Println("unable to read file:", err)
		os.Exit(1)
	}

	lib.CalculateAndPrintResultExecution("part_one", data, partOne)

	lib.CalculateAndPrintResultExecution("part_two", data, partTwo)
}

func partOne(input []byte) (uint64, bool) {
	lockPosition := 50
	result := uint64(0)

	start := 0

	for i := range len(input) {
		if input[i] != '\n' {
			continue
		}

		line := input[start:i]
		start = i + 1

		if len(line) == 0 {
			continue
		}

		moveSize, err := lib.BytesToInt(line[1:])
		if err != nil {
			fmt.Println("unable to parse int", err)
			return 0, false
		}

		switch line[0] {
		case 'L':
			lockPosition = (lockPosition - moveSize) % 100
			if lockPosition < 0 {
				lockPosition += 100
			}
		case 'R':
			lockPosition = (lockPosition + moveSize) % 100
		}

		if lockPosition == 0 {
			result++
		}
	}

	return result, true
}

func partTwo(input []byte) (uint64, bool) {
	lockPosition := 50
	result := uint64(0)

	start := 0

	for i := range len(input) {
		if input[i] != '\n' {
			continue
		}

		line := input[start:i]
		start = i + 1

		if len(line) == 0 {
			continue
		}

		moveSize, err := lib.BytesToInt(line[1:])
		if err != nil {
			fmt.Println("unable to parse int", err)
			return 0, false
		}

		previousPosition := lockPosition
		rotationTimes := moveSize / 100

		switch line[0] {
		case 'L':
			lockPosition -= (moveSize % 100)
		case 'R':
			lockPosition += (moveSize % 100)
		}

		if lockPosition < 0 {
			lockPosition += 100
			if previousPosition != 0 {
				rotationTimes += 1
			}
		} else if lockPosition > 99 {
			lockPosition -= 100
			if previousPosition != 0 {
				rotationTimes += 1
			}
		} else if lockPosition == 0 {
			if previousPosition != 0 {
				rotationTimes += 1
			}
		}

		result += uint64(rotationTimes)
	}

	return result, true
}
