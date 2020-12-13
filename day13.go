package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func solveDay13Example() {
	g, x, y := extendedGcd(3, 35)
	fmt.Printf("egcd(3, 35) = (%d, %d, %d)\n", g, x, y)
	g, x, y = extendedGcd(5, 21)
	fmt.Printf("egcd(5, 21) = (%d, %d, %d)\n", g, x, y)
	g, x, y = extendedGcd(7, 15)
	fmt.Printf("egcd(7, 15) = (%d, %d, %d)\n", g, x, y)

	timeTable := "7,13,x,x,59,x,31,19"

	busIDs, offsets := parseBusIDsWithOffset(timeTable)
	newOffsets := getNewOffsets(busIDs, offsets)
	x = chineseRemainders(busIDs, newOffsets)
	fmt.Printf("The earliest time for all the bus depart at offset matching their position in the list is %d\n", x)

}

func solveDay13Part1() {
	time, busIDs := getDay13DataPart1()
	findNextDeparture(time, busIDs)
}

func findNextDeparture(time int, busIDs []int) {
	minBusID, minIndex, waiting := 0, -1, 0

	for i, b := range busIDs {
		x := time / b
		next := (x + 1) * b
		if next < minBusID || minBusID == 0 {
			minBusID = next
			minIndex = i
			waiting = next - time
		}
	}
	fmt.Printf("Next departure for bus %d at %d and you'll have to wait %d minutes (%d*%d = %d)\n",
		busIDs[minIndex],
		minBusID,
		waiting,
		busIDs[minIndex],
		waiting,
		busIDs[minIndex]*waiting,
	)
}

// Part 2 ----------

func solveDay13Part2() {
	busIDs, offsets := getDay13DataPart2()
	newOffsets := getNewOffsets(busIDs, offsets)
	x := chineseRemainders(busIDs, newOffsets)
	fmt.Printf("The earliest time for all the bus depart at offset matching their position in the list is %d\n", x)
}

func getNewOffsets(busIDs []int, offsets []int) []int {
	newOffsets := []int{0}
	for i := 1; i < len(offsets); i++ {
		if offsets[i] > busIDs[i] {
			_, r := divide(offsets[i], busIDs[i])
			newOffsets = append(newOffsets, busIDs[i]-r)
		} else {
			newOffsets = append(newOffsets, busIDs[i]-offsets[i])
		}
	}
	return newOffsets
}

func extendedGcd(a int, b int) (int, int, int) {
	oldR, r := a, b
	oldS, s := 1, 0
	oldT, t := 0, 1

	for r != 0 {
		q := oldR / r
		r0 := r
		r = oldR - q*r0
		oldR = r0
		s0 := s
		s = oldS - q*s0
		oldS = s0
		t0 := t
		t = oldT - q*t0
		oldT = t0
	}
	return oldR, oldS, oldT
}

func divide(x int, y int) (int, int) {
	q := x / y
	r := x - q*y
	return q, r
}

func product(n []int) int {
	N := 1
	for _, ni := range n {
		N *= ni
	}
	return N
}

func chineseRemainders(n []int, a []int) int {
	ntilde := []int{}

	for i := range n {
		nt := 1
		for j := range n {
			if i != j {
				nt *= n[j]
			}
		}
		ntilde = append(ntilde, nt)
	}

	e := []int{}

	for i := range n {
		g, _, vi := extendedGcd(n[i], ntilde[i])
		if g != 1 {
			log.Fatalf("gcd(%d, %d) != 1 !\n", n[i], ntilde[i])
		}
		e = append(e, vi*ntilde[i])
	}

	x := 0
	for i := range a {
		x += a[i] * e[i]
	}

	N := product(n)
	_, r := divide(x, N)

	if r < 0 {
		r = N - r
	}
	return r
}

// ----------

func getDay13DataPart1() (int, []int) {
	return 1000390, parseBusIDs("13,x,x,41,x,x,x,x,x,x,x,x,x,997,x,x,x,x,x,x,x,23,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,x,x,x,x,x,29,x,619,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,17")
}

func getDay13DataPart2() ([]int, []int) {
	// BusIDs and Offset
	return parseBusIDsWithOffset("13,x,x,41,x,x,x,x,x,x,x,x,x,997,x,x,x,x,x,x,x,23,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,x,x,x,x,x,29,x,619,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,17")
}

func getDay13ExampleData() (int, []int) {
	return 939, parseBusIDs("7,13,x,x,59,x,31,19")
}

func parseBusIDs(busTimeTable string) []int {
	busIDs := []int{}
	for _, s := range strings.Split(busTimeTable, ",") {
		busID, err := strconv.Atoi(s)
		if err == nil {
			busIDs = append(busIDs, busID)
		}
	}
	return busIDs
}

func parseBusIDsWithOffset(busTimeTable string) ([]int, []int) {
	busIDs, offsets := []int{}, []int{}
	for i, s := range strings.Split(busTimeTable, ",") {
		busID, err := strconv.Atoi(s)
		if err == nil {
			busIDs = append(busIDs, busID)
			offsets = append(offsets, i)
		}
	}
	return busIDs, offsets
}
