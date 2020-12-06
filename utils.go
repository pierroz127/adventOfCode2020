package main

import (
	"bufio"
	"fmt"
	"os"
)

func getDataFromFile(filename string) []string {
	file, err := os.Open(filename)
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
