package main

import "fmt"

func reachBottom(grid []string) {
	width := len(grid[0])
	count := 0
	j := 0
	for i := 0; i < len(grid); i++ {
		if grid[i][j] == '#' {
			count++
		}
		j += 3
		j = j % width
	}
	fmt.Printf("You hit %d trees\n", count)
}

func solveDay03Example() {
	grid := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	reachBottom(grid)
}

func solveDay03Part1() {
	grid := getDay03Data()
	reachBottom(grid)
}

func solveDay03Part2() {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	grid := getDay03Data()
	width := len(grid[0])

	results := []int{}
	for _, slope := range slopes {
		count := 0
		j := 0
		for i := 0; i < len(grid); i += slope[1] {
			if grid[i][j] == '#' {
				count++
			}
			j += slope[0]
			j = j % width
		}
		results = append(results, count)
	}

	res := 1
	for _, r := range results {
		res = res * r
	}
	fmt.Printf("You hit %d trees\n", res)
}

// ----------
func getDay03Data() []string {
	return getDataFromFile("day03")
}
