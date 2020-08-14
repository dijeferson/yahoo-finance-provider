package main

import (
	"encoding/csv"
	"fmt"
	"github.com/dijeferson/yahoo-finance-provider/converter"
	"github.com/dijeferson/yahoo-finance-provider/model"
	"log"
	"strings"
	"time"
)

const baseURL = "https://query1.finance.yahoo.com/v7/finance/download/%s?period1=%d&period2=%d&interval=%s&events=history"

// FetchDailyQuoteHistory fetches daily stock quote data from 1970-01-01 until today for a given symbol.
func FetchDailyQuoteHistory(symbol string) []model.StockQuote {
	return FetchUntilNow(0, symbol, _1d)
}

// FetchUntilNow ...
func FetchUntilNow(startDate int64, symbol string, interval Interval) []model.StockQuote {
	return Fetch(startDate, -1, symbol, interval)
}

// Fetch ...
func Fetch(startDate int64, endDate int64, symbol string, interval Interval) []model.StockQuote {
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

func parse(symbol string, data string) []model.StockQuote {
	lines, err := readCSVFromString(data, true)
	if err != nil {
		log.Panic(err)
	}

	stockList := make([]model.StockQuote, 0)

	// Loop through lines & turn into object
	for _, columns := range lines {
		data := model.StockQuote{
			Symbol:   strings.ToUpper(symbol),
			Date:     columns[0],
			Open:     converter.ConvertToFloat64(columns[1]),
			High:     converter.ConvertToFloat64(columns[2]),
			Low:      converter.ConvertToFloat64(columns[3]),
			Close:    converter.ConvertToFloat64(columns[4]),
			AdjClose: converter.ConvertToFloat64(columns[5]),
			Volume:   converter.ConvertToFloat64(columns[6]),
		}

		stockList = append(stockList, data)
	}

	return stockList
}

// ReadCSVFromString reads and parses CSV from a given string.
func readCSVFromString(contents string, removeHeader bool) ([][]string, error) {
	lines, err := csv.NewReader(strings.NewReader(contents)).ReadAll()

	if err != nil || len(lines) == 0 {
		return [][]string{}, err
	}

	if removeHeader {
		lines = lines[1:]
	}

	return lines, nil
}
