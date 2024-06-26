package stats

import "math"

// standard deviation is simply the square root of variance
func Stdev(values []float64) float64 {
	variance := Variance(values)
	stdev := math.Sqrt(variance)

	return float64(stdev)
}
