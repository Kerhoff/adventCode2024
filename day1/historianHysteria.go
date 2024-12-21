package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func calculateTotalDistance(leftList, rightList []int) int {
    totalDistance := 0

    sort.Ints(leftList)
    sort.Ints(rightList)

    for i :=0; i < len(leftList); i++ {
        totalDistance += abs(leftList[i] - rightList[i])
    }

    return totalDistance
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func readInputFile(filename string) ([]int, []int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var leftList, rightList []int

    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        if len(fields) == 2 {
            lefNum, err := strconv.Atoi(fields[0])
            if err != nil {
                return nil, nil, err
            }
            rightNum, err := strconv.Atoi(fields[1])
            if err != nil {
                return nil, nil, err
            }
            leftList = append(leftList, lefNum)
            rightList = append(rightList, rightNum)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, nil, err
    }

    return leftList, rightList, nil
}


func main() {
    inputFile := "input.txt"

    leftList, rightList, err := readInputFile(inputFile)
    if err != nil {
        fmt.Printf("Error reading input file: %v\n", err)
        return
    }

    totalDistance := calculateTotalDistance(leftList, rightList)
    fmt.Printf("Total distance: %d\n", totalDistance)
}
