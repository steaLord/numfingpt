package models

type FinancialData struct {
	ChartDataFinancials struct {
		CompanyName string `json:"company_name"`
		Ticker      string `json:"ticker"`
	} `json:"chartDataFinancials"`
	BalanceSheet struct {
		AverageYearlyAssetGrowth             string `json:"average_yearly_asset_growth"`
		QuarterOverQuarterAssetChangePercent string `json:"quarter_over_quarter_asset_change_percent"`
		AssetsGrowthRate                     string `json:"assets_growth_rate"`
		TotalCurrentAssets                   string `json:"total_current_assets"`
		CashAndShortTermInvestments          string `json:"cash_and_short_term_investments"`
		CashAndShortTermInvestmentsGrowth    string `json:"cash_and_short_term_investments_growth"`
	} `json:"balanceSheet"`
}
