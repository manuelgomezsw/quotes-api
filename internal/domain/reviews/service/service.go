package service

import (
	"quotes-api/internal/domain/reviews"
	"quotes-api/internal/domain/reviews/repository"
	tagsRepository "quotes-api/internal/domain/tags/repository"
	"quotes-api/internal/domain/tags/service"
	"quotes-api/internal/util/customstrings"
	"strings"
)

func Create(review *reviews.Review) error {
	keywords, err := service.GetTagsAI(review.Review)
	if err != nil {
		return err
	}
	formatReview(review, keywords)

	if err := repository.Create(review); err != nil {
		return err
	}

	if len(review.Tags) > 0 {
		if err = tagsRepository.CreateTags(0, review.ReviewID, review.Tags); err != nil {
			return err
		}
	}

	return nil
}

func Update(currentReview *reviews.Review) error {
	keywords, err := service.GetTagsAI(currentReview.Review)
	if err != nil {
		return err
	}
	formatReview(currentReview, keywords)

	if err := repository.Update(currentReview); err != nil {
		return err
	}

	if len(currentReview.Tags) > 0 {
		if err = tagsRepository.DeleteTags(0, currentReview.ReviewID); err != nil {
			return err
		}

		if err = tagsRepository.CreateTags(0, currentReview.ReviewID, currentReview.Tags); err != nil {
			return err
		}
	}

	return nil
}

func Delete(reviewID int64) error {
	if err := repository.Delete(reviewID); err != nil {
		return err
	}

	return nil
}

func Get() ([]reviews.Review, error) {
	allReviews, err := repository.Get()
	keywordsToTagsReviews(allReviews)

	if err != nil {
		return nil, err
	}

	return allReviews, nil
}

func GetByID(reviewID int64) (reviews.Review, error) {
	review, err := repository.GetByID(reviewID)
	if err != nil {
		return reviews.Review{}, err
	}
	keywordsToTagsReview(&review)

	return review, nil
}

func GetByTitle(title string) ([]reviews.Review, error) {
	reviewsByTitle, err := repository.GetByTitle(title)
	if err != nil {
		return nil, err
	}
	keywordsToTagsReviews(reviewsByTitle)

	return reviewsByTitle, nil
}

// formatReview ajusta el formato estándar y elimina caracteres especiales de la review.
func formatReview(review *reviews.Review, keywords string) {
	review.Review = customstrings.NewStringBuilder(review.Review).TrimSpace().RemoveEndPeriod().CapitalizeFirst().Build()
	review.Author = customstrings.NewStringBuilder(review.Author).TrimSpace().RemoveEndPeriod().CapitalizeFirst().Build()
	review.Source = customstrings.NewStringBuilder(review.Source).TrimSpace().RemoveEndPeriod().CapitalizeFirst().Build()
	review.Tags = strings.Split(customstrings.NewStringBuilder(keywords).RemoveSpecialCharacters().RemoveEndPeriod().Build(), ",")
}

// KeywordsToTagsReview convierte el campo Keywords de un review en un slice de strings y lo asigna a Tags.
func keywordsToTagsReview(review *reviews.Review) {
	if review == nil {
		return
	}
	review.Tags = strings.Split(review.Keywords, ",")
}

// KeywordsToTagsReviews recorre una slice de reviews y aplica la conversión a cada uno.
func keywordsToTagsReviews(reviews []reviews.Review) {
	for i := range reviews {
		keywordsToTagsReview(&reviews[i])
	}
}
