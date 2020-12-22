package main

import "strings"

// "regexp"
// "strconv"

func solveDay21Example() {

}

func solveDay21Part1() {
}

func solveDay21Part2() {

}

// ----------
func getDay21Data() []string {
	return []string{}
}

func parseFoodLines(lines []string) ([][]string, [][]string) {
	ingredients := make([][]string, len(lines))
	allergens := make([][]string, len(lines))

	for i, line := range lines {
		idx := strings.Index(line, " (contains")
		ingredients[i] = strings.Split(line[:idx], " ")
		allergens[i] = strings.Split(line[idx+11:len(line)-1], ", ")
	}
	return ingredients, allergens
}
