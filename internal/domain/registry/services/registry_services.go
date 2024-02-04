package services

import (
	"quotes-api/internal/domain"
	"quotes-api/internal/domain/registry/repository"
	"strings"
)

func CreateQuoteService(quote *domain.Quote) error {
	formatQuote(quote)

	if err := repository.CreateQuote(quote); err != nil {
		return err
	}

	return nil
}

func UpdateQuoteService(quoteID int64, currentQuote *domain.Quote) error {
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

func formatQuote(quote *domain.Quote) {
	trimSpaceQuote(quote)
	removeEndPeriodQuote(quote)
}

func trimSpaceQuote(quote *domain.Quote) {
	quote.Author = strings.TrimSpace(quote.Author)
	quote.Work = strings.TrimSpace(quote.Work)
	quote.Phrase = strings.TrimSpace(quote.Phrase)
}

func removeEndPeriodQuote(quote *domain.Quote) {
	quote.Author = removeEndPeriod(quote.Author)
	quote.Work = removeEndPeriod(quote.Work)
	quote.Phrase = removeEndPeriod(quote.Phrase)
}

func removeEndPeriod(word string) string {
	lastCharacter := word[len(word)-1:]
	if lastCharacter == "." {
		return word[0 : len(word)-1]
	}

	return word
}
