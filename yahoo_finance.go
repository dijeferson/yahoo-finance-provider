package yahoofinance

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type method string

const (
	download     method = "download"
	quote        method = "quote"
	quoteSummary method = "quoteSummary"
)

var baseURLs = map[method]string{
	download:     "https://query1.finance.yahoo.com/v7/finance/download/%s",
	quote:        "https://query1.finance.yahoo.com/v7/finance/quote?symbols=%s",
	quoteSummary: "https://query1.finance.yahoo.com/v10/finance/quoteSummary/%s?formatted=false&modules=price%2CsummaryDetail%2CpageViews%2CfinancialsTemplate",
}

func processRequest(url string, params map[string]string) (string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	queryParams := req.URL.Query()

	for key, value := range params {
		queryParams.Add(key, value)
	}

	req.URL.RawQuery = queryParams.Encode()
	contents, err := fetchRequest(req)

	return contents, err
}

func fetchRequest(request *http.Request) (string, error) {
	contents, err := fetchBytesFromRequest(request)

	if err != nil {
		return "", fmt.Errorf("could not fetch url %v. Response error %v", request.URL, err)
	}

	return string(contents), nil
}

func fetchBytesFromRequest(request *http.Request) ([]byte, error) {
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