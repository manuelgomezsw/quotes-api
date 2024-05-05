package services

import (
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/domain/quotes/registry/repository"
	"strings"
)

func CreateQuoteService(quote *quotes.Quote) error {
	formatQuote(quote)

	if err := repository.CreateQuote(quote); err != nil {
		return err
	}

	return nil
}

func UpdateQuoteService(quoteID int64, currentQuote *quotes.Quote) error {
	formatQuote(currentQuote)

	if err := repository.UpdateQuote(quoteID, currentQuote); err != nil {
		return err
	}

	return nil
}

func DeleteQuoteService(quoteID int64) error {
	if err := repository.DeleteQuote(quoteID); err != nil {
		return err
	}

	return nil
}

func formatQuote(quote *quotes.Quote) {
	quote.Author = trimSpaceQuote(quote.Author)
	quote.Phrase = trimSpaceQuote(quote.Phrase)
	quote.Work = trimSpaceQuote(quote.Work)

	quote.Author = removeEndPeriod(quote.Author)
	quote.Work = removeEndPeriod(quote.Work)
	quote.Phrase = removeEndPeriod(quote.Phrase)
}

func trimSpaceQuote(value string) string {
	if value == "" {
		return value
	}

	return strings.TrimSpace(value)
}

func removeEndPeriod(value string) string {
	if value == "" {
		return value
	}

	lastCharacter := value[len(value)-1:]
	if lastCharacter == "." {
		return value[0 : len(value)-1]
	}

	return value
}
