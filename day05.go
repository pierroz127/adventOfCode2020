package main

import (
	"fmt"
	"sort"
)

// "regexp"
// "strconv"

func search(seats string, left int, right int) int {
	//fmt.Printf("seats: %s left: %d right: %d\n", seats, left, right)
	if len(seats) == 1 {
		if seats[0] == 'F' || seats[0] == 'L' {
			return left
		}
		return right
	}
	if seats[0] == 'F' || seats[0] == 'L' {
		return search(seats[1:], left, (left+right)/2)
	}
	return search(seats[1:], (left+right+1)/2, right)
}

func getSeatID(seat string) int {
	row := search(seat[:7], 0, 127)
	col := search(seat[7:], 0, 7)
	return row*8 + col
}

func processBoardingPass(boardingPass []string) {
	maxSeatID := 0
	for _, seat := range boardingPass {
		seatID := getSeatID(seat)
		//fmt.Printf("seat ID for %s is %d\n", seat, seatID)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Printf("Max seat id is %d\n", maxSeatID)
}

func solveDay05Example() {
	boardingPass := []string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
	processBoardingPass(boardingPass)
}

func solveDay05Part1() {
	boardingPass := getDay05Data()
	processBoardingPass(boardingPass)
}

func solveDay05Part2() {
	occupiedSeats := []int{}
	boardingPass := getDay05Data()
	for _, seat := range boardingPass {
		seatID := getSeatID(seat)
		occupiedSeats = append(occupiedSeats, seatID)
	}

	sort.Ints(occupiedSeats[:])

	for idx := 0; idx < len(occupiedSeats)-1; idx++ {
		prev := occupiedSeats[idx]
		next := occupiedSeats[idx+1]
		if next == prev+2 {
			fmt.Printf("Your seat is probably %d\n", prev+1)
		}
	}

}

// ----------
func getDay05Data() []string {
	return getDataFromFile("day05")
}
