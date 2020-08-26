package yahoofinance

import (
	"fmt"
	"testing"
)

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

func TestHistoricQuote_ToJSON(t *testing.T) {
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
			data := &HistoricQuote{
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
				t.Errorf("HistoricQuote.ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistoricQuote_String(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Case Success",
			fields: fields{Symbol: "ABCD", Date: "2020-01-01", Open: 100, High: 150, Low: 50, Close: 125, AdjClose: 130, Volume: 1000},
			want:   `HistoricQuote ABCD => Data=2020-01-01, Open=100.000000, High=150.000000, Low=50.000000, Close=125.000000, AdjClose=130.000000, Volume=1000.000000`,
		},
		{
			name:   "Case Empty",
			fields: fields{},
			want:   `HistoricQuote  => Data=, Open=0.000000, High=0.000000, Low=0.000000, Close=0.000000, AdjClose=0.000000, Volume=0.000000`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &HistoricQuote{
				Symbol:   tt.fields.Symbol,
				Date:     tt.fields.Date,
				Open:     tt.fields.Open,
				High:     tt.fields.High,
				Low:      tt.fields.Low,
				Close:    tt.fields.Close,
				AdjClose: tt.fields.AdjClose,
				Volume:   tt.fields.Volume,
			}

			if got := fmt.Sprint(data); got != tt.want {
				t.Errorf("HistoricQuote.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
