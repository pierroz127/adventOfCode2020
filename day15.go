package main

import "fmt"

func solveDay15Example() {
	mem := map[int]int{0: 1, 3: 2, 6: 3}
	solveDay15(mem, 2020)
}

func solveDay15Part1() {
	mem := getDay15Data()
	solveDay15(mem, 2020)
}

func solveDay15Part2() {
	mem := getDay15Data()
	solveDay15(mem, 30000000)
}

func solveDay15(mem map[int]int, stop int) {
	next := 0
	for turn := len(mem) + 1; turn < stop; turn++ {
		last, ok := mem[next]
		mem[next] = turn
		if ok {
			next = turn - last
		} else {
			next = 0
		}
	}
	fmt.Printf("The %dth number spoken is %d\n", stop, next)
}

// ----------
func getDay15Data() map[int]int {
	return map[int]int{0: 1, 14: 2, 1: 3, 3: 4, 7: 5, 9: 6}

}
