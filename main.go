package main

import (
	"flag"
	"fmt"
)

var day int
var part int

func parseArgs() {
	flag.IntVar(&day, "day", 1, "day of the puzzle")
	flag.IntVar(&part, "part", 1, "part of the puzzle (1 or 2)")
	flag.Parse()
}

func main() {
	parseArgs()
	fmt.Println("Day", day)

	problems := [][]func(){
		[]func(){solveDay01Example, solveDay01Part1, solveDay01Part2},
		[]func(){solveDay02Example, solveDay02Part1, solveDay02Part2},
		[]func(){solveDay03Example, solveDay03Part1, solveDay03Part2},
		[]func(){solveDay04Example, solveDay04Part1, solveDay04Part2},
		[]func(){solveDay05Example, solveDay05Part1, solveDay05Part2},
		[]func(){solveDay06Example, solveDay06Part1, solveDay06Part2},
		[]func(){solveDay07Example, solveDay07Part1, solveDay07Part2},
		[]func(){solveDay08Example, solveDay08Part1, solveDay08Part2},
		[]func(){solveDay09Example, solveDay09Part1, solveDay09Part2},
		[]func(){solveDay10Example, solveDay10Part1, solveDay10Part2},
		[]func(){solveDay11Example, solveDay11Part1, solveDay11Part2},
		[]func(){solveDay12Example, solveDay12Part1, solveDay12Part2},
		[]func(){solveDay13Example, solveDay13Part1, solveDay13Part2},
		[]func(){solveDay14Example, solveDay14Part1, solveDay14Part2},
		[]func(){solveDay15Example, solveDay15Part1, solveDay15Part2},
		[]func(){solveDay16Example, solveDay16Part1, solveDay16Part2},
		[]func(){solveDay17Example, solveDay17Part1, solveDay17Part2},
		[]func(){solveDay18Example, solveDay18Part1, solveDay18Part2},
		[]func(){solveDay19Example, solveDay19Part1, solveDay19Part2},
		[]func(){solveDay20Example, solveDay20Part1, solveDay20Part2},
		[]func(){solveDay21Example, solveDay21Part1, solveDay21Part2},
		[]func(){solveDay22Example, solveDay22Part1, solveDay22Part2},
		[]func(){solveDay23Example, solveDay23Part1, solveDay23Part2},
		[]func(){solveDay24Example, solveDay24Part1, solveDay24Part2},
		[]func(){solveDay25Example, solveDay25Part1, solveDay25Part2},
	}

	if day > 0 && day <= len(problems) && part >= 0 && part < len(problems[day-1]) {
		problems[day-1][part]()
	} else {
		fmt.Println("No puzzle for this day (or this part of day), yet!")
	}
}
