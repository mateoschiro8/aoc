package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var total int = 0

	for fileScanner.Scan() {
		total += processLine2(fileScanner.Text())
	}

	readFile.Close()

	fmt.Println(total)

}

func processLine1(line string) int {

	var first, last int
	for _, v := range line {
		if unicode.IsNumber(v) {
			if first == 0 {
				first = int(v - '0')
			}
			last = int(v - '0')
		}
	}
	return first*10 + last
}

func tieneNumero(line string) (int, bool) {
	switch {
	case unicode.IsNumber(rune(line[0])):
		return int(line[0] - '0'), true
	case strings.HasPrefix(line, "one"):
		return 1, true
	case strings.HasPrefix(line, "two"):
		return 2, true
	case strings.HasPrefix(line, "three"):
		return 3, true
	case strings.HasPrefix(line, "four"):
		return 4, true
	case strings.HasPrefix(line, "five"):
		return 5, true
	case strings.HasPrefix(line, "six"):
		return 6, true
	case strings.HasPrefix(line, "seven"):
		return 7, true
	case strings.HasPrefix(line, "eight"):
		return 8, true
	case strings.HasPrefix(line, "nine"):
		return 9, true
	default:
		return 0, false
	}

}

func processLine2(line string) int {

	var first, last int
	for i := range line {
		value, ok := tieneNumero(line[i:])
		if ok {
			if first == 0 {
				first = value
			}
			last = value
		}
	}

	return first*10 + last
}
