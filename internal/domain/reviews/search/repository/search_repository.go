package repository

import (
	"quotes-api/internal/domain/reviews"
	"quotes-api/internal/util/mysql"
)

func GetByID(reviewID int64) (reviews.Review, error) {
	resultReview, err := mysql.ClientDB.Query(
		"SELECT review_id, title, review, date_created FROM quotes.reviews WHERE review_id = ?", reviewID)
	if err != nil {
		return reviews.Review{}, err
	}

	var review reviews.Review
	for resultReview.Next() {
		err = resultReview.Scan(&review.ReviewID, &review.Title, &review.Review, &review.DateCreated)
		if err != nil {
			return reviews.Review{}, err
		}
	}

	return review, nil
}

func GetByTitle(title string) ([]reviews.Review, error) {
	resultReview, err := mysql.ClientDB.Query(
		"SELECT review_id, title, review, date_created FROM quotes.reviews WHERE title LIKE ?", "%"+title+"%")
	if err != nil {
		return nil, err
	}

	var reviewsSearched []reviews.Review
	for resultReview.Next() {
		var review reviews.Review

		err = resultReview.Scan(&review.ReviewID, &review.Title, &review.Review, &review.DateCreated)
		if err != nil {
			return nil, err
		}

		reviewsSearched = append(reviewsSearched, review)
	}

	return reviewsSearched, nil
}
