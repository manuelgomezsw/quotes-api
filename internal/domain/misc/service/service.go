package service

import "quotes-api/internal/domain/misc/repository"

func GetAuthors() ([]string, error) {
	return repository.GetAuthors()
}

func GetWorks() ([]string, error) {
	return repository.GetWorks()
}
