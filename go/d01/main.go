package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

const INPUT = "input/d01/input"

func main() {
	file, err := os.Open(INPUT)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sum := 0
	for _, line := range lines {
		n1, _ := p1GetFirstNum(line)
		n2, _ := p1GetReverseNum([]rune(line))
		sum += n1*10 + n2
	}
	log.Printf("Part 1: %d", sum)

	sum = 0
	for _, line := range lines {
		n1, _ := p2([]byte(line))
		n2, _ := p2Rev([]byte(line))
		sum += n1*10 + n2
	}

	log.Printf("Part 2: %d", sum)

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

func p2(line []byte) (int, error) {
	state := uint64(0)

	for i := 0; i < len(line); i++ {

		if unicode.IsDigit(rune(line[i])) {
			return strconv.Atoi(string(line[i]))
		}

		for j := 0; j < 9; j++ {
			if (state>>(j*4))&0b1 == 1 {
				count := ((state >> ((j * 4) + 1)) & 0b111) + 1
				// fmt.Printf("%d: Count %d", j, int(count))

				if len(words[j])-1 == int(count) && line[i] == words[j][count] {
					return j + 1, nil
				}

				if len(words[j]) <= int(count) || line[i] != words[j][count] {
					state &= ^(0b1111 << (j * 4))
				} else {
					state += 0b10 << (j * 4)
				}
			}
		}

		switch line[i] {
		case 'o':
			state |= uint64(0b1)
		case 't':
			state |= uint64(0b10001) << 4
		case 'f':
			state |= uint64(0b10001) << 3 * 4
		case 's':
			state |= uint64(0b10001) << 5 * 4
		case 'e':
			state |= uint64(0b10001) << 7 * 4
		case 'n':
			state |= uint64(0b1) << 8 * 4
		}
	}
	return 0, nil
}

var wordsRev = [9]string{"eno", "owt", "eerht", "rouf", "evif", "xis", "neves", "thgie", "enin"}

func p2Rev(line []byte) (int, error) {
	state := uint64(0)

	for i := len(line) - 1; i >= 0; i-- {

		if unicode.IsDigit(rune(line[i])) {
			return strconv.Atoi(string(line[i]))
		}

		for j := 0; j < 9; j++ {
			if (state>>(j*4))&0b1 == 1 {
				count := ((state >> ((j * 4) + 1)) & 0b111) + 1
				// fmt.Printf("%d: Count %d", j, int(count))

				if len(wordsRev[j])-1 == int(count) && line[i] == wordsRev[j][count] {
					return j + 1, nil
				}

				if len(wordsRev[j]) <= int(count) || line[i] != wordsRev[j][count] {
					state &= ^(0b1111 << (j * 4))
				} else {
					state += 0b10 << (j * 4)
				}
			}
		}

		switch line[i] {
		case 'o':
			state |= uint64(0b1)
		case 't':
			state |= uint64(0b10001) << 4
		case 'f':
			state |= uint64(0b10001) << 3 * 4
		case 's':
			state |= uint64(0b10001) << 5 * 4
		case 'e':
			state |= uint64(0b10001) << 7 * 4
		case 'n':
			state |= uint64(0b1) << 8 * 4
		}
	}
	return 0, nil
}
