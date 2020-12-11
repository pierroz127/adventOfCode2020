package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getDataFromFile(filename string) []string {
	file, err := os.Open(fmt.Sprintf("inputs/%s.txt", filename))
	if err != nil {
		fmt.Printf("can't open %s\n", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func convertToIntArray(lines []string) []int {
	integers := []int{}
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		integers = append(integers, i)
	}
	return integers
}

func contains(arr []int, value int) bool {
	for _, el := range arr {
		if el == value {
			return true
		}
	}
	return false
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i int, j int) int {
	if i < j {
		return j
	}
	return i
}

func minInt(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("Input array is empty")
	}
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min, nil
}

type error interface {
	Error() string
}

func minIntIndex(arr []int) (int, error) {
	if len(arr) == 0 {
		return -1, fmt.Errorf("Input array is empty")
	}
	minIdx := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[minIdx] {
			minIdx = i
		}
	}
	return minIdx, nil
}

func maxInt(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("Input array is empty") // By convention... we should
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max, nil
}

func maxIntIndex(arr []int) (int, error) {
	if len(arr) == 0 {
		return -1, fmt.Errorf("Input array is empty") // By convention... we should
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max, nil
}

func stringContains(arr []string, value string) bool {
	for _, el := range arr {
		if el == value {
			return true
		}
	}
	return false
}

func areEqual(expected []string, actual []string) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i, s := range expected {
		if s != actual[i] {
			return false
		}
	}

	return true
}
