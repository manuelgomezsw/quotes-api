package service

import (
	"errors"
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/domain/quotes/repository"
	tagsRepository "quotes-api/internal/domain/tags/repository"
	"quotes-api/internal/domain/tags/service"
	"quotes-api/internal/util/cache"
	"quotes-api/internal/util/customstrings"
	"strings"
)

func Create(newQuote *quotes.Quote) error {
	keywords, err := service.GetTagsAI(newQuote.Phrase)
	if err != nil {
		return err
	}
	formatQuote(newQuote, keywords)

	if err := repository.CreateQuote(newQuote); err != nil {
		return err
	}

	if len(newQuote.Tags) > 0 {
		if err = tagsRepository.CreateTags(newQuote.QuoteID, 0, newQuote.Tags); err != nil {
			return err
		}
	}

	return nil
}

func Update(quoteID int64, currentQuote *quotes.Quote) error {
	keywords, err := service.GetTagsAI(currentQuote.Phrase)
	if err != nil {
		return err
	}
	formatQuote(currentQuote, keywords)

	if err := repository.UpdateQuote(quoteID, currentQuote); err != nil {
		return err
	}

	if len(currentQuote.Tags) > 0 {
		if err = tagsRepository.DeleteTags(quoteID, 0); err != nil {
			return err
		}

		if err = tagsRepository.CreateTags(quoteID, 0, currentQuote.Tags); err != nil {
			return err
		}
	}

	return nil
}

func Delete(quoteID int64) error {
	if err := repository.DeleteQuote(quoteID); err != nil {
		return err
	}

	return nil
}

func GetByID(quoteID int64) (quotes.Quote, error) {
	quote, err := repository.GetQuoteByID(quoteID)
	if err != nil {
		return quotes.Quote{}, err
	}
	keywordsToTagsQuote(&quote)

	return quote, nil
}

func GetByKeyword(keyword string) ([]quotes.Quote, error) {
	quotesByKeyword, err := repository.GetQuotesByKeyword(keyword)
	if err != nil {
		return nil, err
	}
	keywordsToTagsQuotes(quotesByKeyword)

	return quotesByKeyword, nil
}

func GetByAuthor(author string) ([]quotes.Quote, error) {
	quoteByAuthor, err := repository.GetQuotesByAuthor(author)
	if err != nil {
		return nil, err
	}
	keywordsToTagsQuotes(quoteByAuthor)

	return quoteByAuthor, nil
}

func GetByWork(work string) ([]quotes.Quote, error) {
	quotesByWork, err := repository.GetQuotesByWork(work)
	if err != nil {
		return nil, err
	}
	keywordsToTagsQuotes(quotesByWork)

	return quotesByWork, nil
}

func GetRandomQuote() (quotes.Quote, error) {
	item, err := cache.GetRandomItem("quote", getQuoteByIDWrapper, repository.GetMinMaxQuotes)
	if err != nil {
		return quotes.Quote{}, err
	}

	// Convertimos el item a quotes.Quote
	quote, ok := item.(quotes.Quote)
	if !ok {
		return quotes.Quote{}, errors.New("error de conversión al tipo Quote")
	}
	keywordsToTagsQuote(&quote)

	return quote, nil
}

// getQuoteByIDWrapper adapta GetQuoteByID al tipo esperado por GetRandomItem.
func getQuoteByIDWrapper(quoteID int64) (interface{}, error) {
	return GetByID(quoteID) // Retorna un `quotes.Quote`, que es compatible con `interface{}`
}

// formatQuote limpia los caracteres especiales y espacios que retorna OpenAI.
func formatQuote(quote *quotes.Quote, keywords string) {
	quote.Author = customstrings.NewStringBuilder(quote.Author).TrimSpace().RemoveEndPeriod().CapitalizeFirst().Build()
	quote.Phrase = customstrings.NewStringBuilder(quote.Phrase).TrimSpace().RemoveEndPeriod().CapitalizeFirst().Build()
	quote.Work = customstrings.NewStringBuilder(quote.Work).TrimSpace().RemoveEndPeriod().CapitalizeFirst().Build()
	quote.Tags = strings.Split(customstrings.NewStringBuilder(keywords).RemoveSpecialCharacters().Build(), ",")
}

// keywordsToTagsQuote convierte el campo Keywords de una quote en un slice de strings y lo asigna a Tags.
func keywordsToTagsQuote(quote *quotes.Quote) {
	if quote == nil {
		return
	}
	quote.Tags = strings.Split(quote.Keywords, ",")
}

// keywordsToTagsQuotes recorre una slice de quotes y aplica la conversión a cada uno.
func keywordsToTagsQuotes(quotes []quotes.Quote) {
	for i := range quotes {
		keywordsToTagsQuote(&quotes[i])
	}
}
