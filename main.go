package main

import (
	"fmt"
	"math"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type FinacialMonthDetail struct {
	Year                   int
	Month                  int
	Incomess               []FinacialItem
	Assets                 []FinacialItem
	Expenses               []FinacialItem
	Liabilities            []FinacialItem
	TotalIncomes           float64
	TotalExpenses          float64
	PreAllocationCashFlow  float64
	PostAllocationCashFlow float64

	TotalAssets      float64
	TotalLiabilities float64
	NetWorth         float64
}

type FinancialPeriod struct {
	StartYear  int
	StartMonth int
	EndYear    int
	EndMonth   int
	Amount     float64
	GrowthRate float64 // monthly growth
}

type ProjectionPlans struct {
	Plans []FinacialItem
}

type FinancialModel struct {
	Assets             []FinacialItem
	Expenses           []FinacialItem
	Incomes            []FinacialItem
	Liability          []FinacialItem
	CashflowPriority   []Priority
	TransactionHistory []Transaction
}

type Transaction struct {
	From   string
	To     string
	Amount string
}

type Priority struct {
	ItemID    int
	MaxInsert int
}

type FinacialItem struct {
	ID         int
	Name       string
	StartYear  int
	StartMonth int
	EndYear    int
	EndMonth   int
	Amount     float64
	GrowthRate float64 // monthly growth
}

// Returns the updated FinacialItem List after cashflow has been distrubuted to accounts per Cashflow Priority Order
func (sp *FinancialModel) priorityEngine(items []FinacialItem, cashflow float64, total float64) (float64, []FinacialItem, float64) {
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

func growthLogger(growth float64) {

}

// Returns the total amount for the array of items and the updated item list.
func (j *FinancialModel) totalForMonth(year, month int, items []FinacialItem) (float64, []FinacialItem) {
	total := 0.0
	var new_items []FinacialItem
	for _, p := range items {
		if periodApplies(year, month, p) {
			// Apply growth for this period (plus 1 so that we get the value at the end of the month)

			// We would need to calculate the change in months as its going up by 1
			// monthsElapsed := (year-p.StartYear)*12 + (month - p.StartMonth) + 1
			fmt.Println("months elapsed", 1)
			growthAmount := p.Amount * math.Pow(1+p.GrowthRate, float64(1))

			currentAmount := growthAmount

			fmt.Println("current amount", currentAmount)

			item := p
			item.Amount = currentAmount
			new_items = append(new_items, item)
			total += currentAmount

		}
	}
	return total, new_items
}

// Finds whether a item is currently active or not -> returns bool
func periodApplies(year, month int, p FinacialItem) bool {
	start := p.StartYear*12 + p.StartMonth
	end := p.EndYear*12 + p.EndMonth
	current := year*12 + month
	return current >= start && current <= end
}

func printProjectionTable(projection_plan []FinacialMonthDetail) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Year", "Month", "Total Incomes", "Total Expenses", "PRE ALLOC Cash Flow", "Total Assets", "Total Liabilities", "Net Worth", "POST ALLOC CASHFLOW"})

	for _, m := range projection_plan {
		t.AppendRow(table.Row{
			m.Year,
			fmt.Sprintf("%02d", m.Month),
			fmt.Sprintf("£%.2f", m.TotalIncomes),
			fmt.Sprintf("£%.2f", m.TotalExpenses),
			fmt.Sprintf("£%.2f", m.PreAllocationCashFlow),
			fmt.Sprintf("£%.2f", m.TotalAssets),
			fmt.Sprintf("£%.2f", m.TotalLiabilities),
			fmt.Sprintf("£%.2f", m.NetWorth),
			fmt.Sprintf("%.2f", m.PostAllocationCashFlow),
		})
	}

	t.Render()

}

func (p *FinancialModel) generateProjection(startYear, startMonth, endYear, endMonth int) []FinacialMonthDetail {

	year, month := startYear, startMonth

	var report []FinacialMonthDetail

	current := p

	for {
		fmt.Println("current assets", current.Assets)
		totalIncomes, newIncomess := p.totalForMonth(year, month, current.Incomes)
		totalExpense, newExpenses := p.totalForMonth(year, month, current.Expenses)

		preAllocCashflow := totalIncomes - totalExpense

		// where the excess cash will go

		// we will calculate the excess cash after the growth so after get the figures back its both easier and more realistic.
		//
		totalAssetsG, newAssetsGrowth := p.totalForMonth(year, month, current.Assets)
		fmt.Println("Total Assets", totalAssetsG)

		// May be better off inside of the totalformonth func to be selected during the loop for performance gain
		totalAssets, newAssets, postAllocationCashFlow := p.priorityEngine(newAssetsGrowth, preAllocCashflow, totalAssetsG)

		totalLiabilites, newLiabilities := p.totalForMonth(year, month, current.Liability)

		networth := totalAssets - totalLiabilites

		report = append(report, FinacialMonthDetail{Year: year, Month: month, Incomess: newIncomess, TotalIncomes: totalIncomes, Expenses: newExpenses, Assets: newAssets, Liabilities: newLiabilities, NetWorth: networth, PreAllocationCashFlow: preAllocCashflow, PostAllocationCashFlow: postAllocationCashFlow, TotalExpenses: totalExpense, TotalAssets: totalAssets, TotalLiabilities: totalLiabilites})

		if year == endYear && month == endMonth {
			break
		}
		month++

		if month > 12 {
			month = 0
			year++
		}
		current.Assets = newAssets
		current.Liability = newLiabilities

	}

	return report

}

func yearRatetoMonthlyRate(yr float64) float64 {
	return math.Pow((1+yr), 1.0/12.0) - 1
}

func main() {

	fmt.Println("Monthly Rate: ", yearRatetoMonthlyRate(0.005))
	assets := []FinacialItem{
		{ID: 1, Name: "Stocks", StartYear: 2025, StartMonth: 1, EndYear: 2060, EndMonth: 1, Amount: 40000, GrowthRate: yearRatetoMonthlyRate(0.05)}, // 2% per month
		// {ID: 2, Name: "Savings", StartYear: 2050, StartMonth: 1, EndYear: 2080, EndMonth: 12, Amount: 30000, GrowthRate: yearRatetoMonthlyRate(0.05)}, // 2% per month
	}

	Incomes := []FinacialItem{
		{Name: "Salary", StartYear: 2025, StartMonth: 1, EndYear: 2080, EndMonth: 12, Amount: 1000, GrowthRate: 0},
		{Name: "StatePension", StartYear: 2064, StartMonth: 5, EndYear: 2080, EndMonth: 12, Amount: 1200, GrowthRate: 0}, // 2% per month
		// 2% per month
	}

	expenses := []FinacialItem{
		{Name: "Rent", StartYear: 2025, StartMonth: 1, EndYear: 2025, EndMonth: 12, Amount: 500, GrowthRate: 0}, // 2% per month
	}
	Liabilities := []FinacialItem{
		{Name: "House", StartYear: 2025, StartMonth: 1, EndYear: 2025, EndMonth: 12, Amount: 500, GrowthRate: 0}, // 2% per month
	}

	// amount, updatedFinance := totalForMonth(2064, 12, assets)
	pro := Priority{ItemID: 1, MaxInsert: 0}
	model := FinancialModel{Assets: assets, Liability: Liabilities, Expenses: expenses, Incomes: Incomes, CashflowPriority: []Priority{pro}}

	printProjectionTable(model.generateProjection(2025, 1, 2030, 12))

	// fmt.Println("Report", model.generateProjection(2025, 1, 2025, 12))

}
