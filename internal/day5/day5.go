/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2022 - Go Edition
	 o'.o.'.o.     day5!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"676f.dev/utilities/structures/stack"
)

func Day5(filename string) {
	fmt.Printf("AoC day5 2022\n\n")

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
	emptyLine := 0
	for idx, line := range input {
		if line == "" {
			emptyLine = idx
			break
		}
	}

	crates := input[:emptyLine-1]
	commands := input[emptyLine+1:]

	parsedCrates := crateParser(crates)

	parsedCommands := commandParser(commands)

	finalCrates := commandEngine(parsedCrates, parsedCommands)

	for _, crate := range finalCrates {
		fmt.Print(crate.Pop())
	}
	fmt.Println()
}

func part2(input []string) {
	emptyLine := 0
	for idx, line := range input {
		if line == "" {
			emptyLine = idx
			break
		}
	}

	crates := input[:emptyLine-1]
	commands := input[emptyLine+1:]

	parsedCrates := crateParser(crates)

	parsedCommands := commandParser(commands)

	finalCrates := commandEnginePreserveOrder(parsedCrates, parsedCommands)

	for _, crate := range finalCrates {
		fmt.Print(crate.Pop())
	}
	fmt.Println()
}

func crateParser(inp []string) []*stack.Stack {

	var table [][]string
	var row []string
	var tmp string

	for _, line := range inp {
		for idx, c := range line {
			if unicode.IsLetter(c) {
				tmp += string(c)
			}
			if idx%4 == 0 || idx == len(line)-1 {
				row = append(row, string(tmp))
				tmp = ""
			}
		}
		table = append(table, row)
		row = nil
	}

	var crates []*stack.Stack
	tableLen := len(table) - 1
	rowLen := len(table[0])
	for j := 1; j < rowLen; j++ {
		crate := stack.NewStack()
		for i := tableLen; i >= 0; i-- {
			if len(table[i][j]) != 0 {
				if unicode.IsLetter(rune(table[i][j][0])) {
					crate.Push(table[i][j])
				}
			}
		}
		crates = append(crates, crate)
	}

	return crates
}

type command struct {
	count int
	from  int
	to    int
}

func commandParser(inp []string) []command {
	var commands []command
	for _, line := range inp {
		splits := strings.Split(line, " ")
		count, _ := strconv.Atoi(splits[1])
		from, _ := strconv.Atoi(splits[3])
		to, _ := strconv.Atoi(splits[5])
		commands = append(commands, command{count, from - 1, to - 1})
	}
	return commands
}

func commandEngine(crates []*stack.Stack, commands []command) []*stack.Stack {
	for _, cmd := range commands {
		for i := 0; i < cmd.count; i++ {
			crates[cmd.to].Push(crates[cmd.from].Pop())
		}
	}

	return crates
}

func commandEnginePreserveOrder(crates []*stack.Stack, commands []command) []*stack.Stack {
	for _, cmd := range commands {
		tmpStack := stack.NewStack()
		for i := 0; i < cmd.count; i++ {
			tmpStack.Push(crates[cmd.from].Pop())
		}

		for tmpStack.Len() != 0 {
			crates[cmd.to].Push(tmpStack.Pop())
		}
	}
	return crates
}
