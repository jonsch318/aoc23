package solve

import (
	"math"

	"github.com/chewxy/math32"
)

//go:noescape
func SolveAVX(times [4]float32, distances [4]float32) [4]float32

//go:noescape
func SolveAVX2(times [4]float64, distances [4]float64) [4]float64

func SolveScalar32(time, distance float32) float32 {
	timeHalf := time * 0.5
	x1 := timeHalf - sign32(timeHalf)*math32.Sqrt(math32.Pow(timeHalf, 2)-distance)
	x2 := math32.Floor(distance / x1)
	return x2 - math32.Ceil(x1) + 1
}

func sign32(f float32) float32 {
	if math32.Signbit(f) {
		return -1
	}
	return 1
}

func SolveScalar(time, distance float64) float64 {
	timeHalf := time * 0.5
	x1 := timeHalf - sign(timeHalf)*math.Sqrt(math.Pow(timeHalf, 2)-distance)
	x2 := math.Floor(distance / x1)
	return x2 - math.Ceil(x1) + 1
}

func sign(f float64) float64 {
	if math.Signbit(f) {
		return -1
	}
	return 1
}
