package domain

import (
	"fmt"
	"testing"
)

func TestTotalForMonth(t *testing.T) {

	year := 2025
	month := 12
	items := []FinacialItem{{ID: 1, Name: "Test", StartYear: 2025, StartMonth: 10, EndYear: 2026, EndMonth: 5, Amount: 1000, GrowthRate: 0.5}}
	total, updatedItems := TotalForMonth(year, month, items)
	fmt.Println("total:", total)
	fmt.Println("updatedItems", updatedItems)
}
