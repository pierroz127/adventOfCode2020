package main

import "fmt"

func solveDay17Example() {
	lines := []string{
		".#.",
		"..#",
		"###",
	}
	fmt.Println("Solving day 17 part 2 with example data")
	hypercube := convertToHypercube(convertTextToCube(lines))
	printHypercube(hypercube)
	for i := 0; i < 6; i++ {
		hypercube = processHypercubeCycle(hypercube)
	}
	fmt.Printf("Number of active cells: %d\n", countHypercubeActiveCells(hypercube))
}

func solveDay17Part1() {
	cube := getDay17Data()
	printCube(cube)
	for i := 0; i < 6; i++ {
		cube = processCubeCycle(cube)
	}
	fmt.Printf("Number of active cells: %d\n", countActiveCells(cube))
}

func processCubeCycle(cube [][][]byte) [][][]byte {
	expandedCube := expandCube(cube)
	newcube := [][][]byte{}

	for z := 0; z < len(expandedCube); z++ {
		square := [][]byte{}
		for y := 0; y < len(expandedCube[z]); y++ {
			row := []byte{}
			for x := 0; x < len(expandedCube[z][y]); x++ {
				count := countNearbyActiveCells(expandedCube, x, y, z)
				cell := expandedCube[z][y][x]
				if cell == 1 {
					if count != 2 && count != 3 {
						cell = 0
					}
				} else {
					if count == 3 {
						cell = 1
					}
				}
				row = append(row, cell)
			}
			square = append(square, row)
		}
		newcube = append(newcube, square)
	}
	return newcube
}

func expandCube(cube [][][]byte) [][][]byte {
	expandedCube := [][][]byte{}
	expandedCube = append(expandedCube, buildSquaresWithZeros(len(cube[0])+2, len(cube[0][0])+2))
	for z := 0; z < len(cube); z++ {
		square := [][]byte{}
		for y := -1; y < len(cube[0])+1; y++ {
			row := []byte{}
			for x := -1; x < len(cube[0][0])+1; x++ {
				if y == -1 || x == -1 || y == len(cube[z]) || x == len(cube[z][y]) {
					row = append(row, 0)
				} else {
					row = append(row, cube[z][y][x])
				}
			}
			square = append(square, row)
		}
		expandedCube = append(expandedCube, square)
	}
	expandedCube = append(expandedCube, buildSquaresWithZeros(len(cube[0])+2, len(cube[0][0])+2))

	return expandedCube
}

func buildSquaresWithZeros(dimy int, dimx int) [][]byte {
	square := [][]byte{}
	for y := 0; y < dimy; y++ {
		row := []byte{}
		for x := 0; x < dimx; x++ {
			row = append(row, 0)
		}
		square = append(square, row)
	}
	return square
}

func countNearbyActiveCells(cube [][][]byte, x int, y int, z int) int {
	count := 0

	for zi := max(0, z-1); zi <= min(z+1, len(cube)-1); zi++ {
		for yj := max(0, y-1); yj <= min(y+1, len(cube[z])-1); yj++ {
			for xk := max(0, x-1); xk <= min(x+1, len(cube[z][y])-1); xk++ {
				if xk != x || yj != y || zi != z {
					count += int(cube[zi][yj][xk])
				}
			}
		}
	}
	return count
}

func printCube(cube [][][]byte) {
	for z := 0; z < len(cube); z++ {
		fmt.Printf("z=%d\n", z)
		for _, row := range cube[z] {
			fmt.Printf("%v\n", row)
		}
	}
}

func countActiveCells(cube [][][]byte) int {
	count := 0
	for z := 0; z < len(cube); z++ {
		for y := 0; y < len(cube[z]); y++ {
			for x := 0; x < len(cube[z][y]); x++ {
				count += int(cube[z][y][x])
			}
		}
	}
	return count
}

func solveDay17Part2() {
	hypercube := convertToHypercube(getDay17Data())
	printHypercube(hypercube)
	for i := 0; i < 6; i++ {
		hypercube = processHypercubeCycle(hypercube)
	}
	fmt.Printf("Number of active cells: %d\n", countHypercubeActiveCells(hypercube))
}

