/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2022 - Go Edition
	 o'.o.'.o.     day4!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4(filename string) {
	fmt.Printf("AoC day4 2022\n\n")

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
	fmt.Println(getCleanersCompletelyOverlapSum(input))

	// Part 2
	fmt.Println(getCleanersOverlapSum(input))

}

type cleaner struct {
	lowerBound int
	upperBound int
}

func cleanerParser(line string) []cleaner {
	var cleanerPair []cleaner
	twoCleaners := strings.Split(line, ",")
	for _, c := range twoCleaners {
		bounds := strings.Split(c, "-")
		lwr, _ := strconv.Atoi(bounds[0])
		upr, _ := strconv.Atoi(bounds[1])

		c := cleaner{
			lowerBound: lwr,
			upperBound: upr,
		}

		cleanerPair = append(cleanerPair, c)
	}

	return cleanerPair
}

func doCleanersCompletelyOverlap(c1, c2 cleaner) bool {
	if (c1.lowerBound <= c2.lowerBound && c1.upperBound >= c2.upperBound) ||
		(c2.lowerBound <= c1.lowerBound && c2.upperBound >= c1.upperBound) {
		return true
	}
	return false
}

func isWithin(i int, lower int, upper int) bool {
	if i >= lower && i <= upper {
		return true
	}
	return false
}

func doCleanersOverlap(c1, c2 cleaner) bool {
	if isWithin(c1.lowerBound, c2.lowerBound, c2.upperBound) ||
		isWithin(c1.upperBound, c2.lowerBound, c2.upperBound) ||
		isWithin(c2.lowerBound, c1.lowerBound, c1.upperBound) ||
		isWithin(c2.upperBound, c1.lowerBound, c1.upperBound) {
		return true
	}
	return false
}

func getCleanersOverlapSum(data []string) int {
	sum := 0
	for _, line := range data {
		cleanerPair := cleanerParser(line)
		if doCleanersOverlap(cleanerPair[0], cleanerPair[1]) {
			sum++
		}
	}
	return sum
}

func getCleanersCompletelyOverlapSum(data []string) int {
	sum := 0
	for _, line := range data {
		cleanerPair := cleanerParser(line)
		if doCleanersCompletelyOverlap(cleanerPair[0], cleanerPair[1]) {
			sum++
		}
	}
	return sum
}
