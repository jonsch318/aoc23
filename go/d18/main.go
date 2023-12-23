package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jonsch318/aoc23/go/d18/solve"
)

const INPUT = "input/d18/input"
const TEST = "input/d18/test"

//not 53019
//not 46067
//correct 92758

const (
	dirR byte = iota
	dirD byte = iota
	dirL byte = iota
	dirU byte = iota
)

func main() {

	commands := setup(INPUT, readP2)
	/* for _, c := range commands {
		log.Println(c)
	} */

	area := solve.Solve(commands)
	log.Printf("Area %v", area)
}

func readP1(file *os.File) (int, solve.CMD) {

	var l, len int
	var d rune

	_, err := fmt.Fscanf(file, "%c %d (#%x)\n", &d, &len, &l)

	if err != nil {
		return 0, solve.CMD{}
	}

	switch d {
	case 'R':
		return 1, solve.CMD{Dir: dirR, Len: len}
	case 'L':
		return 1, solve.CMD{Dir: dirL, Len: len}
	case 'U':
		return 1, solve.CMD{Dir: dirU, Len: len}
	case 'D':
		return 1, solve.CMD{Dir: dirD, Len: len}
	}
	return 0, solve.CMD{}
}

func readP2(file *os.File) (int, solve.CMD) {
	var cmd solve.CMD
	var l int

	_, err := fmt.Fscanf(file, "%c %d (#%x)\n", &cmd.Dir, &l, &cmd.Len)

	if err != nil {
		return 0, solve.CMD{}
	}

	cmd.Dir = byte(cmd.Len & 3)
	cmd.Len = cmd.Len >> 4
	return 1, cmd
}

func setup(path string, read func(*os.File) (int, solve.CMD)) []solve.CMD {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	commands := make([]solve.CMD, 0, 1000)

	for {
		ret, cmd := read(file)
		if ret != 1 {
			break
		}
		commands = append(commands, cmd)

	}

	return commands
}
