package yahoofinance

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dijeferson/yahoo-finance-provider/interval"
)

const baseURL = "https://query1.finance.yahoo.com/v7/finance/download/%s"

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

	url := fmt.Sprintf(baseURL, strings.ToUpper(symbol))

	req, _ := http.NewRequest("GET", url, nil)
	queryParams := req.URL.Query()
	queryParams.Add("period1", fmt.Sprintf("%v", startDate))
	queryParams.Add("period2", fmt.Sprintf("%v", endDate))
	queryParams.Add("interval", fmt.Sprintf("%v", interval))
	queryParams.Add("events", "history")
	req.URL.RawQuery = queryParams.Encode()

	contents, err := fetchURL(req)

	if err != nil {
		return nil, err
	}

	return parse(symbol, contents)
}

// FetchURL fetches URL content and return as string.
func fetchURL(request *http.Request) (string, error) {
	contents, err := fetchBytesFromURL(request)

	if err != nil {
		return "", fmt.Errorf("could not fetch url %v. Response error %v", request.URL, err)
	}

	return string(contents), nil
}

func fetchBytesFromURL(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	res, err := client.Do(request)

	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte{}, err
	}

	return contents, nil
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
