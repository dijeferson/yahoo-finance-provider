package yahoofinance

type QuoteSummary struct {
	Result []QuoteSummaryResult `json:"result"`
	Error  interface{}          `json:"error"`
}

type QuoteSummaryResult struct {
	SummaryDetail      QuoteSummaryDetail             `json:"summaryDetail"`
	PageViews          QuoteSummaryPageViews          `json:"pageViews"`
	FinancialsTemplate QuoteSummaryFinancialsTemplate `json:"financialsTemplate"`
	Price              QuoteSummaryPrice              `json:"price"`
}

type QuoteSummaryFinancialsTemplate struct {
	Code   string `json:"code"`
	MaxAge int64  `json:"maxAge"`
}

type QuoteSummaryPageViews struct {
	ShortTermTrend string `json:"shortTermTrend"`
	MidTermTrend   string `json:"midTermTrend"`
	LongTermTrend  string `json:"longTermTrend"`
	MaxAge         int64  `json:"maxAge"`
}

type QuoteSummaryPrice struct {
	MaxAge                     int64       `json:"maxAge"`
	PreMarketChangePercent     float64     `json:"preMarketChangePercent"`
	PreMarketChange            float64     `json:"preMarketChange"`
	PreMarketTime              int64       `json:"preMarketTime"`
	PreMarketPrice             float64     `json:"preMarketPrice"`
	PreMarketSource            string      `json:"preMarketSource"`
	PostMarketChangePercent    float64     `json:"postMarketChangePercent"`
	PostMarketChange           float64     `json:"postMarketChange"`
	PostMarketTime             int64       `json:"postMarketTime"`
	PostMarketPrice            float64     `json:"postMarketPrice"`
	PostMarketSource           string      `json:"postMarketSource"`
	RegularMarketChangePercent float64     `json:"regularMarketChangePercent"`
	RegularMarketChange        float64     `json:"regularMarketChange"`
	RegularMarketTime          int64       `json:"regularMarketTime"`
	PriceHint                  int64       `json:"priceHint"`
	RegularMarketPrice         float64     `json:"regularMarketPrice"`
	RegularMarketDayHigh       float64     `json:"regularMarketDayHigh"`
	RegularMarketDayLow        float64     `json:"regularMarketDayLow"`
	RegularMarketVolume        int64       `json:"regularMarketVolume"`
	AverageDailyVolume10Day    int64       `json:"averageDailyVolume10Day"`
	AverageDailyVolume3Month   int64       `json:"averageDailyVolume3Month"`
	RegularMarketPreviousClose float64     `json:"regularMarketPreviousClose"`
	RegularMarketSource        string      `json:"regularMarketSource"`
	RegularMarketOpen          float64     `json:"regularMarketOpen"`
	Exchange                   string      `json:"exchange"`
	ExchangeName               string      `json:"exchangeName"`
	ExchangeDataDelayedBy      int64       `json:"exchangeDataDelayedBy"`
	MarketState                string      `json:"marketState"`
	QuoteType                  string      `json:"quoteType"`
	Symbol                     string      `json:"symbol"`
	UnderlyingSymbol           interface{} `json:"underlyingSymbol"`
	ShortName                  string      `json:"shortName"`
	LongName                   string      `json:"longName"`
	Currency                   string      `json:"currency"`
	QuoteSourceName            string      `json:"quoteSourceName"`
	CurrencySymbol             string      `json:"currencySymbol"`
	FromCurrency               interface{} `json:"fromCurrency"`
	ToCurrency                 interface{} `json:"toCurrency"`
	LastMarket                 interface{} `json:"lastMarket"`
	MarketCap                  int64       `json:"marketCap"`
}

type QuoteSummaryDetail struct {
	MaxAge                       int64       `json:"maxAge"`
	PriceHint                    int64       `json:"priceHint"`
	PreviousClose                float64     `json:"previousClose"`
	Open                         float64     `json:"open"`
	DayLow                       float64     `json:"dayLow"`
	DayHigh                      float64     `json:"dayHigh"`
	RegularMarketPreviousClose   float64     `json:"regularMarketPreviousClose"`
	RegularMarketOpen            float64     `json:"regularMarketOpen"`
	RegularMarketDayLow          float64     `json:"regularMarketDayLow"`
	RegularMarketDayHigh         float64     `json:"regularMarketDayHigh"`
	PayoutRatio                  float64     `json:"payoutRatio"`
	Beta                         float64     `json:"beta"`
	TrailingPE                   float64     `json:"trailingPE"`
	ForwardPE                    float64     `json:"forwardPE"`
	Volume                       int64       `json:"volume"`
	RegularMarketVolume          int64       `json:"regularMarketVolume"`
	AverageVolume                int64       `json:"averageVolume"`
	AverageVolume10Days          int64       `json:"averageVolume10days"`
	AverageDailyVolume10Day      int64       `json:"averageDailyVolume10Day"`
	Bid                          float64     `json:"bid"`
	Ask                          float64     `json:"ask"`
	BidSize                      int64       `json:"bidSize"`
	AskSize                      int64       `json:"askSize"`
	MarketCap                    int64       `json:"marketCap"`
	FiftyTwoWeekLow              float64     `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh             float64     `json:"fiftyTwoWeekHigh"`
	PriceToSalesTrailing12Months float64     `json:"priceToSalesTrailing12Months"`
	FiftyDayAverage              float64     `json:"fiftyDayAverage"`
	TwoHundredDayAverage         float64     `json:"twoHundredDayAverage"`
	Currency                     string      `json:"currency"`
	FromCurrency                 interface{} `json:"fromCurrency"`
	ToCurrency                   interface{} `json:"toCurrency"`
	LastMarket                   interface{} `json:"lastMarket"`
	Algorithm                    interface{} `json:"algorithm"`
	Tradeable                    bool        `json:"tradeable"`
}
