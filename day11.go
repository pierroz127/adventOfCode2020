package main

import "fmt"

// Example ----------

func solveDay11Example() {
	seats := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	occupiedSeats := getFinalOccupiedSeatsCountPart2(seats)
	fmt.Printf("The final number of occupied seats is %d\n", occupiedSeats)
}

// Part 1 ----------

func solveDay11Part1() {
	seats := getDay11Data()
	occupiedSeats := getFinalOccupiedSeatsCount(seats)
	fmt.Printf("The final number of occupied seats is %d\n", occupiedSeats)
}

func getFinalOccupiedSeatsCount(seats []string) int {
	for true {
		newSeats := processSeatOccupation(seats)
		if areEqual(newSeats, seats) {
			occupiedSeats := countOccupiedSeats(newSeats)
			return occupiedSeats
		}
		seats = newSeats
	}
	return 0
}

func processSeatOccupation(seats []string) []string {
	newSeats := []string{}
	for row := 0; row < len(seats); row++ {
		newRow := ""
		for col := 0; col < len(seats[row]); col++ {
			count := getOccupiedAdjacentSeats(row, col, seats)
			if count == 0 && seats[row][col] == 'L' {
				newRow += "#"
			} else if count >= 4 && seats[row][col] == '#' {
				newRow += "L"
			} else {
				newRow += seats[row][col : col+1]
			}
		}
		newSeats = append(newSeats, newRow)
	}
	return newSeats
}

func getOccupiedAdjacentSeats(row int, col int, seats []string) int {
	count := 0
	for i := max(row-1, 0); i <= min(row+1, len(seats)-1); i++ {
		for j := max(col-1, 0); j <= min(col+1, len(seats[i])-1); j++ {
			if (i != row || j != col) && seats[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func countOccupiedSeats(seats []string) int {
	count := 0
	for _, row := range seats {
		for col := 0; col < len(row); col++ {
			if row[col] == '#' {
				count++
			}
		}
	}
	return count
}

// Part 2 ----------

func solveDay11Part2() {
	seats := getDay11Data()
	occupiedSeats := getFinalOccupiedSeatsCountPart2(seats)
	fmt.Printf("The final number of occupied seats is %d\n", occupiedSeats)
}

func getFinalOccupiedSeatsCountPart2(seats []string) int {
	step := 0
	for true {
		newSeats := processSeatOccupationPart2(seats)
		if areEqual(newSeats, seats) {
			occupiedSeats := countOccupiedSeats(newSeats)
			return occupiedSeats
		}
		for _, row := range newSeats {
			fmt.Println(row)
		}
		seats = newSeats
		step++
		fmt.Printf(" Step %d ----------\n", step)
	}
	return 0
}

func processSeatOccupationPart2(seats []string) []string {
	newSeats := []string{}
	for row := 0; row < len(seats); row++ {
		newRow := ""
		for col := 0; col < len(seats[row]); col++ {
			count := getOccupiedSeatsAround(row, col, seats)
			if count == 0 && seats[row][col] == 'L' {
				newRow += "#"
			} else if count >= 5 && seats[row][col] == '#' {
				newRow += "L"
			} else {
				newRow += seats[row][col : col+1]
			}
		}
		newSeats = append(newSeats, newRow)
	}
	return newSeats
}

func getOccupiedSeatsAround(row int, col int, seats []string) int {
	count := 0
	if isLeftOccupied(row, col, seats) {
		count++
	}
	if isRightOccupied(row, col, seats) {
		count++
	}
	if isUpOccupied(row, col, seats) {
		count++
	}
	if isDownOccupied(row, col, seats) {
		count++
	}
	if isTopLeftOccupied(row, col, seats) {
		count++
	}
	if isTopRightOccupied(row, col, seats) {
		count++
	}
	if isDownLeftOccupied(row, col, seats) {
		count++
	}
	if isDownRightOccupied(row, col, seats) {
		count++
	}
	return count
}

func isLeftOccupied(row int, col int, seats []string) bool {
	for j := col - 1; j >= 0; j-- {
		if seats[row][j] == '#' {
			return true
		} else if seats[row][j] == 'L' {
			return false
		}
	}
	return false
}

func isRightOccupied(row int, col int, seats []string) bool {
	for j := col + 1; j < len(seats[row]); j++ {
		if seats[row][j] == '#' {
			return true
		} else if seats[row][j] == 'L' {
			return false
		}
	}
	return false
}

func isUpOccupied(row int, col int, seats []string) bool {
	for i := row - 1; i >= 0; i-- {
		if seats[i][col] == '#' {
			return true
		} else if seats[i][col] == 'L' {
			return false
		}
	}
	return false
}

func isDownOccupied(row int, col int, seats []string) bool {
	for i := row + 1; i < len(seats); i++ {
		if seats[i][col] == '#' {
			return true
		} else if seats[i][col] == 'L' {
			return false
		}
	}
	return false
}

func isTopLeftOccupied(row int, col int, seats []string) bool {
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if seats[i][j] == '#' {
			return true
		} else if seats[i][j] == 'L' {
			return false
		}
	}
	return false
}

func isTopRightOccupied(row int, col int, seats []string) bool {
	for i, j := row-1, col+1; i >= 0 && j < len(seats[row]); i, j = i-1, j+1 {
		if seats[i][j] == '#' {
			return true
		} else if seats[i][j] == 'L' {
			return false
		}
	}
	return false
}

func isDownLeftOccupied(row int, col int, seats []string) bool {
	for i, j := row+1, col-1; i < len(seats) && j >= 0; i, j = i+1, j-1 {
		if seats[i][j] == '#' {
			return true
		} else if seats[i][j] == 'L' {
			return false
		}
	}
	return false
}

func isDownRightOccupied(row int, col int, seats []string) bool {
	for i, j := row+1, col+1; i < len(seats) && j < len(seats[row]); i, j = i+1, j+1 {
		if seats[i][j] == '#' {
			return true
		} else if seats[i][j] == 'L' {
			return false
		}
	}
	return false
}

// Data ----------

func getDay11Data() []string {
	return getDataFromFile("day11")
}
