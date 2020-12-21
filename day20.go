package main

import (
	"fmt"
	"math"
	"strconv"
)

// Tile ----------
type Tile struct {
	ID     int
	Square [10][10]byte
}

func (tile *Tile) rotate() {
	square := [10][10]byte{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			square[j][9-i] = tile.Square[i][j]
		}
	}
	tile.Square = square
}

func (tile *Tile) flip() *Tile {
	square := [10][10]byte{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			square[i][j] = tile.Square[9-i][j]
		}
	}
	return &Tile{tile.ID, square}
}

func (tile *Tile) isUsed(rows []*Row) bool {
	for _, row := range rows {
		if contains(row.TileIDs, tile.ID) {
			return true
		}
	}
	return false
}

func (tile *Tile) match(right *Tile) []*Row {
	possibleRows := []*Row{}
	leftTiles := []*Tile{tile, tile.flip()}

	for _, left := range leftTiles {
		for leftRotation := 0; leftRotation < 4; leftRotation++ {
			for rightRotation := 0; rightRotation < 4; rightRotation++ {
				matched := true
				for i := 0; i < 10; i++ {
					if left.Square[i][9] != right.Square[i][0] {
						matched = false
						break
					}
				}
				if matched {
					rectangle := [10][]byte{}
					for i := 0; i < 10; i++ {
						line := []byte{}
						for j := 0; j < 10; j++ {
							line = append(line, left.Square[i][j])
						}
						for j := 0; j < 10; j++ {
							line = append(line, right.Square[i][j])
						}
						rectangle[i] = line
					}
					tileIDs := []int{tile.ID, right.ID}
					possibleRows = append(possibleRows, &Row{Key: "", TileIDs: tileIDs, Rectangle: rectangle})
				}
				right.rotate()
			}
			left.rotate()
		}
	}
	return possibleRows
}

// Row ----------
type Row struct {
	Key       string
	TileIDs   []int
	Rectangle [10][]byte
}

func (row *Row) buildOrGetKey() string {
	if row.Key == "" {
		row.Key = fmt.Sprintf("%d", row.TileIDs[0])
		for i := 1; i < len(row.TileIDs); i++ {
			row.Key = fmt.Sprintf("%s.%d", row.Key, row.TileIDs[i])
		}
	}
	return row.Key
}

func (row *Row) flipHorizontal() *Row {
	rectangle := [10][]byte{}
	for i := 0; i < 10; i++ {
		rectangle[i] = row.Rectangle[9-i]
	}
	return &Row{"", row.TileIDs, rectangle}
}

func (row *Row) flipVertical() *Row {
	rectangle := [10][]byte{}
	for i := 0; i < 10; i++ {
		rectangle[i] = []byte{}
		for j := 0; j < len(row.Rectangle[0]); j++ {
			rectangle[i] = append(rectangle[i], row.Rectangle[i][len(row.Rectangle[i])-1-j])
		}
	}
	tileIDs := []int{}
	for i := 0; i < len(row.TileIDs); i++ {
		tileIDs = append(tileIDs, row.TileIDs[len(row.TileIDs)-1-i])
	}
	return &Row{"", tileIDs, rectangle}
}

func (row *Row) matchLeft(leftTile *Tile) []*Row {
	possibleRows := []*Row{}
	leftTiles := []*Tile{leftTile, leftTile.flip()}
	for _, left := range leftTiles {
		for rotation := 0; rotation < 4; rotation++ {
			matched := true
			for i := 0; i < 10; i++ {
				if left.Square[i][9] != row.Rectangle[i][0] {
					matched = false
					break
				}
			}
			if matched {
				possibleRows = append(possibleRows, row.addLeft(left))
			}
			left.rotate()
		}
	}
	// rowCache[key] = possibleRows
	return possibleRows
}

func (row *Row) addLeft(tile *Tile) *Row {
	rectangle := [10][]byte{}
	for i := 0; i < 10; i++ {
		line := []byte{}
		for _, value := range tile.Square[i] {
			line = append(line, value)
		}
		for _, value := range row.Rectangle[i] {
			line = append(line, value)
		}
		rectangle[i] = line
	}
	tileIDs := []int{tile.ID}
	for _, tID := range row.TileIDs {
		tileIDs = append(tileIDs, tID)
	}
	return &Row{Key: "", TileIDs: tileIDs, Rectangle: rectangle}
}

