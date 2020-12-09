package main

import (
	"fmt"
	"strconv"
)

// Example ----------

func solveDay09Example() {
	preambleLength := 5
	codes := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	findFirstCodeNotSumOfTwo(codes, preambleLength)
}

// Part 1 ----------

func solveDay09Part1() {
	codes := getDay09Data()
	findFirstCodeNotSumOfTwo(codes, 25)
}

func findFirstCodeNotSumOfTwo(codes []int, preambleLength int) int {
	for i := preambleLength; i < len(codes); i++ {
		if !isSumOfTwo(codes[i-preambleLength:i], codes[i]) {
			fmt.Printf("the first number which is not the sum of two of the %d numbers before it is %d\n", preambleLength, codes[i])
			return codes[i]
		}
	}
	return 0
}

func isSumOfTwo(arr []int, value int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == value {
				return true
			}
		}
	}
	return false
}

// Part 2 ----------

func solveDay09Part2() {
	codes := getDay09Data()
	// codes := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	target := findFirstCodeNotSumOfTwo(codes, 25)
	subset := findContiguousCodes(codes, target)

	minCode, _ := minInt(subset)
	maxCode, _ := maxInt(subset)
	fmt.Printf("%d+%d = %d\n", minCode, maxCode, minCode+maxCode)
}

func findContiguousCodes(codes []int, target int) []int {
	for i := 0; i < len(codes)-1; i++ {
		sum := codes[i]

		for j := i + 1; j < len(codes); j++ {
			sum += codes[j]
			if sum == target {
				return codes[i : j+1]
			}
			if sum > target {
				break
			}
		}
	}
	return []int{}
}

// Data ----------

func getDay09Data() []int {
	scodes := getDataFromFile("day09")
	codes := []int{}
	for _, s := range scodes {
		code, _ := strconv.Atoi(s)
		codes = append(codes, code)
	}
	return codes
}
