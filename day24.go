package main

import "fmt"

var EAST = 0
var NORTHEAST = 1
var NORTHWEST = 2
var WEST = 3
var SOUTHWEST = 4
var SOUTHEAST = 5

var tileDirections = [][]int{
	[]int{-2, 0},
	[]int{-1, 1},
	[]int{1, 1},
	[]int{2, 0},
	[]int{1, -1},
	[]int{-1, -1},
}

var tilesMap = [][]string{}

func getCoordinates(x int, y int) (int, int, int) {
	i := 0
	if x != 0 {
		i = x / abs(x)
	}
	j := 0
	if y != 0 {
		j = y / abs(y)
	}

	return abs(x), abs(y), i + abs(i) + (j+abs(j))/2
}

func solveDay24Example() {
	directions := parseTileLines([]string{
		"sesenwnenenewseeswwswswwnenewsewsw",
		"neeenesenwnwwswnenewnwwsewnenwseswesw",
		"seswneswswsenwwnwse",
		"nwnwneseeswswnenewneswwnewseswneseene",
		"swweswneswnenwsewnwneneseenw",
		"eesenwseswswnenwswnwnwsewwnwsene",
		"sewnenenenesenwsewnenwwwse",
		"wenwwweseeeweswwwnwwe",
		"wsweesenenewnwwnwsenewsenwwsesesenwne",
		"neeswseenwwswnwswswnw",
		"nenwswwsewswnenenewsenwsenwnesesenew",
		"enewnwewneswsewnwswenweswnenwsenwsw",
		"sweneswneswneneenwnewenewwneswswnese",
		"swwesenesewenwneswnwwneseswwne",
		"enesenwswwswneneswsenwnewswseenwsese",
		"wnwnesenesenenwwnenwsewesewsesesew",
		"nenewswnwewswnenesenwnesewesw",
		"eneswnwswnwsenenwnwnwwseeswneewsenese",
		"neswnwewnwnwseenwseesewsenwsweewe",
		"wseweeenwnesenwwwswnew"})
	fmt.Printf("Directions: %v\n", directions)

	for _, direction := range directions {
		moveTile(direction)
	}

	for _, tilesLine := range tilesMap {
		fmt.Printf("%v\n", tilesLine)
	}

	count := 0
	for _, tilesLine := range tilesMap {
		for _, tiles := range tilesLine {
			for k := 0; k < 4; k++ {
				if tiles[k] == 'B' {
					count++
				}
			}
		}
	}
	fmt.Printf("There are %d black tiles\n", count)
}

func moveTile(directions []int) {
	x, y := 0, 0
	for _, direction := range directions {
		x += tileDirections[direction][0]
		y += tileDirections[direction][1]
	}
	i, j, k := getCoordinates(x, y)

	// fmt.Printf("Flip tile at (%d, %d) ==> [%d, %d, %d]\n", x, y, i, j, k)

	for len(tilesMap) <= i {
		tilesMap = append(tilesMap, []string{})
	}
	for len(tilesMap[i]) <= j {
		tilesMap[i] = append(tilesMap[i], "WWWW")
	}

	tile := tilesMap[i][j]
	// fmt.Printf("current tile: %s\n", tile)
	var newTile string
	if tile[k] == 'W' {
		newTile = tile[:k] + "B"
	} else {
		newTile = tile[:k] + "W"
	}
	if k < 3 {
		newTile += tile[k+1:]
	}

	// fmt.Printf("new tile: %s\n", tile)

	tilesMap[i][j] = newTile
}

func solveDay24Part2() {
}

func solveDay24Part1() {
	directions := getDay24Data()
	for _, direction := range directions {
		moveTile(direction)
	}

	for _, tilesLine := range tilesMap {
		fmt.Printf("%v\n", tilesLine)
	}

	count := 0
	for _, tilesLine := range tilesMap {
		for _, tiles := range tilesLine {
			for k := 0; k < 4; k++ {
				if tiles[k] == 'B' {
					count++
				}
			}
		}
	}
	fmt.Printf("There are %d black tiles\n", count)
}

// ----------
func getDay24Data() [][]int {
	return parseTileLines(getDataFromFile("day24"))
}

func parseTileLines(lines []string) [][]int {
	directions := [][]int{}
	for _, line := range lines {
		directions = append(directions, parseTileLine(line))
	}
	return directions
}

func parseTileLine(line string) []int {
	directions := []int{}

	idx := 0
	for idx < len(line) {
		if line[idx] == 'e' {
			directions = append(directions, EAST)
		} else if line[idx] == 'w' {
			directions = append(directions, WEST)
		} else if line[idx] == 's' {
			idx++
			if line[idx] == 'e' {
				directions = append(directions, SOUTHEAST)
			} else if line[idx] == 'w' {
				directions = append(directions, SOUTHWEST)
			} else {
				fmt.Printf("oups! we've got a problem! (at idx %d)\n", idx)
			}
		} else if line[idx] == 'n' {
			idx++
			if line[idx] == 'e' {
				directions = append(directions, NORTHEAST)
			} else if line[idx] == 'w' {
				directions = append(directions, NORTHWEST)
			} else {
				fmt.Printf("oups! we've got a problem! (at idx %d)\n", idx)
			}
		} else {
			fmt.Printf("oups! we've got a problem! (at idx %d)\n", idx)
		}
		idx++
	}

	return directions
}
