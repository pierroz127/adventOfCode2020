package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Example ----------

func solveDay08Example() {
	instructions := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
	executeOnce(instructions)
}

// Part 1 ----------

func solveDay08Part1() {
	instructions := getDay08Data()
	executeOnce(instructions)
}

func executeOnce(instructions []string) {
	executed := []int{}
	step := 0
	acc := 0

	for !contains(executed, step) {
		executed = append(executed, step)
		execStep(instructions, &step, &acc)
	}
	fmt.Printf("Immediately before the program would run an instruction a second time, the value in the accumulator is %d\n", acc)
}

func execStep(instructions []string, step *int, acc *int) {
	cmd := instructions[*step]
	if strings.Contains(cmd, "nop") {
		*step = *step + 1
	} else {
		value, _ := strconv.Atoi(cmd[5:])
		sign := cmd[4:5]
		if strings.Contains(cmd, "acc") {
			if sign == "+" {
				*acc = *acc + value
			} else {
				*acc = *acc - value
			}
			*step = *step + 1
		} else if strings.Contains(cmd, "jmp") {
			if sign == "+" {
				*step = *step + value
			} else {
				*step = *step - value
			}
		}
	}
}

// Part 2 ----------

func solveDay08Part2() {
	instructions := getDay08Data()

	old := ""
	for next := 0; next < len(instructions); {
		if executeOnceOrExit(instructions) {
			break
		}
		if next > 0 {
			instructions[next-1] = old
		}
		changeNext(instructions, &next, &old)
		next++
	}
}

func executeOnceOrExit(instructions []string) bool {
	executed := []int{}
	step := 0
	acc := 0

	for !contains(executed, step) && step < len(instructions)-1 {
		executed = append(executed, step)
		execStep(instructions, &step, &acc)
	}

	if step < len(instructions)-1 {
		// fmt.Printf("Immediately before the program would run an instruction a second time, the value in the accumulator is %d\n", acc)
		return false
	}
	execStep(instructions, &step, &acc)
	fmt.Printf("The value of the accumulator after the program terminates is %d \n", acc)
	return true
}

func changeNext(instructions []string, next *int, old *string) {
	cmd := instructions[*next]

	for strings.Contains(cmd, "acc") && *next < len(instructions)-1 {
		*next++
		cmd = instructions[*next]
		// fmt.Printf("instructions[%d] = %s\n", *next, cmd)
	}
	*old = cmd
	if strings.Contains(cmd, "jmp") {
		cmd = strings.Replace(cmd, "jmp", "nop", 1)
	} else {
		cmd = strings.Replace(cmd, "nop", "jmp", 1)
	}
	instructions[*next] = cmd
}

// ----------

func getDay08Data() []string {
	return getDataFromFile("day08")
}