func (row *Row) matchRight(rightTile *Tile) []*Row {
	possibleRows := []*Row{}
	rightTiles := []*Tile{rightTile, rightTile.flip()}
	for _, right := range rightTiles {
		for rotation := 0; rotation < 4; rotation++ {
			matched := true
			rowLastIdx := len(row.Rectangle[0]) - 1
			for i := 0; i < 10; i++ {
				if right.Square[i][0] != row.Rectangle[i][rowLastIdx] {
					matched = false
					break
				}
			}

			if matched {
				possibleRows = append(possibleRows, row.addRight(right))
			}
			right.rotate()
		}
	}
	return possibleRows
}

func (row *Row) addRight(tile *Tile) *Row {
	rectangle := [10][]byte{}
	for i := 0; i < 10; i++ {
		line := []byte{}
		for _, value := range row.Rectangle[i] {
			line = append(line, value)
		}
		for _, value := range tile.Square[i] {
			line = append(line, value)
		}
		rectangle[i] = line
	}
	tileIDs := []int{}
	for _, tID := range row.TileIDs {
		tileIDs = append(tileIDs, tID)
	}
	tileIDs = append(tileIDs, tile.ID)
	return &Row{Key: "", TileIDs: tileIDs, Rectangle: rectangle}
}

func (row *Row) isMatchUp(rectangle [][]byte) bool {
	height := len(rectangle)
	for i, value := range row.Rectangle[0] {
		if value != rectangle[height-1][i] {
			return false
		}
	}
	return true
}

func (row *Row) isMatchDown(rectangle [][]byte) bool {
	for i := 0; i < len(row.Rectangle[9]); i++ {
		if row.Rectangle[9][i] != rectangle[0][i] {
			return false
		}
	}
	return true
}

func (currentRow *Row) extend(dimension int, tiles []*Tile) []*Row {
	if len(currentRow.TileIDs) == dimension {
		return []*Row{currentRow}
	}

	extendedRows := []*Row{}
	for _, tile := range tiles {
		if !contains(currentRow.TileIDs, tile.ID) {
			for _, rightMatched := range currentRow.matchRight(tile) {
				for _, rightExtendedRow := range rightMatched.extend(dimension, tiles) {
					if !containsRow(extendedRows, rightExtendedRow) {
						extendedRows = append(extendedRows, rightExtendedRow)
					}
				}
			}
			for _, leftMatched := range currentRow.matchLeft(tile) {
				for _, leftExtendedRow := range leftMatched.extend(dimension, tiles) {
					if !containsRow(extendedRows, leftExtendedRow) {
						extendedRows = append(extendedRows, leftExtendedRow)
					}
				}
			}
		}
	}
	return extendedRows
}

func containsRow(rows []*Row, other *Row) bool {
	for _, row := range rows {
		if other.buildOrGetKey() == row.buildOrGetKey() {
			matched := true
			for i := 0; i < 10; i++ {
				for j := 0; j < len(row.Rectangle[i]); j++ {
					if row.Rectangle[i][j] != other.Rectangle[i][j] {
						matched = false
						break
					}
				}

				if !matched {
					break
				}
			}

			if matched {
				return true
			}
		}
	}
	return false
}

func (row *Row) toPuzzle() *Puzzle {
	rectangle := [][]byte{}
	for _, line := range row.Rectangle {
		newLine := line
		rectangle = append(rectangle, newLine)
	}
	newTileIDs := row.TileIDs
	return &Puzzle{TileIDs: [][]int{newTileIDs}, Rectangle: rectangle}
}

// Puzzle ----------
type Puzzle struct {
	TileIDs   [][]int
	Rectangle [][]byte
	ones      int
	twos      int
}

func (puzzle *Puzzle) print() {
	fmt.Println("Puzzle:")
	for _, tileIDs := range puzzle.TileIDs {
		fmt.Printf("%v\n", tileIDs)
	}
	reduced := len(puzzle.Rectangle) != 10*len(puzzle.TileIDs)
	for i := 0; i < len(puzzle.Rectangle); i++ {
		if !reduced && i > 0 && i%10 == 0 {
			fmt.Printf("[ ")
			for j := 0; j < len(puzzle.Rectangle[i]); j++ {
				if j > 0 && j%10 == 0 {
					fmt.Printf(" - ")
				}
				fmt.Printf("-")
			}
			fmt.Printf(" ]\n")
		}
		fmt.Printf("[ ")
		for j := 0; j < len(puzzle.Rectangle[i]); j++ {
			if !reduced && j > 0 && j%10 == 0 {
				fmt.Printf(" - ")
			}
			fmt.Printf("%d", puzzle.Rectangle[i][j])
		}
		fmt.Printf(" ]\n")
	}
}

