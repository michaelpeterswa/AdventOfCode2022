/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2022 - Go Edition
	 o'.o.'.o.     day1!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"676f.dev/utilities/algorithms/sorting"
)

func Day1(filename string) {
	fmt.Printf("AoC day1 2022\n\n")

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var sums []int
	currSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			currSum += num
		} else {
			sums = append(sums, currSum)
			currSum = 0
		}
	}

	// part 1
	fmt.Println(sliceMax(sums))

	// part 2
	fmt.Println(sliceSum(sorting.Sort(sums, true)[0:3]))
}

func sliceMax(inp []int) (int, int) {
	max := 0
	maxIndex := 0
	for i, v := range inp {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	return maxIndex, max
}

func sliceSum(inp []int) int {
	sum := 0
	for _, v := range inp {
		sum += v
	}
	return sum
}
