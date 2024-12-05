package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	report struct {
		levels []int
	}
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalln("a path must be provided")
	}
	path := os.Args[1]
	fmt.Printf("reading path: %s\n", path)
	reports, err := parseReportsFromPath(path)
	if err != nil {
		log.Fatal(err)
	}
	numSafe := 0
	numSafeWithOneRemoved := 0
	for _, r := range reports {
		if isReportSafe(r) {
			numSafe++
		}
		if isReportSafeWithOneRemoved(r) {
			numSafeWithOneRemoved++
		}
	}
	fmt.Printf("Part 1 - Number of safe reports: %d\n", numSafe)
	fmt.Printf("Part 2 - Number of safe reports: %d\n", numSafeWithOneRemoved)
}

func parseReportsFromPath(filePath string) ([]report, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer file.Close()
	reports := make([]report, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		levels := make([]int, 0, len(fields))
		for _, potentialLevel := range fields {
			parsedInt, err := strconv.Atoi(potentialLevel)
			if err != nil {
				return nil, fmt.Errorf("strconv.Atoi: %w", err)
			}
			levels = append(levels, parsedInt)
		}
		reports = append(reports, report{levels})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner.Err: %w", err)
	}
	return reports, nil
}

func isReportSafe(r report) bool {
	if len(r.levels) == 0 {
		return false
	}
	// state represents the unset (-1)/ decreasing (0)/ increasing(1)
	prevState := -1
	for i := 1; i < len(r.levels); i++ {
		curr, prev := r.levels[i], r.levels[i-1]
		diff := absDiff(curr, prev)
		if diff < 1 || diff > 3 {
			return false
		}
		currentState := -1
		if curr > prev {
			currentState = 1
		} else {
			currentState = 0
		}
		if prevState != -1 && currentState != prevState {
			return false
		}
		prevState = currentState
	}
	return true
}

func isReportSafeWithOneRemoved(r report) bool {
	if len(r.levels) == 0 {
		return false
	}
	if isReportSafe(r) {
		return true
	}
	for i := 0; i < len(r.levels); i++ {
		// Get the levels excluding the current index, but with copy
		copied := make([]int, len(r.levels))
		copy(copied, r.levels)
		levelsWithoutIdx := append(copied[:i], copied[i+1:]...)
		newReport := report{levels: levelsWithoutIdx}
		if isReportSafe(newReport) {
			return true
		}
	}
	return false
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	if a < b {
		return b - a
	}
	return 0
}
