/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2022 - Go Edition
	 o'.o.'.o.     day3!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day3(filename string) {
	fmt.Printf("AoC day3 2022\n\n")

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
	fmt.Println(getIndividualSum(input))

	// Part 2
	fmt.Println(getGroupSum(input))

}

func getCommonPriority(c rune) int {

	if isRuneUpperCase(c) {
		// ASCII value mapping to 27-52
		return int(c) - 64 + 26
	} else {
		// ASCII value mapping to 1-26
		return int(c) - 96
	}
}

func getCommonCharacter(s string) rune {
	firstHalf := s[:len(s)/2]
	secondHalf := s[len(s)/2:]

	for _, c := range firstHalf {
		if strings.Contains(secondHalf, string(c)) {
			return c
		}
	}

	// should never get here
	return 0
}

func isRuneUpperCase(c rune) bool {
	return string(c) == strings.ToUpper(string(c))
}

func getIndividualSum(data []string) int {
	sum := 0
	for _, line := range data {
		c := getCommonCharacter(line)
		sum += getCommonPriority(c)
	}
	return sum
}

func getGroupSum(data []string) int {
	var groups [][]string
	var group []string
	for idx, line := range data {
		group = append(group, line)
		if (idx+1)%3 == 0 {
			groups = append(groups, group)
			group = []string{}
		}
	}

	sum := 0
	for _, group := range groups {
		sum += getCommonPriority(getCommonCharacterFromGroup(group))
	}

	return sum
}

func getCommonCharacterFromGroup(group []string) rune {
	for _, c := range group[0] {
		if strings.Contains(group[1], string(c)) && strings.Contains(group[2], string(c)) {
			return c
		}
	}
	return 0
}
