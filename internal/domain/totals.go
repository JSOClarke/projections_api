package domain

import (
	"math"
)

// Returns the Total Value and
func TotalForMonth(year, month int, items []FinacialItem) (float64, []FinacialItem) {
	total := 0.0
	var updated_items []FinacialItem
	for _, item := range items {
		if IsPeriodApplied(year, month, item) {
			// Apply growth for this period (plus 1 so that we get the value at the end of the month)

			// We would need to calculate the change in months as its going up by 1
			// monthsElapsed := (year-p.StartYear)*12 + (month - p.StartMonth) + 1
			// fmt.Println("months elapsed", 1)
			growthAmount := item.Amount * math.Pow(1+item.GrowthRate, float64(1))

			currentAmount := growthAmount

			// fmt.Println("current amount", currentAmount)

			updated_item := item
			updated_item.Amount = currentAmount
			updated_items = append(updated_items, item)
			total += currentAmount

		}
	}
	return total, updated_items
}
