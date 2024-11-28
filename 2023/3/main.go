package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var schem [][]rune

	for fileScanner.Scan() {
		schem = append(schem, []rune(fileScanner.Text()))
	}

	readFile.Close()

	sum, ratios := checkParts(schem)

	fmt.Println("Primera parte:", sum)
	fmt.Println("Segunda parte:", ratios)

}

func checkParts(schem [][]rune) (int, int) {

	var sum, ratios int
	gears := make(map[string][]int)

	for i, v := range schem {
		for j := 0; j < len(v); j++ {

			if unicode.IsNumber(schem[i][j]) {
				newSearchPos, lastPos, number := extractNumber(schem[i], j)
				if isNearSymbol(schem, i, j, lastPos, number, gears) {
					sum += number
					// fmt.Println("El numero", number, "esta cerca de un símbolo")
				}
				j = newSearchPos
			}
		}
	}

	for _, v := range gears {
		if len(v) == 2 {
			// fmt.Println("Son parte los numeros", v[0], "y", v[1])
			ratios += (v[0] * v[1])
		}
	}

	return sum, ratios
}

func extractNumber(row []rune, j int) (int, int, int) {

	var number, k int
	for k = j; k < len(row) && unicode.IsNumber(row[k]); k++ {
		number = number*10 + int(row[k]-'0')
	}
	// fmt.Println(k, k-1, number)
	return k, k - 1, number
}

func isNearSymbol(schem [][]rune, row, startPos, lastPos, number int, gears map[string][]int) bool {

	for pos := startPos; pos <= lastPos; pos++ {

		for i := row - 1; i <= row+1; i++ {
			for j := pos - 1; j <= pos+1; j++ {

				if i >= 0 && j >= 0 && i < len(schem) && j < len(schem[0]) && !unicode.IsNumber(schem[i][j]) && schem[i][j] != '.' {
					if schem[i][j] == '*' {
						gears[string(i)+string(j)] = append(gears[string(i)+string(j)], number)
					}
					return true
				}
			}
		}
	}

	return false
}
