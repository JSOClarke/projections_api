package domain

// Finds whether a item is currently active or not -> returns bool
func IsPeriodApplied(year, month int, p FinacialItem) bool {
	start := p.StartYear*12 + p.StartMonth
	end := p.EndYear*12 + p.EndMonth
	current := year*12 + month
	return current >= start && current <= end
}
