package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JonathanSudibya/advent-of-code-go-2025/lib"
)

const FILENAME = "./inputs/11.txt"

func main() {
	data, err := os.ReadFile(FILENAME)
	if err != nil {
		fmt.Println("unable to read file:", err)
		os.Exit(1)
	}

	lib.CalculateAndPrintResultExecution("part_one", data, partOne)

	lib.CalculateAndPrintResultExecution("part_two", data, partTwo)
}

func buildPathsDictionary(input []byte) map[string][]string {
	pathsDictionary := make(map[string][]string)
	for line := range strings.SplitSeq(string(input), "\n") {
		if len(line) == 0 {
			continue
		}
		lineSplit := strings.Split(line, ":")
		if len(lineSplit) == 2 {
			directionSplit := strings.Split(
				strings.TrimSpace(lineSplit[1]),
				" ",
			)
			pathsDictionary[lineSplit[0]] = directionSplit
		}
	}
	return pathsDictionary
}

func partOne(input []byte) (uint64, bool) {
	pathsDictionary := buildPathsDictionary(input)

	paths := searchPaths("you", pathsDictionary)
	return paths, true
}

func searchPaths(target string, pathsDictionary map[string][]string) uint64 {
	if target == "out" {
		return 1
	}

	count := uint64(0)
	for _, destination := range pathsDictionary[target] {
		count += searchPaths(destination, pathsDictionary)
	}
	return count
}

func partTwo(input []byte) (uint64, bool) {
	pathsDictionary := buildPathsDictionary(input)

	memo := make(map[cacheKey]uint64)
	pathsCount := searchPathsServer("svr", pathsDictionary, false, false, memo)
	return pathsCount, true
}

type cacheKey struct {
	target string
	dacVisited bool
	fftVisited bool
}

func searchPathsServer(target string, pathsDictionary map[string][]string, dacVisited bool, fftVisited bool, memo map[cacheKey]uint64) uint64 {
	key := cacheKey{
		target: target,
		dacVisited: dacVisited,
		fftVisited: fftVisited,
	}
	if count, ok := memo[key]; ok {
		return count
	}

	if target == "out" {
		if dacVisited && fftVisited {
			return 1
		}
		return 0
	}

	count := uint64(0)
	
	for _, destination := range pathsDictionary[target] {
		dacV := (destination == "dac") || dacVisited
		fftV := (destination == "fft") || fftVisited
		count += searchPathsServer(destination, pathsDictionary, dacV, fftV, memo)
	}
	memo[key] = count
	return count
}
