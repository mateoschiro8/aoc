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

	sum1, sum2 = checkSafeReports(input)

	fmt.Println("Primera parte:", sum1)
	fmt.Println("Segunda parte:", sum2)
}

func checkSafeReports(input []string) (int, int) {

	var sum1, sum2 int
	for _, v := range input {

		var rp []int
		for _, v := range strings.Fields(v) {
			value, _ := strconv.Atoi(v)
			rp = append(rp, value)
		}

		sum1 += isSafeReport(rp)
		sum2 += isSafeWithTol(rp)
	}

	return sum1, sum2
}

func isSafeReport(report []int) int {

	safe := 1
	for i := 0; i < len(report)-1; i++ {
		dif := report[i+1] - report[i]
		if dif < 1 || dif > 3 {
			safe = 0
		}
	}

	if safe == 0 {
		slices.Reverse(report)
		safe = 1
		for i := 0; i < len(report)-1; i++ {
			dif := report[i+1] - report[i]
			if dif < 1 || dif > 3 {
				safe = 0
			}
		}
	}

	return safe
}

func isSafeWithTol(report []int) int {

	var sum int
	for i := range report {
		tmp := slices.Delete(slices.Clone(report), i, i+1)
		if isSafeReport(tmp) == 1 {
			sum++
			break
		}
	}

	return sum
}
