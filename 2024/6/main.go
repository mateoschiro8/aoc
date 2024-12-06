package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func main() {

	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var sum1, sum2 int
	var input []string
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	readFile.Close()

	res := make([][]byte, len(input))
	for i, str := range input {
		res[i] = []byte(str)
	}
	sum1, sum2 = countSteps(res)

	fmt.Println("Primera parte:", sum1)
	fmt.Println("Segunda parte:", sum2)
}

func countSteps(input [][]byte) (int, int) {

	var sum1, sum2, startI, startJ int

	pos := make([][]int, len(input))

	for i, v := range input {

		pos[i] = make([]int, len(v))

		for j := range v {
			if input[i][j] == '^' {
				startI, startJ = i, j
			}
		}
	}

	sum1 = step(input, startI, startJ, 1, Point{x: -1, y: 0}, pos)
	/*
		for i, v := range input {
			for j := range v {
				if input[i][j] != '^' {
					tmp1 := input[i][j]
					input[i][j] = '#'
					step(input, startI, startJ, 1, Point{x: -1, y: 0}, pos)
					input[i][j] = tmp1
				}
			}
		}
	*/
	return sum1, sum2
}

// 1 arriba, 2 derecha, 3 abajo, 4 izquierda
func step(input [][]byte, i, j, heading int, dir Point, pos [][]int) int {

	/*
		for _, v := range input {
			fmt.Println(string(v))
		}
		fmt.Println()
	*/
	if !inBounds(input, i, j) {
		return 0
	}

	if input[i][j] == '#' {
		switch heading {
		case 1:
			return step(input, i+1, j+1, 2, Point{x: 0, y: 1}, pos)
		case 2:
			return step(input, i+1, j-1, 3, Point{x: 1, y: 0}, pos)
		case 3:
			return step(input, i-1, j-1, 4, Point{x: 0, y: -1}, pos)
		case 4:
			return step(input, i-1, j+1, 1, Point{x: -1, y: 0}, pos)
		}
	} else {
		if pos[i][j] == 1 {
			return step(input, i+dir.x, j+dir.y, heading, dir, pos)
		} else {
			pos[i][j] = 1
			return 1 + step(input, i+dir.x, j+dir.y, heading, dir, pos)
		}
	}
	return 0
}

func inBounds(input [][]byte, i, j int) bool {
	return 0 <= i && i < len(input) && 0 <= j && j < len(input[0])
}
