/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2022 - Go Edition
	 o'.o.'.o.     day6!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Day6(filename string) {
	fmt.Printf("AoC day6 2022\n\n")

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Part 1
	part1(input)

	// Part 2
	part2(input)
}

func part1(input []string) {
	for _, line := range input {
		for i := 4; i < len(line); i++ {
			if !containsDuplicate(line[i-4 : i]) {
				fmt.Println(i)
				break
			}
		}
	}
}

func part2(input []string) {
	for _, line := range input {
		for i := 14; i < len(line); i++ {
			if !containsDuplicate(line[i-14 : i]) {
				fmt.Println(i)
				break
			}
		}
	}
}

func containsDuplicate(s string) bool {
	var runeMap = make(map[rune]bool)
	for _, c := range s {
		if _, ok := runeMap[c]; !ok {
			runeMap[c] = true
		} else {
			return true
		}
	}
	return false
}
