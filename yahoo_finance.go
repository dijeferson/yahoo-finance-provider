package yahoofinance

import (
	"fmt"
	"strings"
	"time"

	"github.com/dijeferson/yahoo-finance-provider/interval"
)

const baseURL = "https://query1.finance.yahoo.com/v7/finance/download/%s?period1=%d&period2=%d&interval=%s&events=history"

// FetchDailyQuoteHistory fetches daily stock quote values from 1970-01-01 until today for a given symbol.
func FetchDailyQuoteHistory(symbol string) ([]StockQuote, error) {
	return FetchUntilNow(symbol, 0, interval.OneDays)
}

// FetchUntilNow fetches stock quote values for a given symbol from a startDate until now, separated by interval.
func FetchUntilNow(symbol string, startDate int64, interval interval.Interval) ([]StockQuote, error) {
	return Fetch(symbol, startDate, -1, interval)
}

// Fetch fetches stock quotes for a given symbol, in a given timespan (startDate and endDate), separated by interval.
func Fetch(symbol string, startDate int64, endDate int64, interval interval.Interval) ([]StockQuote, error) {
	if endDate == -1 {
		endDate = time.Now().Unix()
	}

	url := fmt.Sprintf(baseURL,
		strings.ToUpper(symbol),
		startDate,
		endDate,
		interval)

	contents, err := fetchURL(url)

	if err != nil {
		return nil, err
	}

	return parse(symbol, contents)
}

func parse(symbol string, data string) ([]StockQuote, error) {
	lines, err := readCSVFromString(data, true)
	stockList := make([]StockQuote, 0)

	if err != nil {
		return nil, fmt.Errorf("could not parse Stock data")
	}

	// Loop through lines & turn into object
	for _, columns := range lines {
		open, err := convertToFloat64(columns[1])
		if err != nil {
			return stockList, err
		}

		high, err := convertToFloat64(columns[2])
		if err != nil {
			return stockList, err
		}

		low, err := convertToFloat64(columns[3])
		if err != nil {
			return stockList, err
		}

		closure, err := convertToFloat64(columns[4])
		if err != nil {
			return stockList, err
		}

		adjClose, err := convertToFloat64(columns[5])
		if err != nil {
			return stockList, err
		}

		volume, err := convertToFloat64(columns[6])
		if err != nil {
			return stockList, err
		}

		data := StockQuote{
			Symbol:   strings.ToUpper(symbol),
			Date:     columns[0],
			Open:     open,
			High:     high,
			Low:      low,
			Close:    closure,
			AdjClose: adjClose,
			Volume:   volume,
		}

		stockList = append(stockList, data)
	}

	return stockList, nil
}
