package main

import "fmt"

func solveDay25Example() {
	loopSize := findLoopSize(5764801, 7)
	fmt.Printf("Card's loop size is %d\n", loopSize)
	doorLoop := findLoopSize(17807724, 7)
	fmt.Printf("Door's loop size is %d\n", doorLoop)
	encryptionKey := encrypt(5764801, doorLoop)
	fmt.Printf("Encryption's key: %d\n", encryptionKey)
}

func solveDay25Part1() {
	cardLoop := findLoopSize(2959251, 7)
	fmt.Printf("Card's loop size is %d\n", cardLoop)
	doorLoop := findLoopSize(4542595, 7)
	fmt.Printf("Door's loop size is %d\n", doorLoop)
	encryptionKey := encrypt(2959251, doorLoop)
	fmt.Printf("Encryption's key: %d\n", encryptionKey)
}

func findLoopSize(publicKey int, subject int) int {
	value := 1
	loop := 0
	for value != publicKey {
		value *= subject
		value = value % 20201227
		loop++
	}
	return loop
}

func encrypt(subject int, loop int) int {
	value := 1
	i := 0
	for i < loop {
		value *= subject
		value = value % 20201227
		i++
	}
	return value
}

func solveDay25Part2() {
	fmt.Println("The light turns green and the door unlocks.")
}
