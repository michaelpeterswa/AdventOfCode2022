/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2022 - Go Edition
	 o'.o.'.o.     day2!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day2

import (
	"bufio"
	"fmt"
	"os"
)

func Day2(filename string) {
	fmt.Printf("AoC day2 2022\n\n")

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	sum1 := 0
	sum2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum1 += rps1(scanner.Text())
		sum2 += rps2(scanner.Text())
	}

	// part 1
	fmt.Println(sum1)

	// part 2
	fmt.Println(sum2)
}

func rps1(line string) int {
	moveMap := map[byte]byte{
		'A': 'Z', // rock beats scissors
		'B': 'X', // paper beats rock
		'C': 'Y', // scissors beats paper
		'X': 'C', // rock beats scissors
		'Y': 'A', // paper beats rock
		'Z': 'B', // scissors beats paper
	}

	equalMap := map[byte]byte{
		'A': 'X',
		'B': 'Y',
		'C': 'Z',
	}

	pointMap := map[byte]int{
		'X': 1, // rock
		'Y': 2, // paper
		'Z': 3, // scissors
	}

	opponent := line[0]
	you := line[2]

	// draw condition
	if you == equalMap[opponent] {
		return 3 + pointMap[you]
	}

	// win condition
	if moveMap[you] == opponent {
		return 6 + pointMap[you]
	}

	// lose condition
	return 0 + pointMap[you]
}

func rps2(line string) int {
	moveMap := map[byte]byte{
		'A': 'Z', // rock beats scissors
		'B': 'X', // paper beats rock
		'C': 'Y', // scissors beats paper
		'X': 'C', // rock beats scissors
		'Y': 'A', // paper beats rock
		'Z': 'B', // scissors beats paper
	}

	// opposite of above
	loseMap := map[byte]byte{
		'Z': 'A',
		'X': 'B',
		'Y': 'C',
		'C': 'X',
		'A': 'Y',
		'B': 'Z',
	}

	equalMap := map[byte]byte{
		'A': 'X',
		'B': 'Y',
		'C': 'Z',
	}

	pointMap := map[byte]int{
		'X': 1, // rock
		'Y': 2, // paper
		'Z': 3, // scissors
	}

	opponent := line[0]
	you := line[2]

	// draw condition
	if you == 'Y' {
		return 3 + pointMap[equalMap[opponent]]
	}

	// win condition
	if you == 'Z' {
		return 6 + pointMap[loseMap[opponent]]
	}

	// lose condition
	return 0 + pointMap[moveMap[opponent]]
}
