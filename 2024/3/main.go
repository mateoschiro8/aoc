package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	sum1, sum2 = findMultiplications(input)

	fmt.Println("Primera parte:", sum1)
	fmt.Println("Segunda parte:", sum2)
}

func findMultiplications(input []string) (int, int) {

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	var sum1, sum2, a, b, value int
	puedoSumar := true
	for _, l := range input {

		muls := re.FindAllString(l, -1)
		for _, m := range muls {

			value = 0
			switch m {
			case "do()":
				puedoSumar = true
			case "don't()":
				puedoSumar = false
			default:
				fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
				// fmt.Println(a, "*", b)
				value = a * b
			}

			sum1 += value

			if puedoSumar {
				sum2 += value
			}
		}
	}
	return sum1, sum2
}
