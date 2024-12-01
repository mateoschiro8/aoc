package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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

	sum1, sum2 = findDif(input)

	fmt.Println("Primera parte:", sum1)
	fmt.Println("Segunda parte:", sum2)
}

func findDif(input []string) (int, int) {

	var s1, s2 []int
	for _, v := range input {
		values := strings.Fields(v)
		q1, _ := strconv.Atoi(values[0])
		q2, _ := strconv.Atoi(values[1])
		s1 = append(s1, q1)
		s2 = append(s2, q2)
	}

	slices.Sort(s1)
	slices.Sort(s2)

	var sum1, sum2, tmp int
	for i, v1 := range s1 {
		val := v1 - s2[i]
		sum1 += max(val, -1*val)

		for _, v2 := range s2 {
			if v1 == v2 {
				tmp++
			}
		}

		sum2 += v1 * tmp
		tmp = 0

	}

	return sum1, sum2

}
