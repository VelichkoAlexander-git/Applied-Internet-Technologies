package math

// Найти среднее в массиве чисел.
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Найти минимальное в массиве чисел.
func Min(xs []float64) (float64, bool) {
	if len(xs) == 0 {
		return 0.0, false
	}
	total := xs[0]
	for _, x := range xs {
		if total > x {
			total = x
		}
	}
	return total, true
}

// Найти максимальное в массиве чисел.
func Max(xs []float64) (float64, bool) {
	if len(xs) == 0 {
		return 0.0, false
	}
	total := xs[0]
	for _, x := range xs {
		if total < x {
			total = x
		}
	}
	return total, true
}
