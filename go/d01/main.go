package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

const INPUT = "/home/jonas/src/aoc23/input/d01/input"

func main() {
	file, err := os.Open(INPUT)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	p1Sum := 0
	p2Sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		p1n1, _ := p1GetFirstNum(line)
		p1n2, _ := p1GetReverseNum([]rune(line))
		p2n1 := p2([]byte(line))
		p2n2 := p2Rev([]byte(line))
		p1Sum += p1n1*10 + p1n2
		p2Sum += p2n1*10 + p2n2
	}
	log.Printf("Part 1: %d", p1Sum)
	log.Printf("Part 2: %d", p2Sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func p1GetReverseNum(line []rune) (int, error) {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(line[i]) {
			return strconv.Atoi(string(line[i]))
		}
	}
	return 0, nil
}

func p1GetFirstNum(line string) (int, error) {
	for _, char := range line {
		if unicode.IsDigit(char) {
			return strconv.Atoi(string(char))
		}
	}
	return 0, nil
}

var words = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var wordsRev = [9]string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

func stateMachine(line []byte, state uint64, i int, words [9]string) (int, uint64) {
	if unicode.IsDigit(rune(line[i])) {
		res, _ := strconv.Atoi(string(line[i]))
		return res, state
	}

	for j := 0; j < 9; j++ {
		if (state>>(j*4))&0b1 == 1 {
			count := ((state >> ((j * 4) + 1)) & 0b111) + 1
			// fmt.Printf("%d: Count %d", j, int(count))

			if len(words[j])-1 == int(count) && line[i] == words[j][count] {
				return j + 1, state
			}

			if len(words[j]) <= int(count) || line[i] != words[j][count] {
				state &= ^(0b1111 << (j * 4))
			} else {
				state += 0b10 << (j * 4)
			}
		}
	}

	return -1, state
}

func p2(line []byte) int {
	state := uint64(0)
	res := -1
	for i := 0; i < len(line); i++ {
		res, state = stateMachine(line, state, i, words)
		if res != -1 {
			return res
		}

		switch line[i] {
		case 'o':
			state |= uint64(0b1)
		case 't':
			state |= uint64(0b10001) << 4
		case 'f':
			state |= uint64(0b10001) << (3 * 4)
		case 's':
			state |= uint64(0b10001) << (5 * 4)
		case 'e':
			state |= uint64(0b10001) << (7 * 4)
		case 'n':
			state |= uint64(0b1) << (8 * 4)
		}
	}
	return 0
}

func p2Rev(line []byte) int {
	state := uint64(0)
	res := -1
	for i := len(line) - 1; i >= 0; i-- {
		res, state = stateMachine(line, state, i, wordsRev)
		if res != -1 {
			return res
		}

		switch line[i] {
		case 'e':
			state |= uint64(0b1)
			state |= uint64(0b1) << (2 * 4)
			state |= uint64(0b1) << (4 * 4)
			state |= uint64(0b1) << (8 * 4)
		case 'o':
			state |= uint64(0b1) << 4
		case 'r':
			state |= uint64(0b1) << (3 * 4)
		case 'x':
			state |= uint64(0b1) << (5 * 4)
		case 'n':
			state |= uint64(0b1) << (6 * 4)
		case 't':
			state |= uint64(0b1) << (7 * 4)
		}
	}
	return 0
}