func (puzzle *Puzzle) rotate() *Puzzle {
	rotatedRectangle := [][]byte{}
	for i := 0; i < len(puzzle.Rectangle[0]); i++ {
		line := []byte{}
		for j := 0; j < len(puzzle.Rectangle); j++ {
			line = append(line, puzzle.Rectangle[len(puzzle.Rectangle)-1-j][i])
		}
		rotatedRectangle = append(rotatedRectangle, line)
	}
	return &Puzzle{Rectangle: rotatedRectangle}
}

func (puzzle *Puzzle) flip() *Puzzle {
	flippedRectangle := [][]byte{}
	for i := 0; i < len(puzzle.Rectangle); i++ {
		newLine := puzzle.Rectangle[len(puzzle.Rectangle)-1-i]
		flippedRectangle = append(flippedRectangle, newLine)
	}
	return &Puzzle{Rectangle: flippedRectangle}
}

func (puzzle *Puzzle) findSeamonster(i int, j int) *Puzzle {

	if i > len(puzzle.Rectangle)-3 {
		return puzzle
	}

	var result *Puzzle
	if j < len(puzzle.Rectangle[0])-20 {
		result = puzzle.findSeamonster(i, j+1)
	} else {
		result = puzzle.findSeamonster(i+1, 0)
	}

	seaMonster := getSeaMonster()
	matched := puzzle.matchSeamonster(i, j, seaMonster)
	if matched {
		next := puzzle.addSeamonster(i, j, seaMonster)
		if j < len(puzzle.Rectangle[0])-20 {
			next = next.findSeamonster(i, j+20)
		} else {
			next = next.findSeamonster(i+1, 0)
		}
		if next.twos > result.twos {
			result = next
		}
	}
	return result
}

func (puzzle *Puzzle) count() (int, int) {
	if puzzle.ones != 0 {
		return puzzle.ones, puzzle.twos
	}
	ones, twos := 0, 0
	for i := 0; i < len(puzzle.Rectangle); i++ {
		for j := 0; j < len(puzzle.Rectangle[i]); j++ {
			if puzzle.Rectangle[i][j] == 1 {
				ones++
			} else if puzzle.Rectangle[i][j] == 2 {
				twos++
			}
		}
	}
	puzzle.ones = ones
	puzzle.twos = twos
	return ones, twos
}

func (puzzle *Puzzle) matchSeamonster(i int, j int, seaMonster [][]byte) bool {
	if i+len(seaMonster) > len(puzzle.Rectangle) {
		return false
	}
	if j+len(seaMonster[0]) > len(puzzle.Rectangle[j]) {
		return false
	}

	for ii := 0; ii < len(seaMonster); ii++ {
		for jj := 0; jj < len(seaMonster[ii]); jj++ {
			if seaMonster[ii][jj] == 1 && puzzle.Rectangle[ii+i][jj+j] != 1 {
				return false
			}
		}
	}
	return true
}

func (puzzle *Puzzle) addSeamonster(i int, j int, seaMonster [][]byte) *Puzzle {
	rectangle := [][]byte{}
	for _, line := range puzzle.Rectangle {
		newLine := line
		rectangle = append(rectangle, newLine)
	}
	for ii := 0; ii < len(seaMonster); ii++ {
		for jj := 0; jj < len(seaMonster[ii]); jj++ {
			if seaMonster[ii][jj] == 1 {
				rectangle[ii+i][jj+j] = 2
			}
		}
	}
	return &Puzzle{TileIDs: puzzle.TileIDs, Rectangle: rectangle}
}

func (puzzle *Puzzle) addDown(row *Row) *Puzzle {
	rectangle := [][]byte{}
	for _, line := range puzzle.Rectangle {
		newLine := line
		rectangle = append(rectangle, newLine)
	}
	for _, line := range row.Rectangle {
		newLine := line
		rectangle = append(rectangle, newLine)
	}
	newTileIDs := [][]int{}
	for _, tileIDsLine := range puzzle.TileIDs {
		newTileIDsLine := tileIDsLine
		newTileIDs = append(newTileIDs, newTileIDsLine)
	}
	rowTileIDs := row.TileIDs
	newTileIDs = append(newTileIDs, rowTileIDs)
	return &Puzzle{TileIDs: newTileIDs, Rectangle: rectangle}
}

