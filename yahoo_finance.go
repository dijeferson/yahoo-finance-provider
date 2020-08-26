package yahoofinance

type method string

const (
	download     method = "download"
	quote        method = "quote"
	quoteSummary method = "quoteSummary"
)

var baseURLs = map[method]string{
	download:     "https://query1.finance.yahoo.com/v7/finance/download/%s",
	quote:        "https://query1.finance.yahoo.com/v7/finance/quote?symbols=%s",
	quoteSummary: "https://query1.finance.yahoo.com/v10/finance/quoteSummary/%s?formatted=false&modules=price%2CsummaryDetail%2CpageViews%2CfinancialsTemplate",
}
