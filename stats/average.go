package stats

func Average(values []float64) float64 {
	sum := 0.0
	for _, num := range values {
		sum += num
	}
	average := sum / float64(len(values))
	return average
}