func (puzzle *Puzzle) addUp(row *Row) *Puzzle {
	rectangle := [][]byte{}
	for _, line := range row.Rectangle {
		newLine := line
		rectangle = append(rectangle, newLine)
	}
	for _, line := range puzzle.Rectangle {
		newLine := line
		rectangle = append(rectangle, newLine)
	}
	rowTileIDs := row.TileIDs
	newTileIDs := [][]int{rowTileIDs}
	for _, tileIDsLine := range puzzle.TileIDs {
		newTileIDsLine := tileIDsLine
		newTileIDs = append(newTileIDs, newTileIDsLine)
	}
	return &Puzzle{TileIDs: newTileIDs, Rectangle: rectangle}
}

func (puzzle *Puzzle) reduce() {
	rectangle := [][]byte{}
	for i := 0; i < len(puzzle.TileIDs); i++ {
		for ii := 1; ii < 9; ii++ {
			row := []byte{}
			for j := 0; j < len(puzzle.TileIDs); j++ {
				for jj := 1; jj < 9; jj++ {
					row = append(row, puzzle.Rectangle[10*i+ii][10*j+jj])
				}
			}
			rectangle = append(rectangle, row)
		}
	}
	puzzle.Rectangle = rectangle
}

// Find Puzzle ----------

func findPuzzle(tiles []*Tile, rows []*Row) (*Puzzle, bool) {
	if len(rows) == dimension(tiles) {
		return tryToMatchRows(rows, []int{0}, rows[0].toPuzzle())
	}
	for _, possibleRow := range buildNewPossibleRows(tiles, rows) {
		puzzle, ok := findPuzzle(tiles, append(rows, possibleRow))
		if ok {
			return puzzle, true
		}
	}
	return nil, false
}

func dimension(tiles []*Tile) int {
	return int(math.Sqrt(float64(len(tiles))))
}

func buildNewPossibleRows(tiles []*Tile, rows []*Row) []*Row {
	possibleRows := []*Row{}
	var next *Tile
	otherTiles := []*Tile{}
	for _, tile := range tiles {
		if !tile.isUsed(rows) {
			if next == nil {
				next = tile
			} else {
				otherTiles = append(otherTiles, tile)
			}
		}
	}
	if next == nil {
		return possibleRows
	}
	dimension := dimension(tiles)
	for _, other := range otherTiles {
		for _, pair := range next.match(other) {
			for _, row := range pair.extend(dimension, otherTiles) {
				possibleRows = append(possibleRows, row)
			}
		}
	}
	return possibleRows
}

func tryToMatchRows(rows []*Row, usedindexes []int, puzzle *Puzzle) (*Puzzle, bool) {
	if len(usedindexes) == len(rows) {
		fmt.Println("\n***** Hurrah, we've got a solution *****")
		for _, tileIDs := range puzzle.TileIDs {
			fmt.Printf("%v\n", tileIDs)
		}
		fmt.Println("*****************************\n")
		return puzzle, true
	}

	for i, row := range rows {
		if contains(usedindexes, i) {
			continue
		}
		allRows := []*Row{row, row.flipHorizontal(), row.flipVertical()}
		for _, row := range allRows {
			if row.isMatchUp(puzzle.Rectangle) {
				updatedPuzzle := puzzle.addDown(row)
				finalPuzzle, ok := tryToMatchRows(rows, append(usedindexes, i), updatedPuzzle)
				if ok {
					return finalPuzzle, true
				}
			}
			if row.isMatchDown(puzzle.Rectangle) {
				updatedPuzzle := puzzle.addUp(row)
				finalPuzzle, ok := tryToMatchRows(rows, append(usedindexes, i), updatedPuzzle)
				if ok {
					return finalPuzzle, true
				}
			}
		}
	}
	return nil, false
}

func makeOneRowWithTwoTiles(tiles []*Tile, rows []*Row) []*Row {
	possibleRows := []*Row{}
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < i; j++ {
			if !tiles[i].isUsed(rows) && !tiles[j].isUsed(rows) {
				for _, possibleRow := range tiles[i].match(tiles[j]) {
					possibleRows = append(possibleRows, possibleRow)
				}
			}
		}
	}
	return possibleRows
}

