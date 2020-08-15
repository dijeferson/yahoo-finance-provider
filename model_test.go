package yahoofinance

import "testing"

type fields struct {
	Symbol   string
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	AdjClose float64
	Volume   float64
}

func TestStockQuote_ToJSON(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Case Success",
			fields: fields{Symbol: "ABCD", Date: "2020-01-01", Open: 100, High: 150, Low: 50, Close: 125, AdjClose: 130, Volume: 1000},
			want:   `{"Symbol":"ABCD","Date":"2020-01-01","Open":100,"High":150,"Low":50,"Close":125,"AdjClose":130,"Volume":1000}`,
		},
		{
			name:   "Case Empty",
			fields: fields{},
			want:   `{"Symbol":"","Date":"","Open":0,"High":0,"Low":0,"Close":0,"AdjClose":0,"Volume":0}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &StockQuote{
				Symbol:   tt.fields.Symbol,
				Date:     tt.fields.Date,
				Open:     tt.fields.Open,
				High:     tt.fields.High,
				Low:      tt.fields.Low,
				Close:    tt.fields.Close,
				AdjClose: tt.fields.AdjClose,
				Volume:   tt.fields.Volume,
			}
			if got := data.ToJSON(); got != tt.want {
				t.Errorf("StockQuote.ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStockQuote_ToString(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Case Success",
			fields: fields{Symbol: "ABCD", Date: "2020-01-01", Open: 100, High: 150, Low: 50, Close: 125, AdjClose: 130, Volume: 1000},
			want:   `ABCD => Data=2020-01-01, Open=100.000000, High=150.000000, Low=50.000000, Close=125.000000, AdjClose=130.000000, Volume=1000.000000`,
		},
		{
			name:   "Case Empty",
			fields: fields{},
			want:   ` => Data=, Open=0.000000, High=0.000000, Low=0.000000, Close=0.000000, AdjClose=0.000000, Volume=0.000000`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &StockQuote{
				Symbol:   tt.fields.Symbol,
				Date:     tt.fields.Date,
				Open:     tt.fields.Open,
				High:     tt.fields.High,
				Low:      tt.fields.Low,
				Close:    tt.fields.Close,
				AdjClose: tt.fields.AdjClose,
				Volume:   tt.fields.Volume,
			}
			if got := data.ToString(); got != tt.want {
				t.Errorf("StockQuote.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}