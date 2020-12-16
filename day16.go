package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveDay16Example() {
	lines := []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
		"",
		"your ticket:",
		"7,1,14",
		"",
		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}
	fields, _, nearbyTickets := (parseFullTicket(lines))
	checkNearbyTicketsValidity(fields, nearbyTickets)
}

func solveDay16Part1() {
	lines := getDataFromFile("day16")
	fields, _, nearbyTickets := (parseFullTicket(lines))
	checkNearbyTicketsValidity(fields, nearbyTickets)
}

func checkNearbyTicketsValidity(rules [][]int, nearbyTickets [][]int) {
	errorRate := 0
	for _, nearbyTicket := range nearbyTickets {
		for _, fieldValue := range nearbyTicket {
			if !isTicketFieldValid(fieldValue, rules) {
				errorRate += fieldValue
			}
		}
	}

	fmt.Printf("error rate: %d\n", errorRate)
}

func isTicketFieldValid(field int, rules [][]int) bool {
	for _, rule := range rules {
		if field >= rule[0] && field <= rule[1] {
			return true
		}
		if field >= rule[2] && field <= rule[3] {
			return true
		}
	}
	return false
}

func solveDay16Part2() {
	lines := getDataFromFile("day16")
	fields, ticket, nearbyTickets := (parseFullTicket(lines))
	checkFieldsTicketOrder(fields, ticket, nearbyTickets)
}

func isTicketValid(filedRules [][]int, ticket []int) bool {
	for _, fieldValue := range ticket {
		if !isTicketFieldValid(fieldValue, filedRules) {
			return false
		}
	}
	return true
}

func checkFieldsTicketOrder(fieldRules [][]int, ticket []int, nearbyTickets [][]int) {
	matrix := buildMatrixWithZeros(len(fieldRules))

	for rank, fieldValue := range ticket {
		possibleFields := getPossibleFieldsForValue(fieldValue, fieldRules)
		for _, field := range possibleFields {
			matrix[rank][field] = 1
		}
	}

	for _, nearbyTicket := range nearbyTickets {
		if !isTicketValid(fieldRules, nearbyTicket) {
			continue
		}

		otherMatrix := buildMatrixWithZeros(len(fieldRules))
		for rank, fieldValue := range nearbyTicket {
			possibleFields := getPossibleFieldsForValue(fieldValue, fieldRules)
			for _, field := range possibleFields {
				otherMatrix[rank][field] = 1
			}
		}
		mergeMatrix(matrix, otherMatrix)
	}

	// for i
	fmt.Println("=========")
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}

	reduced := true
	for reduced {
		reduced = false
		for col := 0; col < len(matrix); col++ {
			rank, index := columnRank(matrix, col)
			if rank == 1 {
				for j := 0; j < len(matrix); j++ {
					if j != col && matrix[index][j] == 1 {
						reduced = true
						matrix[index][j] = 0
					}
				}
			}
		}
	}

	fmt.Println("=========")
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}

	res := 1
	for field := 0; field < 6; field++ {
		_, order := columnRank(matrix, field)
		fmt.Printf("field %d is at order %d (value on your ticket: %d)\n", field, order, ticket[order])
		res *= ticket[order]
	}
	fmt.Printf("Result: %d\n", res)
}

func buildMatrixWithZeros(dim int) [][]int {
	matrix := [][]int{}
	for i := 0; i < dim; i++ {
		row := []int{}
		for j := 0; j < dim; j++ {
			row = append(row, 0)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func mergeMatrix(matrix [][]int, other [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 || other[i][j] == 0 {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = 1
			}
		}
	}
}

func columnRank(matrix [][]int, col int) (int, int) {
	firstIndex := -1
	rank := 0
	for idx, row := range matrix {
		rank += row[col]
		if rank == 1 && firstIndex == -1 {
			firstIndex = idx
		}
	}
	return rank, firstIndex
}

func getPossibleFieldsForValue(fieldValue int, fieldRules [][]int) []int {
	results := []int{}
	for idx, rule := range fieldRules {
		if fieldValue >= rule[0] && fieldValue <= rule[1] {
			results = append(results, idx)
			continue
		}
		if fieldValue >= rule[2] && fieldValue <= rule[3] {
			results = append(results, idx)
		}
	}
	return results
}

// Data parsing ----------

func parseFullTicket(lines []string) ([][]int, []int, [][]int) {
	fields, idx := parseFields(lines)
	ticket, idx := parseDay16Ticket(idx, lines)
	nearbyTickets := parseDay16NearbyTickets(idx, lines)
	// fmt.Printf("%v\n", fields)
	// fmt.Printf("%v\n", ticket)
	// fmt.Printf("%v\n", nearbyTickets)
	// fmt.Println("==========")
	return fields, ticket, nearbyTickets
}

func parseFields(lines []string) ([][]int, int) {
	idx := 0
	fieldRules := [][]int{}
	for lines[idx] != "" {
		fields := strings.Split(lines[idx], ": ")
		fields = strings.Split(fields[1], " or ")
		fieldValues := []int{}
		for i := 0; i < 2; i++ {
			values := strings.Split(fields[i], "-")
			for j := 0; j < 2; j++ {
				v, _ := strconv.Atoi(values[j])
				fieldValues = append(fieldValues, v)
			}
		}
		fieldRules = append(fieldRules, fieldValues)
		idx++
	}
	return fieldRules, (idx + 1)
}

func parseDay16Ticket(idx int, lines []string) ([]int, int) {
	for lines[idx] != "your ticket:" {
		idx++
	}
	ticket := []int{}
	arr := strings.Split(lines[idx+1], ",")
	for _, a := range arr {
		i, _ := strconv.Atoi(a)
		ticket = append(ticket, i)
	}

	return ticket, idx + 2
}

func parseDay16NearbyTickets(idx int, lines []string) [][]int {
	for lines[idx] != "nearby tickets:" {
		idx++
	}
	idx++

	nearbyTickets := [][]int{}
	for idx < len(lines) {
		ticket := []int{}
		arr := strings.Split(lines[idx], ",")
		for _, a := range arr {
			i, _ := strconv.Atoi(a)
			ticket = append(ticket, i)
		}
		nearbyTickets = append(nearbyTickets, ticket)
		idx++
	}
	return nearbyTickets
}
