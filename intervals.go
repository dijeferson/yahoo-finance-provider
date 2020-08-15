package yahoofinance

// Interval type
// Based on valid intervals: [1m, 2m, 5m, 15m, 30m, 60m, 90m, 1h, 1d, 5d, 1wk, 1mo, 3mo]
type Interval string

// Interval constants
// Based on valid intervals: [1m, 2m, 5m, 15m, 30m, 60m, 90m, 1h, 1d, 5d, 1wk, 1mo, 3mo]
const (
	OneMinutes     Interval = "1m"
	TwoMinutes     Interval = "2m"
	FiveMinutes    Interval = "5m"
	FifteenMinutes Interval = "15m"
	ThirtyMinutes  Interval = "30m"
	SixtyMinutes   Interval = "60m"
	NinetyMinutes  Interval = "90m"
	OneHour        Interval = "1h"
	OneDays        Interval = "1d"
	FiveDays       Interval = "5d"
	SevenDays      Interval = "1wk"
	ThirtyDays     Interval = "1mo"
	NinetyDays     Interval = "3mo"

	// Also in another notation
	// .. I couldn't decide. ¯\_(ツ)_/¯
	_1m  Interval = "1m"
	_2m  Interval = "2m"
	_5m  Interval = "5m"
	_15m Interval = "15m"
	_30m Interval = "30m"
	_60m Interval = "60m"
	_90m Interval = "90m"
	_1h  Interval = "1h"
	_1d  Interval = "1d"
	_5d  Interval = "5d"
	_7d  Interval = "1wk"
	_30d Interval = "1mo"
	_90d Interval = "3mo"
)
