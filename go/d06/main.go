package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jonsch318/aoc23/go/d06/add"
)

type Race struct {
	time     float64
	distance float64
}

func main() {
	races := setup("/home/jonas/src/aoc23/input/d06/test")
	log.Printf("SETUP: %v", races)
	res := solveP1(races)
	log.Printf("P1: %v", res)
}

func solveP1(races []Race) uint64 {
	return add.Add(1, 2)
}

func toDouble(a string) float64 {
	res, _ := strconv.ParseFloat(a, 64)
	return res
}

func setup(path string) []Race {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	text := string(data)

	lines := strings.Split(text, "\n")

	if len(lines) != 2 {
		return nil
	}

	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	races := make([]Race, 0, len(times))

	for i := range times {
		races = append(races, Race{
			time:     toDouble(times[i]),
			distance: toDouble(distances[i]),
		})
	}

	return races
}
