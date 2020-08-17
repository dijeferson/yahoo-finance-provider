package yahoofinance

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ConvertToFloat64 converts a string to float64. If it fails, zero is returned.
func convertToFloat64(value string) float64 {
	result, err := strconv.ParseFloat(strings.TrimSpace(value), 64)

	if err != nil {
		log.Printf("Error %v. Returning value = 0", err)
		return 0
	}

	return result
}

// FetchURL fetches URL content and return as string.
func fetchURL(url string) string {
	contents, err := fetchBytesFromURL(url)

	if err != nil {
		log.Fatalf("Error while fetching url %s. Error: %v", url, err)
	}

	return string(contents)
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
