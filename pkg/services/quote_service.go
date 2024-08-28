package services

import (
	"context"

	"github.com/dreadster3/goquotes/database"
	"github.com/dreadster3/goquotes/database/generated"
	"github.com/dreadster3/goquotes/pkg/utils/lists"
)

type QuoteService struct{}

func NewQuoteService() *QuoteService {
	return &QuoteService{}
}

func (qs *QuoteService) GetQuotes() ([]string, error) {
	ctx := context.Background()

	quotes, err := database.Queries.GetQuotes(ctx)
	if err != nil {
		return nil, err
	}

	return lists.Map(quotes, func(q generated.TbQuote) string { return q.Quote }), nil
}

func (qs *QuoteService) GetRandomQuote() (string, error) {
	ctx := context.Background()

	quote, err := database.Queries.GetRandomQuote(ctx)
	if err != nil {
		return "", err
	}

	return quote.Quote, nil
}

func (qs *QuoteService) CreateQuote(quote string) (string, error) {
	ctx := context.Background()

	createQuote, err := database.Queries.CreateQuote(ctx, quote)
	if err != nil {
		return "", err
	}

	return createQuote.Quote, nil
}