func replaceLastRow(rows []*Row, last *Row) []*Row {
	return append(rows[:len(rows)-1], last)
}

// Day 20 Part 1 ----------

func solveDay20Part1() {
	tiles := parseTiles(getDataFromFile("day20"))

	dim := int(math.Sqrt(float64(len(tiles))))
	fmt.Printf("Puzzle dimension: %d\n", dim)
	findPuzzle(tiles, []*Row{})
}

// Day 20 Part 1 ----------

func solveDay20Part2() {
	tiles := parseTiles(getDataFromFile("day20"))

	dim := int(math.Sqrt(float64(len(tiles))))
	fmt.Printf("Puzzle dimension: %d\n", dim)
	puzzle, _ := findPuzzle(tiles, []*Row{})
	// puzzle, _ := findPuzzle(tiles[:9], []*Row{})
	puzzle.print()

	fmt.Println("Looking for sea monsters...")
	puzzle.reduce()

	puzzles := []*Puzzle{puzzle, puzzle.flip()}

	var best *Puzzle
	for _, pzzl := range puzzles {
		for i := 0; i < 4; i++ {
			res := pzzl.findSeamonster(0, 0)
			res.count()
			if best == nil || res.twos > best.twos {
				best = res
			}
			pzzl = pzzl.rotate()
		}
	}

	best.print()
	ones, twos := best.count()
	fmt.Printf("%d ones and %d sea monsters detected\n", ones, twos/15)
}

// ----------

func (tile *Tile) print() {
	fmt.Printf("Tile %d:\n", tile.ID)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", tile.Square[i])
	}
}

func (row *Row) print() {
	fmt.Printf("Tiles %v:\n", row.TileIDs)
	for i := 0; i < 10; i++ {
		fmt.Printf("[ ")
		for j := 0; j < len(row.Rectangle[i]); j++ {
			if j > 0 && j%10 == 0 {
				fmt.Printf(" - ")
			}
			fmt.Printf("%d", row.Rectangle[i][j])
		}
		fmt.Printf(" ]\n")
	}
}

func printTileIDs(rows []*Row) {
	for _, row := range rows {
		fmt.Printf("%v\n", row.TileIDs)
	}
}

func printSeaMonster() {
	seaMonster := getSeaMonster()
	for _, line := range seaMonster {
		fmt.Printf("%v\n", line)
	}
}

func getSeaMonster() [][]byte {
	return [][]byte{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		[]byte{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1},
		[]byte{0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0},
	}
}

// Example ----------

func solveDay20Example() {
	tiles := parseTiles(getDay20ExampleData())
	dim := int(math.Sqrt(float64(len(tiles))))
	fmt.Printf("Puzzle dimension: %d\n", dim)
	printSeaMonster()
	puzzle, _ := findPuzzle(tiles[:9], []*Row{})
	puzzle.print()
	puzzle.reduce()
	puzzles := []*Puzzle{puzzle, puzzle.flip()}
	var best *Puzzle
	for _, pzzl := range puzzles {
		for i := 0; i < 4; i++ {
			res := pzzl.findSeamonster(0, 0)
			res.count()
			if best == nil || res.twos > best.twos {
				best = res
			}
			pzzl = pzzl.rotate()
		}
	}
	best.print()
	ones, twos := best.count()
	fmt.Printf("%d ones and %d sea monsters detected\n", ones, twos/15)
}

func parseTiles(lines []string) []*Tile {
	tiles := []*Tile{}
	for idx := 0; idx < len(lines); idx++ {
		tileID := parseTileID(lines[idx])
		idx++
		tileSquare := parseTileSquare(idx, lines)
		idx += 10
		tile := &Tile{ID: tileID, Square: tileSquare}
		tiles = append(tiles, tile)
	}
	return tiles
}

func parseTileID(s string) int {
	s = s[5:]
	s = s[:len(s)-1]
	id, _ := strconv.Atoi(s)
	return id
}

func parseTileSquare(idx int, lines []string) [10][10]byte {
	square := [10][10]byte{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if lines[idx+i][j] == '.' {
				square[i][j] = 0
			} else {
				square[i][j] = 1
			}
		}
	}
	return square
}

