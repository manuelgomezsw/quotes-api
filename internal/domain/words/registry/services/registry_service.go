package services

import (
	"quotes-api/internal/domain/words"
	"quotes-api/internal/domain/words/registry/repository"
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

func formatWord(word *words.Word) {
	word.Word = strings.TrimSpace(word.Word)
	word.Word = strings.ToLower(word.Word)
	word.Meaning = strings.TrimSpace(word.Meaning)
}
