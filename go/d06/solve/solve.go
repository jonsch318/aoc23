package solve

import (
	"math"
	"unsafe"
)

//go:noescape
func SolveAVX(times [4]float32, distances [4]float32) [4]float32

//go:noescape
func SolveAVX2(times, distances unsafe.Pointer) float32

func SolveScalar(time, distance float32) float32 {
	pT, pD := float64(time), float64(distance)
	timeHalf := pT * 0.5
	x1 := math.Ceil(timeHalf - math.Sqrt(math.Pow(timeHalf, 2)-pD))
	x2 := math.Floor(timeHalf + math.Sqrt(math.Pow(timeHalf, 2)-pD))
	return float32(x2 - x1 + 1)
}
