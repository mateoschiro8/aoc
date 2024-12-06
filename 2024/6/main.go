package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Point struct {
	dir int
	x   int
	y   int
}

func main() {

	readFile, err := os.Open("input.txt")

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
	var pos1 []Point

	for i, v := range input {
		for j := range v {
			if input[i][j] == '^' {
				startI, startJ = i, j
			}
		}
	}

	step(input, startI, startJ, Point{dir: 1, x: -1, y: 0}, &pos1, true)
	sum1 = len(pos1)

	for _, v := range pos1 {
		i, j := v.x, v.y
		if input[i][j] != '^' {
			var pos2 []Point
			tmp1 := input[i][j]
			input[i][j] = '#'
			step(input, startI, startJ, Point{dir: 1, x: -1, y: 0}, &pos2, false)
			input[i][j] = tmp1
			if len(pos2) == 0 {
				sum2++
			}
		}

	}

	return sum1, sum2
}

// 1 arriba, 2 derecha, 3 abajo, 4 izquierda
func step(input [][]byte, i, j int, point Point, pos *[]Point, w bool) {

	if !inBounds(input, i, j) {
		return
	}

	if input[i][j] == '#' {
		switch point.dir {
		case 1:
			step(input, i+1, j+1, Point{dir: 2, x: 0, y: 1}, pos, w)
		case 2:
			step(input, i+1, j-1, Point{dir: 3, x: 1, y: 0}, pos, w)
		case 3:
			step(input, i-1, j-1, Point{dir: 4, x: 0, y: -1}, pos, w)
		case 4:
			step(input, i-1, j+1, Point{dir: 1, x: -1, y: 0}, pos, w)
		}
	} else {
		p := Point{point.dir, i, j}
		if w {
			if !slices.ContainsFunc(*pos, func(p1 Point) bool { return p1.x == p.x && p1.y == p.y }) {
				*pos = append(*pos, p)
			}
		} else {
			if slices.Contains(*pos, p) {
				*pos = slices.Delete(*pos, 0, len(*pos))
				return
			}
			*pos = append(*pos, p)
		}
		step(input, i+point.x, j+point.y, point, pos, w)
	}
}

func inBounds(input [][]byte, i, j int) bool {
	return 0 <= i && i < len(input) && 0 <= j && j < len(input[0])
}
