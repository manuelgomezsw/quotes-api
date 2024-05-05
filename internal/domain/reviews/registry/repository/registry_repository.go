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
