package main

import (
	"bufio"
	"cmp"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalln("a path must be provided")
	}
	path := os.Args[1]
	list1, list2, err := parseLists(path)
	if err != nil {
		log.Fatal(err)
	}
	result, err := calculateTotalDistance(list1, list2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total distance result: %d\n", result)
	similarityScore, err := calculateSimilarityScore(list1, list2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Similarity score: %d\n", similarityScore)
}

func parseLists(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("os.Open: %w", err)
	}
	defer file.Close()
	var (
		list1, list2 []int
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, nil, errors.New("expected only two fields")
		}
		list1Val, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("strconv.Itoa: %w", err)
		}
		list2Val, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("strconv.Itoa: %w", err)
		}
		list1 = append(list1, list1Val)
		list2 = append(list2, list2Val)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("scanner.Err: %w", err)
	}
	return list1, list2, nil
}

func calculateTotalDistance(list1, list2 []int) (int, error) {
	if len(list1) != len(list2) {
		return -1, errors.New("lists must be of equal length")
	}
	// Sort both lists in ascending order.
	slices.SortFunc(list1, cmp.Compare[int])
	slices.SortFunc(list2, cmp.Compare[int])
	// Loop over both and calculate distance, and sum them.
	length := len(list1)
	sum := 0
	for i := 0; i < length; i++ {
		sum += absDiff(list1[i], list2[i])
	}
	return sum, nil
}

func calculateSimilarityScore(list1, list2 []int) (int, error) {
	if len(list1) != len(list2) {
		return -1, errors.New("lists must be of equal length")
	}
	// For list2, create a map with the number of appearances.
	numAppsMap := make(map[int]int, 0)
	for _, el := range list2 {
		currVal, ok := numAppsMap[el]
		if !ok {
			numAppsMap[el] = 1
		}
		numAppsMap[el] = currVal + 1
	}
	// Loop over first list and calculate similarity score.
	similarityScore := 0
	for _, el := range list1 {
		numApps, ok := numAppsMap[el]
		if !ok {
			continue
		}
		similarityScore = similarityScore + (el * numApps)
	}
	return similarityScore, nil
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
