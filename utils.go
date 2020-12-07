package main

import (
	"bufio"
	"fmt"
	"os"
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

func contains(arr []int, value int) bool {
	for _, el := range arr {
		if el == value {
			return true
		}
	}
	return false
}

func stringContains(arr []string, value string) bool {
	for _, el := range arr {
		if el == value {
			return true
		}
	}
	return false
}
