package services

var (
	Quote *QuoteService
)

func init() {
	Quote = NewQuoteService()
}
