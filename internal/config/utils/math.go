package config

func Sum(values []float64) float64 {
	var sum float64

	for _, value := range values {
		sum += value
	}

	return sum
}

func Prod(value1, value2 float64) float64 {
	return value1 * value2
}
