package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	sum1, sum2 = buscarXMAS(input)

	fmt.Println("Primera parte:", sum1)
	fmt.Println("Segunda parte:", sum2)
}

func buscarXMAS(input []string) (int, int) {

	var sum1, sum2 int
	for i, s := range input {
		for j, v := range s {

			if v == 'X' {
				for a := -1; a <= 1; a++ {
					for b := -1; b <= 1; b++ {
						sum1 += buscarEnDir(input, i, j, a, b, "MAS")
					}
				}
			}

			if v == 'A' {
				sum2 += esXMAS(input, i, j)
			}

		}
	}

	return sum1, sum2
}

func buscarEnDir(input []string, i, j, a, b int, s string) int {

	if s == "" {
		return 1
	}

	newI, newJ := i+a, j+b
	if !inBounds(input, newI, newJ) {
		return 0
	}

	if input[newI][newJ] == s[0] {
		return buscarEnDir(input, newI, newJ, a, b, s[1:])
	}

	return 0
}

func esXMAS(input []string, i, j int) int {

	is := []int{-1, -1, 1, 1}
	js := []int{-1, 1, 1, -1}

	var patas []byte
	for z, v := range is {

		newI, newJ := i+v, j+js[z]
		if !inBounds(input, newI, newJ) {
			return 0
		}

		val := input[newI][newJ]
		if val == 'M' || val == 'S' {
			patas = append(patas, val)
		}
	}

	if len(patas) != 4 || patas[0] == patas[2] || patas[1] == patas[3] {
		return 0
	}

	return 1
}

func inBounds(input []string, i, j int) bool {
	return 0 <= i && i < len(input) && 0 <= j && j < len(input[0])
}
