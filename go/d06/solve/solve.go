package solve

import (
	"math"
)

//go:noescape
func SolveAVX(times [4]float32, distances [4]float32) [4]float32

//go:noescape
func SolveAVX2(times [4]float64, distances [4]float64) [4]float64

func SolveScalar(time, distance float64) float64 {
	timeHalf := time * 0.5
	x1 := math.Ceil(timeHalf - math.Sqrt(math.Pow(timeHalf, 2)-distance))
	x2 := math.Floor(timeHalf + math.Sqrt(math.Pow(timeHalf, 2)-distance))
	return x2 - x1 + 1
}
