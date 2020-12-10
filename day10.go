package main

import (
	"fmt"
	"sort"
)

func solveDay10Example() {
	// adapters := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	adapters := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
	computeDifferencesDistribution(adapters)

}

func solveDay10Part1() {
	adapters := getDay10Data()
	computeDifferencesDistribution(adapters)
}

func computeDifferencesDistribution(adapters []int) {
	fmt.Printf("%d adapters\n", len(adapters))
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	ones, threes := 0, 0 // final diff is 3
	for i := 1; i < len(adapters); i++ {
		if adapters[i]-adapters[i-1] == 1 {
			ones++
		} else if adapters[i]-adapters[i-1] == 3 {
			threes++
		}
	}
	fmt.Printf("There are %d differences of 1 jolt and %d differences of 3 jolt (multplied: %d)\n", ones, threes, ones*threes)
}

func solveDay10Part2() {
	// adapters := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	// adapters := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
	adapters := getDay10Data()

	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	arrangements := make([]int, len(adapters))
	count := countArrangements(adapters, 0, arrangements)
	fmt.Printf("There are a total number of %d arrangements that connect the charging outlet to the device\n", count)
}

func countArrangements(adapters []int, index int, arrangements []int) int {
	if index == len(adapters)-1 {
		return 1
	}

	if arrangements[index] > 0 {
		return arrangements[index]
	}
	count := 0
	j := index + 1
	for j < len(adapters) && adapters[j]-adapters[index] <= 3 {
		count += countArrangements(adapters, j, arrangements)
		j++
	}
	arrangements[index] = count
	return count
}

// ----------
func getDay10Data() []int {
	return convertToIntArray(getDataFromFile("day10"))
}
