package domain

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
	Results            []FinacialMonthDetail
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
