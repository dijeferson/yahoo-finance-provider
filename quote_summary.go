package yahoofinance

import "log"

// FetchQuoteSummaries ...
func FetchQuoteSummaries(symbols ...string) []QuoteSummary {
	result := make([]QuoteSummary, 0)
	for _, symbol := range symbols {
		FetchQuoteSummary(symbol)
	}

	return result
}

// FetchQuoteSummary ...
func FetchQuoteSummary(symbol string) QuoteSummary {
	result := QuoteSummary{}

	log.Print(symbol)

	return result
}
