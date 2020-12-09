package main

import "fmt"

func getPair(numbers []int) {
	for i, x := range numbers {
		for j := 0; j < i; j++ {
			if x+numbers[j] == 2020 {
				fmt.Printf("%d * %d = %d\n", x, numbers[j], x*numbers[j])
				break
			}
		}
	}
}

func solveDay01Example() {
	numbers := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	getPair(numbers)
}

func solveDay01Part1() {
	numbers := getDay01Data()
	getPair(numbers)
}

func solveDay01Part2() {
	numbers := getDay01Data()
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < i; j++ {
			for t := 0; t < j; t++ {
				x, y, z := numbers[i], numbers[j], numbers[t]
				if x+y+z == 2020 {
					fmt.Printf("%d * %d * %d = %d\n", x, y, z, x*y*z)
					break
				}
			}
		}
	}
}

// ----------
func getDay01Data() []int {
	return convertToIntArray(getDataFromFile("day01"))
}
