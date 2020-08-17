# Yahoo Finance Provider

[![Build Status](https://travis-ci.com/dijeferson/yahoo-finance-provider.svg?branch=mainline)](https://travis-ci.com/dijeferson/yahoo-finance-provider)
[![Go Report Card](https://goreportcard.com/badge/github.com/dijeferson/yahoo-finance-provider?)](https://goreportcard.com/report/github.com/dijeferson/yahoo-finance-provider)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/dijeferson/yahoo-finance-provider)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/dijeferson/yahoo-finance-provider#license-mit)

Go client for [YahooFinance](https://finance.yahoo.com) stock quote historical and real time data.

## Installation

    go get -u github.com/dijeferson/yahoo-finance-provider

## Usage and Examples

### 1. Historical stock quote data for MSFT 

```go
package main

import (
	"fmt"

	"github.com/dijeferson/yahoo-finance-provider"
)

func main() {
	response := yahoofinance.FetchDailyQuoteHistory("MSFT")

	for _, stock := range response {
		fmt.Println(stock.ToJSON())
	}
}
```

### 2. Latest stock quote data for MSFT 

Use the `interval` type to specify the aggregation. In the example below, the last 15 minutes of stock values for MSFT, aggregated by 5 min averages.  

```go
package main

import (
	"fmt"
	"time"
	"github.com/dijeferson/yahoo-finance-provider"
    "github.com/dijeferson/yahoo-finance-provider/interval"
)

func main() {
	// currentTime - 15 minutes
	start := time.Now().Unix() - 900
	response := yahoofinance.FetchUntilNow("MSFT", start, interval.FiveMinutes)

	for _, stock := range response {
		fmt.Println(stock.ToJSON())
	}
}
```

**Available Intervals**

- OneMinutes
- TwoMinutes
- FiveMinutes
- FifteenMinutes
- ThirtyMinutes
- SixtyMinutes
- NinetyMinutes
- OneHour
- OneDays
- FiveDays
- SevenDays
- ThirtyDays
- NinetyDays   
