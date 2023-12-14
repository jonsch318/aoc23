package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

const INPUT = "/home/jonas/src/aoc23/input/d01/input"

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	mainArgs(true)
}

func mainArgs(print bool) {
	start := time.Now()
	lines := setup()

	var p1Sum, p2Sum int
	//for i := 0; i < 100000; i++ {
	p1Sum, p2Sum = run(lines)
	//}

	elapsed := time.Since(start)
	if print {
		log.Printf("Part 1: %d", p1Sum)
		log.Printf("Part 2: %d", p2Sum)
		log.Printf("%v ms | %v ns", elapsed.Milliseconds(), elapsed.Nanoseconds())
	}
}

func setup() [][]byte {
	file, err := os.Open(INPUT)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]byte, 0)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func run(lines [][]byte) (p1Sum int, p2Sum int) {
	p1Sum = runP1(lines)
	p2Sum = runP2(lines)
	return
}

func runP1(lines [][]byte) (sum int) {
	var n1, n2 int
	for _, line := range lines {
		n1 = p1GetFirstNum(line)
		n2 = p1GetReverseNum(line)
		sum += n1*10 + n2
	}
	return
}

func runP2(lines [][]byte) (sum int) {
	var n1, n2 int
	for _, line := range lines {
		n1 = p2(line)
		n2 = p2Rev(line)
		sum += n1*10 + n2
	}
	return
}

func p1GetReverseNum(line []byte) int {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= 48 && line[i] <= 57 {
			return int(line[i] - 48)
		}
	}
	return 0
}

func p1GetFirstNum(line []byte) int {
	for i := range line {
		if line[i] >= 48 && line[i] <= 57 {
			return int(line[i] - 48)
		}
	}
	return 0
}

var words = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var wordsRev = [9]string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

func p2(line []byte) int {
	state := uint64(0)

	for i := 0; i < len(line); i++ {
		if line[i] >= 48 && line[i] <= 57 {
			return int(line[i] - 48)
		}

		for j := 0; j < 9; j++ {
			shift := j << 2 //j * 4
			if (state & (1 << shift)) != 0 {
				count := ((state >> (shift + 1)) & 0b111) + 1

				if len(words[j])-1 == int(count) && line[i] == words[j][count] {
					return j + 1
				}

				if len(words[j]) <= int(count) || line[i] != words[j][count] {
					state &= ^(0b1111 << shift)
				} else {
					state += 0b10 << shift
				}
			}
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
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= 48 && line[i] <= 57 {
			return int(line[i] - 48)
		}

		for j := 0; j < 9; j++ {
			shift := j << 2 //j * 4
			if (state & (1 << shift)) != 0 {
				count := ((state >> (shift + 1)) & 0b111) + 1

				if len(wordsRev[j])-1 == int(count) && line[i] == wordsRev[j][count] {
					return j + 1
				}

				if len(wordsRev[j]) <= int(count) || line[i] != wordsRev[j][count] {
					state &= ^(0b1111 << shift)
				} else {
					state += 0b10 << shift
				}
			}
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
