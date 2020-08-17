package yahoofinance

import (
	"fmt"
	"github.com/dijeferson/yahoo-finance-provider/interval"
	"log"
	"strings"
	"time"
)

const baseURL = "https://query1.finance.yahoo.com/v7/finance/download/%s?period1=%d&period2=%d&interval=%s&events=history"

// FetchDailyQuoteHistory fetches daily stock quote values from 1970-01-01 until today for a given symbol.
func FetchDailyQuoteHistory(symbol string) []StockQuote {
	return FetchUntilNow(symbol, 0, interval.OneDays)
}

// FetchUntilNow fetches stock quote values for a given symbol from a startDate until now, separated by interval.
func FetchUntilNow(symbol string, startDate int64, interval interval.Interval) []StockQuote {
	return Fetch(symbol, startDate, -1, interval)
}

// Fetch fetches stock quotes for a given symbol, in a given timespan (startDate and endDate), separated by interval.
func Fetch(symbol string, startDate int64, endDate int64, interval interval.Interval) []StockQuote {
	if endDate == -1 {
		endDate = time.Now().Unix()
	}

	url := fmt.Sprintf(baseURL,
		strings.ToUpper(symbol),
		startDate,
		endDate,
		interval)

	contents := fetchURL(url)

	return parse(symbol, contents)
}

func parse(symbol string, data string) []StockQuote {
	lines, err := readCSVFromString(data, true)
	if err != nil {
		log.Panic(err)
	}

	stockList := make([]StockQuote, 0)

	// Loop through lines & turn into object
	for _, columns := range lines {
		data := StockQuote{
			Symbol:   strings.ToUpper(symbol),
			Date:     columns[0],
			Open:     convertToFloat64(columns[1]),
			High:     convertToFloat64(columns[2]),
			Low:      convertToFloat64(columns[3]),
			Close:    convertToFloat64(columns[4]),
			AdjClose: convertToFloat64(columns[5]),
			Volume:   convertToFloat64(columns[6]),
		}

		stockList = append(stockList, data)
	}

	return stockList
}
