package provider

import (
	"fmt"
	"github.com/dijeferson/worth-report-dataloader/providers/web"
	"strings"
	"time"
)

const baseURL = "https://query1.finance.yahoo.com/v7/finance/download/%s?period1=%d&period2=%d&interval=%s&events=history"

// FetchDailyQuoteHistory fetches daily stock quote data from 1970-01-01 until today for a given symbol.
func FetchDailyQuoteHistory(symbol string) string {
	return fetchUntilNow(0, symbol, _1d)
}

func fetch(startDate int64, endDate int64, symbol string, interval Interval) string {

	if endDate == -1 {
		endDate = time.Now().Unix()
	}

	url := fmt.Sprintf(baseURL,
		strings.ToUpper(symbol),
		startDate,
		endDate,
		interval)

	contents := web.FetchURL(url)

	return contents
}

func fetchUntilNow(startDate int64, symbol string, interval Interval) string {
	return fetch(startDate, -1, symbol, interval)
}
