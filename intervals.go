package main

// Interval type
type Interval string

// Interval constants
// Based on valid intervals: [1m, 2m, 5m, 15m, 30m, 60m, 90m, 1h, 1d, 5d, 1wk, 1mo, 3mo]
const (
	oneMinutes     Interval = "1m"
	twoMinutes     Interval = "2m"
	fiveMinutes    Interval = "5m"
	fifteenMinutes Interval = "15m"
	thirtyMinutes  Interval = "30m"
	sixtyMinutes   Interval = "60m"
	ninetyMinutes  Interval = "90m"
	oneHour        Interval = "1h"
	oneDays        Interval = "1d"
	fiveDays       Interval = "5d"
	sevenDays      Interval = "1wk"
	thirtyDays     Interval = "1mo"
	ninetyDays     Interval = "3mo"

	// Also in another notation
	// .. I couldn't decide.
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
