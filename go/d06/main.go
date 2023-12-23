package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	solve "github.com/jonsch318/aoc23/go/d06/solve"
)

var impl = Default

const (
	Default int = iota
	AVX
	AVX2
)

func switchImpl(nImpl int) {
	impl = nImpl
}

func main() {
	races, length := setup32("/home/jonas/src/aoc23/input/d06/input")
	log.Printf("SETUP: %v", races)
	res := solve32(races, length)
	log.Printf("P1: %v", res)

	races64, length64 := setup64("/home/jonas/src/aoc23/input/d06/input")
	res = solveP2(races64, length64)
	log.Printf("P2: %v", res)

}

func solve32(timesDistances []float32, length int) int64 {
	prod := int64(1)
	switch impl {
	case AVX:
		for i := 0; i < len(timesDistances); i += 8 {
			ptr := solve.SolveAVX([4]float32(timesDistances[i:i+4]), [4]float32(timesDistances[i+4:i+8]))
			prod *= int64(ptr[0])
			prod *= int64(ptr[1])
			prod *= int64(ptr[2])
			prod *= int64(ptr[3])
		}
	case AVX2:
	case Default:
		for i := 0; i < len(timesDistances); i += 2 {
			prod *= int64(solve.SolveScalar(float64(timesDistances[i]), float64(timesDistances[i+1])))
		}
	}
	return prod
}

func solve64(timesDistances []float64, length int) int64 {
	prod := int64(1)
	switch impl {
	case AVX:
	case AVX2:
		for i := 0; i < len(timesDistances); i += 8 {
			ptr := solve.SolveAVX2([4]float64(timesDistances[i:i+4]), [4]float64(timesDistances[i+4:i+8]))
			prod *= int64(ptr[0])
			prod *= int64(ptr[1])
			prod *= int64(ptr[2])
			prod *= int64(ptr[3])
		}
	case Default:
		for i := 0; i < len(timesDistances); i += 2 {
			prod *= int64(solve.SolveScalar(timesDistances[i], timesDistances[i+1]))
		}
	}
	return prod
}

func solveP2(timesDistances []float64, length int) int64 {
	var time float64
	var dist float64
	powerTime := float64(1.)
	powerDist := float64(1.)
	for i := 0; i < length; i++ {
		magnitudeTime := math.Pow(10, math.Floor(math.Log10(float64(timesDistances[i*2])))+1)
		time = time*magnitudeTime + timesDistances[i*2]
		powerTime *= magnitudeTime
		magnitudeDist := math.Pow(10, math.Floor(math.Log10(float64(timesDistances[(i*2)+1])))+1)
		dist = dist*magnitudeDist + timesDistances[(i*2)+1]
		powerDist *= magnitudeDist
	}

	return int64(solve.SolveScalar(time, dist))
}

func toDouble(a string) float64 {
	res, _ := strconv.ParseFloat(a, 64)
	return float64(res)
}

func toFloat(a string) float32 {
	res, _ := strconv.ParseFloat(a, 32)
	return float32(res)
}

func setup32(path string) ([]float32, int) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("not good %v", err)
		return nil, 0
	}
	text := string(data)

	lines := strings.Split(text, "\n")
	lines = lines[0:2]

	timeStrings := strings.Fields(lines[0])[1:]
	distanceStrings := strings.Fields(lines[1])[1:]

	// cache efficient array. If AVX2 => [8 times...][8 distances...][8 times...]... else
	// Default => [time][distance][time]...

	var vectorLength = 1
	length := len(timeStrings)
	switch impl {
	case AVX:
		vectorLength = 128 / (4 * 8)
	case AVX2:
		vectorLength = 256 / (4 * 8)
	}

	if length%vectorLength != 0 {
		length += vectorLength - (len(timeStrings) % vectorLength)
	}

	timesDistances := make([]float32, length*2)
	for i := 0; i < len(timeStrings); i += vectorLength {
		for j := 0; j < vectorLength; j++ {
			x := (i * 2) + j
			if len(timeStrings) <= i+j {
				timesDistances[x] = 0
				timesDistances[x+vectorLength] = 0
			} else {
				timesDistances[x] = toFloat(timeStrings[i+j])
				timesDistances[x+vectorLength] = toFloat(distanceStrings[i+j])
			}
		}
	}
	return timesDistances, len(timeStrings)
}

func setup64(path string) ([]float64, int) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("not good %v", err)
		return nil, 0
	}
	text := string(data)

	lines := strings.Split(text, "\n")
	lines = lines[0:2]

	timeStrings := strings.Fields(lines[0])[1:]
	distanceStrings := strings.Fields(lines[1])[1:]

	// cache efficient array. If AVX2 => [8 times...][8 distances...][8 times...]... else
	// Default => [time][distance][time]...

	var vectorLength = 1
	length := len(timeStrings)
	switch impl {
	case AVX:
		vectorLength = 128 / (4 * 8)
	case AVX2:
		vectorLength = 256 / (8 * 8)
	}

	if length%vectorLength != 0 {
		length += vectorLength - (len(timeStrings) % vectorLength)
	}

	timesDistances := make([]float64, length*2)
	for i := 0; i < len(timeStrings); i += vectorLength {
		for j := 0; j < vectorLength; j++ {
			x := (i * 2) + j
			if len(timeStrings) <= i+j {
				timesDistances[x] = 0
				timesDistances[x+vectorLength] = 0
			} else {
				timesDistances[x] = toDouble(timeStrings[i+j])
				timesDistances[x+vectorLength] = toDouble(distanceStrings[i+j])
			}
		}
	}
	return timesDistances, len(timeStrings)
}
