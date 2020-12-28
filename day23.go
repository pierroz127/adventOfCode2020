package main

import "fmt"

type Cup struct {
	label int
	next  *Cup
	prev  *Cup
}

func solveDay23Example() {
	cups := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

	cup := playCupRounds(cups, 100)
	fmt.Println("order: ", cup.order())
	cup1 := cache[1]
	label2 := cup1.next.label
	label3 := cup1.next.next.label
	fmt.Printf("The two cups after 1 are %d and %d and they produce %d\n", label2, label3, label2*label3)
}

func (cup *Cup) print() {
	fmt.Printf("(%d)", cup.label)
	next := cup.next
	count := 0
	for next != cup && count < 15 {
		fmt.Printf(" %d ", next.label)
		next = next.next
		count++
	}
	fmt.Printf("...\n")
}

var cache []*Cup
var maxLabel int

func playCupRounds(cups []int, nbMove int) *Cup {
	maxLabel, _ = maxInt(cups)
	fmt.Printf("Max Label: %d\n", maxLabel)
	cache = make([]*Cup, len(cups)+1)
	fmt.Printf("cache of length: %d\n", len(cache))

	first := &Cup{label: cups[0]}
	cache[first.label] = first
	previous := first
	i := 1
	for i < len(cups) {
		current := &Cup{label: cups[i]}
		if current.label >= len(cache) {
			fmt.Printf("Oups, we've got a prb with label %d \n", current.label)
		}
		cache[current.label] = current
		current.prev = previous
		previous.next = current
		previous = current
		i++
	}
	previous.next = first
	first.prev = previous

	cup := first
	move := 1
	for move <= nbMove {
		cup = cup.playCupRound()
		move++
	}
	return cup
}

func (cup *Cup) playCupRound() *Cup {
	next := cup.next

	pickedCups := []int{}
	pickedUp := next
	for i := 0; i < 3; i++ {
		pickedCups = append(pickedCups, next.label)
		next = next.next
	}
	next.prev = cup
	cup.next = next

	target := decrement(cup.label)
	for contains(pickedCups, target) {
		target = decrement(target)
	}

	destination := cache[target]
	tmp := destination.next
	destination.next = pickedUp
	pickedUp.prev = destination
	for i := 1; i < 3; i++ {
		pickedUp = pickedUp.next
	}
	pickedUp.next = tmp
	tmp.prev = pickedUp
	return cup.next
}

func decrement(label int) int {
	res := label - 1
	if res == 0 {
		return maxLabel
	}
	return res
}

func (cup *Cup) order() string {
	for cup.label != 1 {
		cup = cup.next
	}
	next := cup.next
	order := ""
	for next != cup {
		order = fmt.Sprintf("%s%d", order, next.label)
		next = next.next
	}
	return order
}

func solveDay23Part1() {
	cups := []int{9, 1, 6, 4, 3, 8, 2, 7, 5}
	printCups(cups, 0)
	cup := playCupRounds(cups, 100)
	fmt.Println(cup.order())
}

func printCups(cups []int, current int) {
	fmt.Printf("Cups: ")
	for i := 0; i < 9; i++ {
		if i == current {
			fmt.Printf("(%d)", cups[i])
		} else {
			fmt.Printf(" %d ", cups[i])
		}
	}
	fmt.Printf("\n")
}

func solveDay23Part2() {
	cups := []int{9, 1, 6, 4, 3, 8, 2, 7, 5}
	for k := 10; k <= 1000000; k++ {
		cups = append(cups, k)
	}
	playCupRounds(cups, 10000000)
	cup1 := cache[1]
	label2 := cup1.next.label
	label3 := cup1.next.next.label
	fmt.Printf("The two cups after 1 are %d and %d and they produce %d\n", label2, label3, label2*label3)
}
