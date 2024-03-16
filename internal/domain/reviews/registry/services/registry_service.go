package services

import (
	"quotes-api/internal/domain/reviews"
	"quotes-api/internal/domain/reviews/registry/repository"
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

func formatReview(review *reviews.Review) {
	review.Title = strings.TrimSpace(review.Title)
	review.Review = strings.TrimSpace(review.Review)
}
