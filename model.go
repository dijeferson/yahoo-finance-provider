package yahoofinance

import (
	"encoding/json"
	"fmt"
)

// StockQuote data type
type StockQuote struct {
	Symbol   string
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	AdjClose float64
	Volume   float64
}

// ToJSON generates the JSON serialization for the model
func (data *StockQuote) ToJSON() string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

// ToString generates a debuggable and easy to read string from the model
func (data *StockQuote) String() string {
	return fmt.Sprintf("%s => Data=%s, Open=%f, High=%f, Low=%f, Close=%f, AdjClose=%f, Volume=%f",
		data.Symbol,
		data.Date,
		data.Open,
		data.High,
		data.Low,
		data.Close,
		data.AdjClose,
		data.Volume)
}
