package presentation

import (
	"projection_test/internal/domain"
	"projection_test/pkg"
	"testing"
)

// Testing Not completed
func TestTableView(t *testing.T) {
	assets := []domain.FinacialItem{
		{ID: 1, Name: "Stocks", StartYear: 2025, StartMonth: 1, EndYear: 2080, EndMonth: 12, Amount: 40000, GrowthRate: pkg.YearRatetoMonthlyRate(0.05)},  // 2% per month
		{ID: 2, Name: "Pension", StartYear: 2025, StartMonth: 1, EndYear: 2080, EndMonth: 12, Amount: 10000, GrowthRate: pkg.YearRatetoMonthlyRate(0.05)}, // 2% per month
	}

	Incomes := []domain.FinacialItem{
		{Name: "Salary", StartYear: 2025, StartMonth: 1, EndYear: 2080, EndMonth: 12, Amount: 1000, GrowthRate: 0},
		// {Name: "StatePension", StartYear: 2064, StartMonth: 5, EndYear: 2080, EndMonth: 12, Amount: 1200, GrowthRate: 0}, // 2% per month
		// 2% per month
	}

	expenses := []domain.FinacialItem{
		{Name: "Rent", StartYear: 2025, StartMonth: 1, EndYear: 2080, EndMonth: 12, Amount: 500, GrowthRate: 0}, // 2% per month
	}
	Liabilities := []domain.FinacialItem{
		// {Name: "House", StartYear: 2025, StartMonth: 1, EndYear: 2025, EndMonth: 12, Amount: 500, GrowthRate: 0}, // 2% per month
	}
	data := []domain.FinacialMonthDetail{{Year: 2025, Month: 1, Incomess: Incomes, Assets: assets, Expenses: expenses, Liabilities: Liabilities, TotalIncomes: 1000, TotalExpenses: 1000, TotalAssets: 1000, TotalLiabilities: 1000}}
	PrintProjectionsTable(data)

}
