package main

import (
	"fmt"
	"log"
	"os"
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

type CMD struct {
	len int
	dir byte
}

func (c CMD) String() string {
	return fmt.Sprintf("(%d, %d)", c.dir, c.len)
}

func main() {

	commands := setup(INPUT, readP2)
	/* for _, c := range commands {
		log.Println(c)
	} */

	area := solve(commands)
	log.Printf("Area %v", area)
}

func readP1(file *os.File) (int, CMD) {

	var l, len int
	var d rune

	_, err := fmt.Fscanf(file, "%c %d (#%x)\n", &d, &len, &l)

	if err != nil {
		return 0, CMD{}
	}

	switch d {
	case 'R':
		return 1, CMD{dir: dirR, len: len}
	case 'L':
		return 1, CMD{dir: dirL, len: len}
	case 'U':
		return 1, CMD{dir: dirU, len: len}
	case 'D':
		return 1, CMD{dir: dirD, len: len}
	}
	return 0, CMD{}
}

func readP2(file *os.File) (int, CMD) {
	var cmd CMD
	var l int

	_, err := fmt.Fscanf(file, "%c %d (#%x)\n", &cmd.dir, &l, &cmd.len)

	if err != nil {
		return 0, CMD{}
	}

	cmd.dir = byte(cmd.len & 3)
	cmd.len = cmd.len >> 4
	return 1, cmd
}

func setup(path string, read func(*os.File) (int, CMD)) []CMD {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	commands := make([]CMD, 0, 1000)

	for {
		ret, cmd := read(file)
		if ret != 1 {
			break
		}
		commands = append(commands, cmd)

	}

	return commands
}

func toInt(a bool) int {
	if a {
		return 1
	}
	return 0
}

func solve(commands []CMD) int64 {
	curDir := commands[0].dir
	var nextDir byte

	curLen := commands[0].len
	nextLen := 0

	y := 0
	area := int64(0)
	clockwise := true

	for i := 1; i < len(commands); i++ {
		nextDir = commands[i].dir
		nextLen = commands[i].len

		nextClockwise := nextDir == ((curDir + 1) & 0b11)
		curLen += toInt(nextClockwise) + toInt(clockwise) - 1
		if curDir&1 != 0 {
			// up or down
			if curDir < 2 {
				//down
				y += curLen
			} else {
				//up
				y -= curLen
			}
		} else {
			// left or right
			if curDir > 1 {
				// left
				area += int64(curLen) * int64(y)
			} else {
				//right
				area += int64(-curLen) * int64(y)
			}
		}
		clockwise = nextClockwise
		curDir = nextDir
		curLen = nextLen
	}

	return area
}
