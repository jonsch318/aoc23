package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jonsch318/aoc23/go/d12/test"
)

const PRETTY_STRING_LENGTH = 20
const INPUT = "input/d12/input"

func parseCounts(numberStrings []string) []int {
	numbers := make([]int, len(numberStrings))
	for i, numberString := range numberStrings {
		numbers[i], _ = strconv.Atoi(numberString)
	}
	return numbers
}

func countPossibilities(s []byte, counts []int) int {
	// solution using NFA.
	// we count all
	possible := 0
	currentStates := map[[4]int]int{{0, 0, 0, 0}: 1} // Tuple sIndex, patternIndex, nextExpected, runCount: possibles in the state
	nextStates := map[[4]int]int{}                   //used for allocation reasons

	for len(currentStates) > 0 {
		//we go over all possible states of the NFA
		for state, num := range currentStates {
			sIndex, countsIndex, nextExpected, runCount := state[0], state[1], state[2], state[3]

			if len(s) <= sIndex {
				// we are finished

				if countsIndex <= len(counts) {
					possible += num
				}
				continue
			}

			char := s[sIndex]
			//log.Printf("(%c, %d, %d, %d): %d", rune(char), counts[countsIndex], nextExpected, runCount, num)
			if char != '.' && nextExpected == 0 && countsIndex < len(counts) {

				if char == '?' && runCount == 0 {
					// we are not in a run: '?' has two options
					nextStates[[4]int{sIndex + 1, countsIndex, nextExpected, runCount}] += num
				}
				//the run can continue
				runCount++
				if runCount >= counts[countsIndex] {
					countsIndex += 1
					//reset the run and expect '.' next
					nextExpected = 1
					runCount = 0
				}
				nextStates[[4]int{sIndex + 1, countsIndex, nextExpected, runCount}] += num
			} else if char != '#' && runCount == 0 {
				nextExpected = 0
				nextStates[[4]int{sIndex + 1, countsIndex, nextExpected, runCount}] += num
			}
		}

		currentStates, nextStates = nextStates, currentStates
		for k := range nextStates {
			delete(nextStates, k)
		}

	}

	return possible
}

func padStringForPrettyLog(s string) string {
	ret := strings.Clone(s)
	for i := len(s); i <= PRETTY_STRING_LENGTH; i++ {
		ret += " "
	}
	return ret
}

func main() {
	file, err := os.Open(INPUT)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0
	sum := 0
	for scanner.Scan() {
		i++
		str := scanner.Text()
		if i != 1 {
			continue
		}
		before, after, _ := strings.Cut(str, " ")
		counts := parseCounts(strings.Split(after, ","))
		possibilities := countPossibilities([]byte(before), counts)
		sum += possibilities

		res := test.Counter{Cache: make(map[test.Entry]int)}.
			CountArrangements(before, counts, -1)

		log.Printf("Processed Line %s %v: %d == %d (%v)", padStringForPrettyLog(before), counts, possibilities, res, res == possibilities)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	log.Printf("Counted Posibilites Part 1: %d", sum)
}
