package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var idValidGames int
	for fileScanner.Scan() {
		idValidGames += isValidGame(fileScanner.Text())
	}

	readFile.Close()

}

func isValidGame(game string) int {
	gameID, gameInfo, _ := strings.Cut(game, ":")
	gameID, _ = strings.CutPrefix(gameID, "Game ")

	rounds := strings.Split(gameInfo, ";")

	for _, v := range rounds {
		if !validRound(v) {
			return 0
		}
	}

	ret, _ := strconv.Atoi(gameID)
	return ret
}

func validRound(round string) bool {
	cubes := strings.Split(round, ",")

	fmt.Println(cubes)

	for _, v := range cubes {

	}

	return true
}
