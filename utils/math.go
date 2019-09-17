package utils

import "math"

func Round(input float32) float32 {
	return float32(math.Round(float64(input)))
}

func ToRadians(degrees int) float64 {
	return float64(degrees) * (math.Pi / float64(180))
}
