package model

import (
	"encoding/json"
	"fmt"
	"log"
)

// StockQuote data type
type StockQuote struct {
	Symbol   string
	Date     string
	Open     string
	High     string
	Low      string
	Close    string
	AdjClose string
	Volume   string
}

// ToJSON generates the JSON serialization for the model
func (data *StockQuote) ToJSON() string {
	jsonData, err := json.Marshal(data)

	if err != nil {

		log.Fatal(err)
	}

	return string(jsonData)
}

// ToString generates a debugable and easy to read string from the model
func (data *StockQuote) ToString() string {
	return fmt.Sprintf("%s => Data=%s, Open=%s, High=%s, Low=%s, Close=%s, AdjClose=%s, Volume=%s",
		data.Symbol,
		data.Date,
		data.Open,
		data.High,
		data.Low,
		data.Close,
		data.AdjClose,
		data.Volume)
}
