package domain

import "math"

// Returns the updated FinacialItem List after cashflow has been distrubuted to accounts per Cashflow Priority Order
func ApplyPriority(items []FinacialItem, cashflow float64, total float64, sp FinancialModel) (float64, []FinacialItem, float64) {

	// No provided cashflow_prority
	if len(sp.CashflowPriority) == 0 {
		// what happens if there are no asset accounts where does the cashflow go. Maybe we create a safety account to hold the excess
		items[0].Amount = items[0].Amount + cashflow
		total += cashflow
		cashflow = 0
		return total, items, cashflow
	}
	for _, i := range sp.CashflowPriority {
		for j := range items {
			if items[j].ID == i.ItemID {
				excess := math.Min(cashflow, float64(i.MaxInsert))
				items[j].Amount = float64(items[j].Amount) + excess
				cashflow = cashflow - excess
				total = total + excess
			}
		}
	}

	if cashflow != 0 {
		items[0].Amount = items[0].Amount + cashflow
		total += cashflow
		cashflow = 0
	}
	return total, items, cashflow

}
