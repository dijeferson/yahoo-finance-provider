package provider

import "testing"

func TestFetchDailyQuoteHistory(t *testing.T) {
	type args struct {
		symbol string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchDailyQuoteHistory(tt.args.symbol); got != tt.want {
				t.Errorf("FetchDailyQuoteHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetch(t *testing.T) {
	type args struct {
		startDate int64
		endDate   int64
		symbol    string
		interval  Interval
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fetch(tt.args.startDate, tt.args.endDate, tt.args.symbol, tt.args.interval); got != tt.want {
				t.Errorf("fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetchUntilNow(t *testing.T) {
	type args struct {
		startDate int64
		symbol    string
		interval  Interval
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fetchUntilNow(tt.args.startDate, tt.args.symbol, tt.args.interval); got != tt.want {
				t.Errorf("fetchUntilNow() = %v, want %v", got, tt.want)
			}
		})
	}
}
