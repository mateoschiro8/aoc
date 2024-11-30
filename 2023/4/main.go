package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
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

	counts := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		counts[i] = 1
	}

	for i, card := range input {
		p1 := pointsForCard(card, i, counts)
		sum1 += p1
		sum2 += counts[i]
	}

	fmt.Println("Primera parte:", sum1)
	fmt.Println("Segunda parte:", sum2)

}

func pointsForCard(card string, cardIndex int, counts []int) int {

	_, numbers, _ := strings.Cut(card, ":")
	winNumbers, numbers, _ := strings.Cut(numbers, "|")

	numbersIHave := strings.Fields(numbers)
	winningNumbers := strings.Fields(winNumbers)

	var sum int
	for _, v := range numbersIHave {
		if slices.Contains(winningNumbers, v) {
			sum++
			// fmt.Println("nro ganador:", v)
		}
	}

	for k := cardIndex + 1; k <= cardIndex+sum; k++ {
		counts[k] += counts[cardIndex]
	}

	return max(MathPow(2, sum-1), 0)

}

func MathPow(n, m int) int {
	return int(math.Pow(float64(n), float64(m)))
}
