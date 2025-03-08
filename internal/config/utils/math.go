package utils

func Sum(values []float64) float64 {
	var sum float64

	for _, value := range values {
		sum += value
	}

	return sum
}

func Prod(value1 float64, value2 int64) float64 {
	return value1 * float64(value2)
}
