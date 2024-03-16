package services

import (
	"quotes-api/internal/domain/reviews"
	"quotes-api/internal/domain/reviews/search/repository"
)

func GetByID(reviewID int64) (reviews.Review, error) {
	review, err := repository.GetByID(reviewID)
	if err != nil {
		return reviews.Review{}, err
	}

	return review, nil
}
