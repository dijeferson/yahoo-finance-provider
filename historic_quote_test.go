package yahoofinance

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		symbol string
		data   string
	}
	tests := []struct {
		name string
		args args
		want []HistoricQuote
	}{
		{
			name: "Success with header",
			args: args{
				symbol: "ABCD",
				data:   "Date,Open,High,Low,Close,Adj Close,Volume\n2020-01-01,100.000,150,50,125.000,130.000000,1000",
			},
			want: []HistoricQuote{
				{Symbol: "ABCD", Date: "2020-01-01", Open: 100, High: 150, Low: 50, Close: 125, AdjClose: 130, Volume: 1000},
			},
		},
		{
			name: "Fail no header",
			args: args{
				symbol: "ABCD",
				data:   "2020-01-01,100.000,150,50,125.000,130.000000,1000",
			},
			want: []HistoricQuote{},
		},
		{
			name: "Fail empty data",
			args: args{
				symbol: "ABCD",
				data:   "",
			},
			want: []HistoricQuote{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := parse(tt.args.symbol, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
