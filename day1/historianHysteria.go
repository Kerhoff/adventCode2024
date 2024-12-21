package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// calculateTotalDistance computes the total distance between the paired elements of two sorted lists.
func calculateTotalDistance(leftList, rightList []int) int {
	totalDistance := 0

	// Ensure both lists are sorted if not already
	if !isSorted(leftList) {
		sort.Ints(leftList)
	}
	if !isSorted(rightList) {
		sort.Ints(rightList)
	}

	// Use a wait group for parallel processing
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < len(leftList); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			distance := abs(leftList[i] - rightList[i])
			mu.Lock()
			totalDistance += distance
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	return totalDistance
}

// abs returns the absolute value of an integer.
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// isSorted checks if a list is already sorted.
func isSorted(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			return false
		}
	}
	return true
}

// readInputFile reads the input file and returns two slices of integers representing the two lists.
func readInputFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file '%s': %v", filename, err)
	}
	defer file.Close() // Use defer immediately after opening the file

	scanner := bufio.NewScanner(file)
	var leftList, rightList []int
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 2 {
			leftNum, err := strconv.Atoi(fields[0])
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing left number on line %d: %v", lineNumber, err)
			}
			rightNum, err := strconv.Atoi(fields[1])
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing right number on line %d: %v", lineNumber, err)
			}
			leftList = append(leftList, leftNum)
			rightList = append(rightList, rightNum)
		} else {
			return nil, nil, fmt.Errorf("invalid format on line %d", lineNumber)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file '%s': %v", filename, err)
	}

	return leftList, rightList, nil
}

func main() {
	// File containing input data
	inputFile := "input.txt"

	// Read input data from file
	leftList, rightList, err := readInputFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Calculate the total distance between the two lists
	totalDistance := calculateTotalDistance(leftList, rightList)

	// Print the result
	fmt.Printf("Total Distance: %d\n", totalDistance)
}
