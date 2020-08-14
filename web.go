package yahoofinance

import (
	"io/ioutil"
	"log"
	"net/http"
)

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