func getDay20ExampleData() []string {
	return []string{
		"Tile 2311:",
		"..##.#..#.",
		"##..#.....",
		"#...##..#.",
		"####.#...#",
		"##.##.###.",
		"##...#.###",
		".#.#.#..##",
		"..#....#..",
		"###...#.#.",
		"..###..###",
		"",
		"Tile 1951:",
		"#.##...##.",
		"#.####...#",
		".....#..##",
		"#...######",
		".##.#....#",
		".###.#####",
		"###.##.##.",
		".###....#.",
		"..#.#..#.#",
		"#...##.#..",
		"",
		"Tile 1171:",
		"####...##.",
		"#..##.#..#",
		"##.#..#.#.",
		".###.####.",
		"..###.####",
		".##....##.",
		".#...####.",
		"#.##.####.",
		"####..#...",
		".....##...",
		"",
		"Tile 1427:",
		"###.##.#..",
		".#..#.##..",
		".#.##.#..#",
		"#.#.#.##.#",
		"....#...##",
		"...##..##.",
		"...#.#####",
		".#.####.#.",
		"..#..###.#",
		"..##.#..#.",
		"",
		"Tile 1489:",
		"##.#.#....",
		"..##...#..",
		".##..##...",
		"..#...#...",
		"#####...#.",
		"#..#.#.#.#",
		"...#.#.#..",
		"##.#...##.",
		"..##.##.##",
		"###.##.#..",
		"",
		"Tile 2473:",
		"#....####.",
		"#..#.##...",
		"#.##..#...",
		"######.#.#",
		".#...#.#.#",
		".#########",
		".###.#..#.",
		"########.#",
		"##...##.#.",
		"..###.#.#.",
		"",
		"Tile 2971:",
		"..#.#....#",
		"#...###...",
		"#.#.###...",
		"##.##..#..",
		".#####..##",
		".#..####.#",
		"#..#.#..#.",
		"..####.###",
		"..#.#.###.",
		"...#.#.#.#",
		"",
		"Tile 2729:",
		"...#.#.#.#",
		"####.#....",
		"..#.#.....",
		"....#..#.#",
		".##..##.#.",
		".#.####...",
		"####.#.#..",
		"##.####...",
		"##..#.##..",
		"#.##...##.",
		"",
		"Tile 3079:",
		"#.#.#####.",
		".#..######",
		"..#.......",
		"######....",
		"####.#..#.",
		".#...#.##.",
		"#.#####.##",
		"..#.###...",
		"..#.......",
		"..#.###...",
		"",
		"Tile 1698:",
		".#.#.#####",
		"#.#..#####",
		"#..#......",
		"#######...",
		"#####.#..#",
		".#...#.##.",
		"..#####.##",
		"..#.###...",
		"..#.......",
		"#.#.###...",
		"",
		"Tile 5022:",
		"..#.###...",
		"#.#..#####",
		"..#.......",
		".######...",
		".####.#..#",
		".#...#.##.",
		"#.#####.##",
		"..#.###...",
		"..#.......",
		".#.#.#####",
		"",
		"Tile 5023:",
		"#.#.###...",
		"#.#..#####",
		"#.........",
		".######...",
		".####.#..#",
		".#...#.##.",
		".#.#####.#",
		"#.###.....",
		"#.........",
		"..#.###...",
		"",
		"Tile 5024:",
		"..#.#####.",
		"#.#..#####",
		"#.........",
		".######...",
		".####.#..#",
		".#...#.##.",
		".#.#####.#",
		"#.###.....",
		"#.........",
		"..#.###...",
		"",
		"Tile 5025:",
		"...##.....",
		"#.#..#####",
		"#........#",
		".######...",
		".####.#...",
		".#...#.##.",
		".#.#####..",
		"#.###....#",
		"#........#",
		"..#.###...",
		"",
		"Tile 5026:",
		"##.#.#....",
		"#.#..#####",
		"#........#",
		".######...",
		".####.#...",
		".#...#.##.",
		".#.#####..",
		"#.###....#",
		"#........#",
		"..#.###...",
		"",
		"Tile 5027:",
		"..#.#....#",
		"#.#..#####",
		"#........#",
		".######...",
		".####.#...",
		".#...#.##.",
		".#.#####..",
		"#.###....#",
		"#........#",
		"..#.###...",
	}
}
