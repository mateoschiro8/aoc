package main

import (
	"bufio"
	"fmt"
	"os"
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

	var idValidGames, powerSets int
	for fileScanner.Scan() {
		id, ps := gameData(fileScanner.Text())
		idValidGames += id
		powerSets += ps
	}

	readFile.Close()
	fmt.Println("Parte 1: ", idValidGames)
	fmt.Println("Parte 2: ", powerSets)
}

func gameData(game string) (int, int) {

	gameID, gameInfo, _ := strings.Cut(game, ":")
	gameID, _ = strings.CutPrefix(gameID, "Game ")
	idIfValid, _ := strconv.Atoi(gameID)

	var maxR, maxG, maxB int

	rounds := strings.Split(gameInfo, ";")
	for _, round := range rounds {

		cubes := strings.Split(round, ",")
		for _, v := range cubes {

			hand := strings.Split(v, " ")

			q, _ := strconv.Atoi(hand[1])
			c := hand[2]

			// fmt.Println("Mostró", q, "cubos de color", c)

			switch c {
			case "red":
				maxR = max(q, maxR)
				if q > 12 {
					idIfValid = 0
				}
			case "green":
				maxG = max(q, maxG)
				if q > 13 {
					idIfValid = 0
				}
			case "blue":
				maxB = max(q, maxB)
				if q > 14 {
					idIfValid = 0
				}
			}

		}

	}

	return idIfValid, maxR * maxG * maxB
}
