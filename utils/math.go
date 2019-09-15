package utils

import "math"

func Round(input float32) float32 {
	return float32(math.Round(float64(input)))
}
