package yahoofinance

import (
	"reflect"
	"testing"

	"github.com/dijeferson/yahoo-finance-provider/model"
)

func TestFetchDailyQuoteHistory(t *testing.T) {
	tests := []struct {
		name   string
		symbol string
		want   []model.StockQuote
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchDailyQuoteHistory(tt.symbol); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchDailyQuoteHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchUntilNow(t *testing.T) {
	type args struct {
		startDate int64
		symbol    string
		interval  Interval
	}
	tests := []struct {
		name string
		args args
		want []model.StockQuote
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchUntilNow(tt.args.startDate, tt.args.symbol, tt.args.interval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUntilNow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetch(t *testing.T) {
	type args struct {
		startDate int64
		endDate   int64
		symbol    string
		interval  Interval
	}
	tests := []struct {
		name string
		args args
		want []model.StockQuote
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fetch(tt.args.startDate, tt.args.endDate, tt.args.symbol, tt.args.interval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		symbol string
		data   string
	}
	tests := []struct {
		name string
		args args
		want []model.StockQuote
	}{
		{
			name: "Success with header",
			args: args{
				symbol: "ABCD",
				data:   "Date,Open,High,Low,Close,Adj Close,Volume\n2020-01-01,100.000,150,50,125.000,130.000000,1000",
			},
			want: []model.StockQuote{
				{Symbol: "ABCD", Date: "2020-01-01", Open: 100, High: 150, Low: 50, Close: 125, AdjClose: 130, Volume: 1000},
			},
		},
		{
			name: "Fail no header",
			args: args{
				symbol: "ABCD",
				data:   "2020-01-01,100.000,150,50,125.000,130.000000,1000",
			},
			want: []model.StockQuote{},
		},
		{
			name: "Fail empty data",
			args: args{
				symbol: "ABCD",
				data:   "",
			},
			want: []model.StockQuote{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.symbol, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readCSVFromString(t *testing.T) {
	type args struct {
		contents     string
		removeHeader bool
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Sucess no header",
			args: args{contents: "A,B", removeHeader: false},
			want: [][]string{
				{"A", "B"},
			},
			wantErr: false,
		},
		{
			name: "Sucess with header",
			args: args{contents: "field1,field2\nA,B", removeHeader: true},
			want: [][]string{
				{"A", "B"},
			},
			wantErr: false,
		},
		{
			name:    "Sucess with empty input and removeHeader",
			args:    args{contents: "", removeHeader: true},
			want:    [][]string{},
			wantErr: false,
		},
		{
			name:    "Sucess with empty input and dont removeHeader",
			args:    args{contents: "", removeHeader: false},
			want:    [][]string{},
			wantErr: false,
		},
		{
			name:    "Sucess with removeHeader when no header",
			args:    args{contents: "A,B", removeHeader: true},
			want:    [][]string{},
			wantErr: false,
		},
		{
			name:    "Error on bad number of columns",
			args:    args{contents: "A,B\nC,D,E", removeHeader: true},
			want:    [][]string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCSVFromString(tt.args.contents, tt.args.removeHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCSVFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCSVFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
