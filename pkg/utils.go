package pkg

import "math"

func YearRatetoMonthlyRate(yr float64) float64 {
	return math.Pow((1+yr), 1.0/12.0) - 1
}
