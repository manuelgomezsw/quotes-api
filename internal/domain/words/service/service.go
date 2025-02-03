package service

import (
	"quotes-api/internal/domain/words"
	"quotes-api/internal/domain/words/repository"
	"quotes-api/internal/util/apierror"
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

func formatWord(word *words.Word) {
	word.Word = strings.TrimSpace(word.Word)
	word.Word = strings.ToLower(word.Word)
	word.Meaning = strings.TrimSpace(word.Meaning)
}
