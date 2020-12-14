package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveDay14Example() {
	instructions := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
	runDockingProgram(instructions)
}

func solveDay14Part1() {
	instructions := getDataFromFile("day14")
	runDockingProgram(instructions)
}

func runDockingProgram(instructions []string) {
	memory := map[string]int{}
	mask := ""
	for _, cmd := range instructions {
		runDockingCommand(cmd, &mask, memory)
	}

	res := 0
	for _, v := range memory {
		res += v
	}
	fmt.Printf("Final result : %d\n", res)
}

func runDockingCommand(cmd string, mask *string, memory map[string]int) {
	pair := strings.Split(cmd, " = ")
	if pair[0] == "mask" {
		*mask = pair[1]
	} else {
		value, _ := strconv.Atoi(pair[1])
		res := 0
		var i uint
		for i = 0; i < 36; i++ {
			if (*mask)[35-i] == '1' {
				res += 1 << i
			} else if (*mask)[35-i] == 'X' {
				res += ((value >> i) & 1) << i
			}
		}
		memory[pair[0]] = res
	}
}

func solveDay14Part2() {
	instructions := getDataFromFile("day14")

	memory := map[int]int{}
	mask := ""
	for _, cmd := range instructions {
		runDockingMemoryCommand(cmd, &mask, memory)
	}
	res := 0
	for _, v := range memory {
		res += v
	}
	fmt.Printf("Final result : %d\n", res)
}

func runDockingMemoryCommand(cmd string, mask *string, memory map[int]int) {
	pair := strings.Split(cmd, " = ")
	if pair[0] == "mask" {
		*mask = pair[1]
	} else {
		memValue, _ := strconv.Atoi(pair[1])
		memAddress := parseMemAddress(pair[0])
		memAddresses := getMemAddresses(memAddress, *mask)
		for _, addr := range memAddresses {
			memory[addr] = memValue
		}
	}
}

func getMemAddresses(memAddress int, mask string) []int {
	memAddresses := []int{0}

	var i uint
	for i = 0; i < 36; i++ {
		newAddresses := []int{}
		for _, addr := range memAddresses {
			if mask[35-i] == '1' {
				addr += 1 << i
				newAddresses = append(newAddresses, addr)
			} else if mask[35-i] == 'X' {
				newAddresses = append(newAddresses, addr)
				addr += 1 << i
				newAddresses = append(newAddresses, addr)
			} else {
				addr += ((memAddress >> i) & 1) << i
				newAddresses = append(newAddresses, addr)
			}
		}
		memAddresses = newAddresses
	}
	return memAddresses
}

func parseMemAddress(memKey string) int {
	memKey = memKey[4:]
	memKey = memKey[:len(memKey)-1]
	memValue, _ := strconv.Atoi(memKey)
	return memValue
}
