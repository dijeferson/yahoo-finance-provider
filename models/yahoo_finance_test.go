package model

import "testing"

func TestStockQuote_ToJSON(t *testing.T) {
	type fields struct {
		Symbol   string
		Date     string
		Open     string
		High     string
		Low      string
		Close    string
		AdjClose string
		Volume   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
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
	type fields struct {
		Symbol   string
		Date     string
		Open     string
		High     string
		Low      string
		Close    string
		AdjClose string
		Volume   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
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
