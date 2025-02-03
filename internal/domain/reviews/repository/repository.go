package repository

import (
	"quotes-api/internal/domain/reviews"
	"quotes-api/internal/util/mysql"
)

func Create(newReview *reviews.Review) error {
	newRecord, err := mysql.ClientDB.Exec(
		"INSERT INTO quotes.reviews (title, review) VALUES (?, ?)",
		newReview.Title,
		newReview.Review,
	)
	if err != nil {
		return err
	}

	newReview.ReviewID, err = newRecord.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func Update(currentReview *reviews.Review) error {
	_, err := mysql.ClientDB.Exec(
		"UPDATE quotes.reviews SET title = ?, review = ?  WHERE review_id = ?",
		currentReview.Title,
		currentReview.Review,
		currentReview.ReviewID,
	)
	if err != nil {
		return err
	}

	return nil
}

func Delete(reviewID int64) error {
	_, err := mysql.ClientDB.Exec(
		"DELETE FROM quotes.reviews WHERE review_id = ?",
		reviewID,
	)
	if err != nil {
		return err
	}

	return nil
}

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
