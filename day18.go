package main

import "fmt"

type Expression struct {
	hasLeft bool;
	leftOp int;
	hasRight bool;
	rightOp int;
	operation byte;
}

type AdvancedExpression struct {
	operands []int;
	operations []byte;
}

func solveDay18Example() {
	lines := []string {
		"5 + 3",
		"1 + 2 * 3 + 4 * 5 + 6",
		"1 + (2 * 3) + (4 * (5 + 6))",
		"2 * 3 + (4 * 5)",
		 "5 + (8 * 3 + 9 + 3 * 4 * 3)",
		  "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
		 "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
	}

	for _, line := range lines {
		_, res := parseAdvancedExpression(line, 0)
		fmt.Printf("%s = %d\n",line, res)
	}
}

func solveDay18Part1() {
	lines := getDataFromFile("day18")
	finalResult := 0
	for _, line := range lines {
		_, res := parseExpression(line, 0)
		fmt.Printf("%s = %d\n",line, res)
		finalResult += res
	}
	fmt.Printf("Final result: %d\n", finalResult)
}

func parseExpression(s string, idx int) (int, int) {
	var expr = &Expression { false, 0, false, 0, ' ' }
	for idx < len(s) && s[idx] != ')' {
		if s[idx] >= '0' && s[idx] <= '9' {
			idx2, number := parseNumber(s, idx)
			idx = idx2
			expr.assign(number)
		} else if s[idx] == '+' || s[idx] == '*' {
			expr.assignOp(s[idx])
			idx++
		} else if s[idx] == '(' {
			idx++
			idx2, number := parseExpression(s, idx)
			idx = idx2
			expr.assign(number)
		} else {
			idx++
		}
	}
	res :=expr.evaluate()
	return idx+1, res
}

func parseNumber(s string, idx int) (int, int) {
	number := 0
	for idx < len(s) && s[idx] >= '0' && s[idx] <= '9' {
		number = 10*number + int(s[idx]) - int('0')
		idx++
	}
	return idx, number
}

func (expr *Expression) evaluate() int {
	if expr.operation == '+' {
		return expr.leftOp + expr.rightOp
	} else if (expr.operation == '*') {
		return expr.leftOp * expr.rightOp
	}
	return 0
}

func (expr *Expression) assign(number int) {
	if !expr.hasLeft {
		expr.leftOp = number
		expr.hasLeft = true
	} else if !expr.hasRight {
		expr.rightOp = number
		expr.hasRight = true
	} else {
		expr.leftOp = expr.evaluate()
		expr.rightOp = number
	}
}

func (expr *Expression) assignOp(op byte) {
	if expr.hasRight {
		expr.leftOp = expr.evaluate()
		expr.hasRight = false
	} 
	expr.operation = op
}

func solveDay18Part2() {
	lines := getDataFromFile("day18")
	finalResult := 0
	for _, line := range lines {
		_, res := parseAdvancedExpression(line, 0)
		fmt.Printf("%s = %d\n",line, res)
		finalResult += res
	}
	fmt.Printf("Final result: %d\n", finalResult)
}

func parseAdvancedExpression(s string, idx int) (int, int) {
	var expr = &AdvancedExpression { []int{}, []byte{} }
	for idx < len(s) && s[idx] != ')' {
		if s[idx] >= '0' && s[idx] <= '9' {
			idx2, number := parseNumber(s, idx)
			idx = idx2
			expr.assign(number)
		} else if s[idx] == '+' || s[idx] == '*' {
			expr.assignOp(s[idx])
			idx++
		} else if s[idx] == '(' {
			idx++
			idx2, number := parseAdvancedExpression(s, idx)
			idx = idx2
			expr.assign(number)
		} else {
			idx++
		}
	}
	res :=expr.evaluate()
	return idx+1, res
}

func (expr *AdvancedExpression) assign(number int) {
	expr.operands = append(expr.operands, number)
}

func (expr *AdvancedExpression) assignOp(op byte) {
	expr.operations = append(expr.operations, op)
}

func (expr *AdvancedExpression) evaluate() int {
	multoperands := []int{expr.operands[0]}
	for i, op := range expr.operations {
		if op == '+' {
			multoperands[len(multoperands)-1] += expr.operands[i+1]
		} else {
			multoperands = append(multoperands, expr.operands[i+1])
		}
	}

	res := 1
	for _, m := range multoperands {
		res *= m
	}
	return res
}