package service

import (
	"errors"
	"quotes-api/internal/domain/words"
	"quotes-api/internal/domain/words/repository"
	"quotes-api/internal/util/apierror"
	"quotes-api/internal/util/cache"
	"quotes-api/internal/util/customstrings"
	"strings"
)

func Create(word *words.Word) apierror.ApiError {
	formatWord(word)

	if err := repository.Create(word); err != nil {
		return err
	}

	return nil
}

func Update(currentWord *words.Word) error {
	formatWord(currentWord)

	if err := repository.Update(currentWord); err != nil {
		return err
	}

	return nil
}

func Delete(wordID int64) error {
	if err := repository.Delete(wordID); err != nil {
		return err
	}

	return nil
}

func GetByID(quoteID int64) (words.Word, error) {
	quote, err := repository.GetByID(quoteID)
	if err != nil {
		return words.Word{}, err
	}

	return quote, nil
}

func GetByKeyword(keyword string) ([]words.Word, error) {
	quote, err := repository.GetByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func GetRandomWord() (words.Word, error) {
	item, err := cache.GetRandomItem("quote", getWordByIDWrapper, repository.GetMinMaxWords)
	if err != nil {
		return words.Word{}, err
	}

	// Convertimos el item a quotes.Quote
	word, ok := item.(words.Word)
	if !ok {
		return words.Word{}, errors.New("error de conversión al tipo word")
	}

	word.Word = customstrings.NewStringBuilder(word.Word).CapitalizeFirst().Build()

	return word, nil
}

// Wrapper para adaptar GetByID al tipo esperado por GetRandomItem
func getWordByIDWrapper(quoteID int64) (interface{}, error) {
	return GetByID(quoteID) // Retorna un `words.Word`, que es compatible con `interface{}`
}

func formatWord(word *words.Word) {
	word.Word = strings.TrimSpace(word.Word)
	word.Word = strings.ToLower(word.Word)
	word.Meaning = strings.TrimSpace(word.Meaning)
}
