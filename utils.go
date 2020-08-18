package yahoofinance

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// ConvertToFloat64 converts a string to float64. If it fails, zero is returned.
func convertToFloat64(value string) (float64, error) {
	result, err := strconv.ParseFloat(strings.TrimSpace(value), 64)

	if err != nil {
		return 0, fmt.Errorf("could not parse string %s into float", value)
	}

	return result, nil
}

// FetchURL fetches URL content and return as string.
func fetchURL(url string) (string, error) {
	contents, err := fetchBytesFromURL(url)

	if err != nil {
		return "", fmt.Errorf("could not fetch url %v. Resonse error %v", url, err)
	}

	return string(contents), nil
}

func fetchBytesFromURL(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	contents, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return []byte{}, err
	}

	return contents, nil
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
