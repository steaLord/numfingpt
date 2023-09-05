package functions

import (
	"encoding/json"
	"fmt"
	"numfingpt/internal/models"
)

func ConvertData(data []byte) (string, error) {

	var fin_data models.FinancialData
	err := json.Unmarshal(data, &fin_data)
	if err != nil {
		return "", fmt.Errorf("Error decoding JSON: %w", err)
	}

	return fmt.Sprintf("Diving into the fiscal trajectory of %s, we observe an average asset growth. This rate, interestingly, stands at %s, reflecting both the company's highs and lows. When compared quarter-over-quarter, this figure adjusts to %s. A look back at the past year reveals a total asset change of %s. In the realm of current assets, %s clocks in at %s in the reporting currency. A significant portion of these assets, precisely %s, is held in cash and short-term investments. This segment shows a change of %s when juxtaposed with last year's data.",
		fin_data.ChartDataFinancials.CompanyName,
		fin_data.BalanceSheet.AverageYearlyAssetGrowth,
		fin_data.BalanceSheet.QuarterOverQuarterAssetChangePercent,
		fin_data.BalanceSheet.AssetsGrowthRate,
		fin_data.ChartDataFinancials.Ticker,
		fin_data.BalanceSheet.TotalCurrentAssets,
		fin_data.BalanceSheet.CashAndShortTermInvestments,
		fin_data.BalanceSheet.CashAndShortTermInvestmentsGrowth), nil
}
