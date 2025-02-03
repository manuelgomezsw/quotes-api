package service

import (
	"quotes-api/internal/domain/reviews"
	"quotes-api/internal/domain/reviews/repository"
	"strings"
)

func Create(review *reviews.Review) error {
	formatReview(review)

	if err := repository.Create(review); err != nil {
		return err
	}

	return nil
}

func Update(currentReview *reviews.Review) error {
	formatReview(currentReview)

	if err := repository.Update(currentReview); err != nil {
		return err
	}

	return nil
}

func Delete(reviewID int64) error {
	if err := repository.Delete(reviewID); err != nil {
		return err
	}

	return nil
}

func GetByID(reviewID int64) (reviews.Review, error) {
	review, err := repository.GetByID(reviewID)
	if err != nil {
		return reviews.Review{}, err
	}

	return review, nil
}

func GetByTitle(title string) ([]reviews.Review, error) {
	reviewsByTitle, err := repository.GetByTitle(title)
	if err != nil {
		return nil, err
	}

	return reviewsByTitle, nil
}

func formatReview(review *reviews.Review) {
	review.Title = strings.TrimSpace(review.Title)
	review.Review = strings.TrimSpace(review.Review)
}
