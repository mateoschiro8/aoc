package main

import (
	"bufio"
	"fmt"
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

	sum1, sum2 = findGoodUpdates(input)

	fmt.Println("Primera parte:", sum1)
	fmt.Println("Segunda parte:", sum2)
}

func findGoodUpdates(input []string) (int, int) {

	rules := map[int][]int{}
	pages := [][]int{}
	var sum1, sum2, tmp int

	findRulesAndPages(input, rules, &pages)

	for _, page := range pages {

		tmp = page[len(page)/2]
		for z, v := range page {
			found := false
			r := rules[v]
			for i := 0; i < z; i++ {
				if slices.Contains(r, page[i]) && !found {
					tmp = 0
					sum2 += correctPage(page, rules)
					found = true
				}
			}
			if found {
				break
			}
		}
		sum1 += tmp
	}
	return sum1, sum2
}

func correctPage(page []int, rules map[int][]int) int {

	fixed := slices.Clone(page)

	// 1 si a < b
	slices.SortFunc(fixed, func(a, b int) int {
		if slices.Contains(rules[a], b) {
			return 1
		}
		return -1
	})

	return fixed[len(fixed)/2]
}

func findRulesAndPages(input []string, rules map[int][]int, pages *[][]int) {

	var a, b, c int
	for _, v := range input {
		if strings.Contains(v, "|") {
			fmt.Sscanf(v, "%d|%d", &a, &b)
			rules[a] = append(rules[a], b)
		}

		if strings.Contains(v, ",") {
			values := strings.Split(v, ",")
			page := []int{}
			for _, z := range values {
				fmt.Sscanf(z, "%d", &c)
				page = append(page, c)
			}
			// fmt.Println(page)
			*pages = append(*pages, page)
		}
	}
}
