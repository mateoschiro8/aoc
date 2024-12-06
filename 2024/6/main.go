package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dir struct {
	heading int // 1 arriba, 2 derecha, 3 abajo, 4 izquierda
	x       int
	y       int
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

	sum1, sum2 = countSteps(input)

	fmt.Println("Primera parte:", sum1)
	fmt.Println("Segunda parte:", sum2)
}

func countSteps(input []string) (int, int) {

	var sum1, sum2, startI, startJ, a int

	pos := make([][]int, len(input))
	obs := make([][]int, len(input))

	for i, v := range input {

		pos[i] = make([]int, len(v))
		obs[i] = make([]int, len(v))

		for j := range v {
			if input[i][j] == '^' {
				startI, startJ = i, j
			}

		}

	}

	// step(input, startI, startJ, Dir{heading: 1, x: -1, y: 0}, pos, obs, &a)
	for i, v := range pos {
		// fmt.Println(v)
		for j := range v {
			sum1 += pos[i][j]
		}
	}

	for i, v := range input {
		tmp := v
		for j := range v {
			if input[i][j] != '^' {
				temp := []byte(v)
				temp[j] = '#'
				input[i] = string(temp)
				step(input, startI, startJ, Dir{heading: 1, x: -1, y: 0}, pos, obs, &a)
				input[i] = tmp
				if a > 0 {
					sum2++
				}
				a = 0
				for i1, v1 := range obs {
					for i2 := range v1 {
						obs[i1][i2] = 0
					}
				}
			}
		}
	}

	return sum1, sum2
}

func step(input []string, i, j int, dir Dir, pos [][]int, obs [][]int, a *int) {

	if !inBounds(input, i, j) {
		return
	}

	if input[i][j] == '#' {
		if obs[i][j] == 1 {
			// fmt.Println("obstaculo visto dos veces", i, j)
			*a = 1
			return
		} else {
			obs[i][j] = 1
		}
		switch dir.heading {
		case 1:
			step(input, i+1, j+1, Dir{heading: 2, x: 0, y: 1}, pos, obs, a)
		case 2:
			step(input, i+1, j-1, Dir{heading: 3, x: 1, y: 0}, pos, obs, a)
		case 3:
			step(input, i-1, j-1, Dir{heading: 4, x: 0, y: -1}, pos, obs, a)
		case 4:
			step(input, i-1, j+1, Dir{heading: 1, x: -1, y: 0}, pos, obs, a)
		}
	} else {
		pos[i][j] = 1
		step(input, i+dir.x, j+dir.y, dir, pos, obs, a)
	}
}

func inBounds(input []string, i, j int) bool {
	return 0 <= i && i < len(input) && 0 <= j && j < len(input[0])
}
