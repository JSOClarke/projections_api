package usecase

import "projection_test/internal/domain"

func GenerateProjection(startYear, startMonth, endYear, endMonth int, p domain.FinancialModel) []domain.FinacialMonthDetail {

	year, month := startYear, startMonth

	var report []domain.FinacialMonthDetail

	current := p

	for {
		// we need to split up into phases from totalForMonth to total and then growth for this
		totalIncomes, newIncomess := domain.TotalForMonth(year, month, current.Incomes)
		totalExpense, newExpenses := domain.TotalForMonth(year, month, current.Expenses)

		preAllocCashflow := totalIncomes - totalExpense

		// Cashflow Calculation

		totalAssetsG, newAssetsGrowth := domain.TotalForMonth(year, month, current.Assets)

		// Asset Calculation

		totalAssets, newAssets, postAllocationCashFlow := domain.ApplyPriority(p.Assets, newAssetsGrowth, preAllocCashflow, totalAssetsG)

		//

		totalLiabilites, newLiabilities := p.totalForMonth(year, month, current.Liability)

		networth := totalAssets - totalLiabilites

		report = append(report, domain.FinacialMonthDetail{Year: year, Month: month, Incomess: newIncomess, TotalIncomes: totalIncomes, Expenses: newExpenses, Assets: newAssets, Liabilities: newLiabilities, NetWorth: networth, PreAllocationCashFlow: preAllocCashflow, PostAllocationCashFlow: postAllocationCashFlow, TotalExpenses: totalExpense, TotalAssets: totalAssets, TotalLiabilities: totalLiabilites})

		if year == endYear && month == endMonth {
			break
		}
		month++

		if month > 12 {
			month = 1
			year++
		}
		current.Assets = newAssets
		current.Liability = newLiabilities

	}
	p.Results = report
	return report

}
