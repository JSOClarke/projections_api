package presentation

import (
	"fmt"
	"os"
	"projection_test/internal/domain"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintProjectionsTable(projection_plan []domain.FinacialMonthDetail) {
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