func buildCubeWithZeros(dimz int, dimy int, dimx int) [][][]byte {
	cube := [][][]byte{}
	for z := 0; z < dimz; z++ {
		square := [][]byte{}
		for y := 0; y < dimy; y++ {
			row := []byte{}
			for x := 0; x < dimx; x++ {
				row = append(row, 0)
			}
			square = append(square, row)
		}
		cube = append(cube, square)
	}
	return cube
}

func expandHypercube(hypercube [][][][]byte) [][][][]byte {
	expandedHypercube := [][][][]byte{}
	expandedHypercube = append(
		expandedHypercube,
		buildCubeWithZeros(len(hypercube[0])+2, len(hypercube[0][0])+2, len(hypercube[0][0][0])+2))

	for w := 0; w < len(hypercube); w++ {
		cube := expandCube(hypercube[w])
		expandedHypercube = append(expandedHypercube, cube)
	}

	return append(
		expandedHypercube,
		buildCubeWithZeros(len(hypercube[0])+2, len(hypercube[0][0])+2, len(hypercube[0][0][0])+2))
}

func countNearbyActiveHypercubeCells(cube [][][][]byte, x int, y int, z int, w int) int {
	count := 0
	for wt := max(0, w-1); wt <= min(w+1, len(cube)-1); wt++ {
		for zi := max(0, z-1); zi <= min(z+1, len(cube[w])-1); zi++ {
			for yj := max(0, y-1); yj <= min(y+1, len(cube[w][z])-1); yj++ {
				for xk := max(0, x-1); xk <= min(x+1, len(cube[w][z][y])-1); xk++ {
					if xk != x || yj != y || zi != z || wt != w {
						count += int(cube[wt][zi][yj][xk])
					}
				}
			}
		}
	}
	return count
}

func processHypercubeCycle(cube [][][][]byte) [][][][]byte {
	expandedHypercube := expandHypercube(cube)
	newHypercube := [][][][]byte{}
	for w := 0; w < len(expandedHypercube); w++ {
		cube := [][][]byte{}
		for z := 0; z < len(expandedHypercube[w]); z++ {
			square := [][]byte{}
			for y := 0; y < len(expandedHypercube[w][z]); y++ {
				row := []byte{}
				for x := 0; x < len(expandedHypercube[w][z][y]); x++ {
					count := countNearbyActiveHypercubeCells(expandedHypercube, x, y, z, w)
					cell := expandedHypercube[w][z][y][x]
					if cell == 1 {
						if count != 2 && count != 3 {
							cell = 0
						}
					} else {
						if count == 3 {
							cell = 1
						}
					}
					row = append(row, cell)
				}
				square = append(square, row)
			}
			cube = append(cube, square)
		}
		newHypercube = append(newHypercube, cube)
	}
	return newHypercube
}

func printHypercube(hypercube [][][][]byte) {
	for w := 0; w < len(hypercube); w++ {
		for z := 0; z < len(hypercube[w]); z++ {
			fmt.Printf("z=%d, w=%d\n", z, w)
			for _, row := range hypercube[w][z] {
				fmt.Printf("%v\n", row)
			}
		}
	}
}

func countHypercubeActiveCells(cube [][][][]byte) int {
	count := 0
	for w := 0; w < len(cube); w++ {
		for z := 0; z < len(cube[w]); z++ {
			for y := 0; y < len(cube[w][z]); y++ {
				for x := 0; x < len(cube[w][z][y]); x++ {
					count += int(cube[w][z][y][x])
				}
			}
		}
	}
	return count
}

// ----------

func getDay17Data() [][][]byte {
	lines := getDataFromFile("day17")
	return convertTextToCube(lines)
}

func convertTextToCube(lines []string) [][][]byte {
	square := [][]byte{}
	for _, line := range lines {
		row := []byte{}
		for i := 0; i < len(line); i++ {
			if line[i] == '.' {
				row = append(row, 0)
			} else {
				row = append(row, 1)
			}
		}
		square = append(square, row)
	}
	return [][][]byte{square}
}

func convertToHypercube(cube [][][]byte) [][][][]byte {
	return [][][][]byte{cube}
}
