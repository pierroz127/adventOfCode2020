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

type TilesMap struct {
	tiles [][][4]byte
}

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

func (tilesMap *TilesMap) moveTile(directions []int) {
	x, y := 0, 0
	for _, direction := range directions {
		x += tileDirections[direction][0]
		y += tileDirections[direction][1]
	}

	tilesMap.flipAt(x, y)
}

func (tilesMap *TilesMap) countBlack() int {
	count := 0
	for _, tilesLine := range tilesMap.tiles {
		for _, tiles := range tilesLine {
			for k := 0; k < 4; k++ {
				count += int(tiles[k])
			}
		}
	}
	return count
}

func (tilesMap *TilesMap) extend() {
	for i := 0; i < len(tilesMap.tiles); i++ {
		tilesMap.tiles[i] = append(tilesMap.tiles[i], [4]byte{0, 0, 0, 0}, [4]byte{0, 0, 0, 0})
	}
	tilesMap.tiles = append(tilesMap.tiles, [][4]byte{[4]byte{0, 0, 0, 0}, [4]byte{0, 0, 0, 0}})
	tilesMap.tiles = append(tilesMap.tiles, [][4]byte{[4]byte{0, 0, 0, 0}})
}

func executeFlipArt(directions [][]int) {
	tilesMap := &TilesMap{tiles: [][][4]byte{}}

	for _, direction := range directions {
		tilesMap.moveTile(direction)
	}

	maxIndex, maxLength := 0, len(tilesMap.tiles)
	for x, tilesLine := range tilesMap.tiles {
		fmt.Printf("%v\n", tilesLine)
		if x+len(tilesLine) >= maxIndex+maxLength {
			maxIndex = x
			maxLength = len(tilesLine)
		}
	}

	for x := 0; x <= maxIndex+maxLength; x++ {
		length := maxLength + maxIndex - x
		if x >= len(tilesMap.tiles) {
			tilesMap.tiles = append(tilesMap.tiles, [][4]byte{})
		}
		tilesLine := tilesMap.tiles[x]
		for len(tilesLine) < length {
			tilesLine = append(tilesLine, [4]byte{0, 0, 0, 0})
		}
		tilesMap.tiles[x] = tilesLine
	}

	for _, tilesLine := range tilesMap.tiles {
		fmt.Printf("%v\n", tilesLine)
	}

	day := 1
	for day <= 100 {
		tilesMap.extend()

		clone := tilesMap.clone()
		for x := 1 - len(tilesMap.tiles); x < len(tilesMap.tiles); x++ {
			y1, y2 := 0, 0
			if abs(x)%2 == 1 {
				y1, y2 = -1, 1
			}
			for y2 < len(tilesMap.tiles[abs(x)]) {
				cb1 := tilesMap.countBlackNeighbors(x, y1)
				v1 := tilesMap.getValueAt(x, y1)
				if (v1 == 1 && (cb1 == 0 || cb1 > 2)) || (v1 == 0 && cb1 == 2) {
					clone.flipAt(x, y1)
				}

				if y2 != 0 {
					cb2 := tilesMap.countBlackNeighbors(x, y2)
					v2 := tilesMap.getValueAt(x, y2)

					if (v2 == 1 && (cb2 == 0 || cb2 > 2)) || (v2 == 0 && cb2 == 2) {
						clone.flipAt(x, y2)
					}
				}
				y1 -= 2
				y2 += 2
			}
		}
		tilesMap.tiles = clone.tiles

		fmt.Printf("Day %d: %d black tiles\n", day, tilesMap.countBlack())
		day++
	}
}

func (tilesMap *TilesMap) clone() *TilesMap {
	clone := make([][][4]byte, len(tilesMap.tiles))
	for i := 0; i < len(tilesMap.tiles); i++ {
		clone[i] = make([][4]byte, len(tilesMap.tiles[i]))
		for j := 0; j < len(tilesMap.tiles[i]); j++ {
			clone[i][j] = tilesMap.tiles[i][j]
		}
	}
	return &TilesMap{tiles: clone}
}

func (tilesMap *TilesMap) getValueAt(x int, y int) byte {
	i, j, k := getCoordinates(x, y)
	if i >= len(tilesMap.tiles) || j >= len(tilesMap.tiles[i]) {
		return 0
	}
	return tilesMap.tiles[i][j][k]
}

func (tilesMap *TilesMap) flipAt(x int, y int) {
	i, j, k := getCoordinates(x, y)
	tiles := tilesMap.tiles
	for len(tiles) <= i {
		tiles = append(tiles, [][4]byte{})
	}
	for len(tiles[i]) <= j {
		tiles[i] = append(tiles[i], [4]byte{0, 0, 0, 0})
	}

	tile := tiles[i][j]
	if tile[k] == 0 {
		tile[k] = 1
	} else {
		tile[k] = 0
	}
	tiles[i][j] = tile
	tilesMap.tiles = tiles
}

func (tilesMap *TilesMap) countBlackNeighbors(x int, y int) int {
	surroundings := [][]int{
		[]int{x - 2, y},
		[]int{x - 1, y + 1},
		[]int{x + 1, y + 1},
		[]int{x + 2, y},
		[]int{x + 1, y - 1},
		[]int{x - 1, y - 1},
	}

	count := 0
	for _, surrounding := range surroundings {
		count += int(tilesMap.getValueAt(surrounding[0], surrounding[1]))
	}
	return count
}

// Day 24 Example ----------

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

	executeFlipArt(directions)
}

// Day 24 Part 1 ----------

func solveDay24Part1() {
	directions := getDay24Data()
	tilesMap := &TilesMap{tiles: [][][4]byte{}}

	for _, direction := range directions {
		tilesMap.moveTile(direction)
	}

	for _, tilesLine := range tilesMap.tiles {
		fmt.Printf("%v\n", tilesLine)
	}

	fmt.Printf("There are %d black tiles\n", tilesMap.countBlack())
}

// Day 24 Part 2 ----------

func solveDay24Part2() {
	directions := getDay24Data()
	executeFlipArt(directions)
}

// Day 24 Data ----------

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
