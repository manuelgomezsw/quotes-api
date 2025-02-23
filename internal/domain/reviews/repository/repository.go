package repository

import (
	"fmt"
	"os"
	"quotes-api/internal/domain/reviews"
	"quotes-api/internal/util/conversions"
	"quotes-api/internal/util/mysql"
)

const (
	basePathSqlQueries = "sql/reviews"

	fileSqlCreate     = "Create.sql"
	fileSqlUpdate     = "Update.sql"
	fileSqlDelete     = "Delete.sql"
	fileSqlGetByID    = "GetByID.sql"
	fileSqlGetByTitle = "GetByTitle.sql"
	fileSqlGet        = "Get.sql"
)

func Create(newReview *reviews.Review) error {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlCreate))
	if err != nil {
		return err
	}

	newRecord, err := mysql.ClientDB.Exec(
		string(query),
		newReview.Title,
		newReview.Review,
		newReview.Author,
		newReview.Source,
		newReview.Column,
	)
	if err != nil {
		return err
	}

	lastID, err := newRecord.LastInsertId()
	if err != nil {
		return err
	}

	newReview.ReviewID, err = conversions.SafeIntConversion(lastID)
	if err != nil {
		return err
	}

	return nil
}

func Update(currentReview *reviews.Review) error {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlUpdate))
	if err != nil {
		return err
	}

	_, err = mysql.ClientDB.Exec(
		string(query),
		currentReview.Title,
		currentReview.Review,
		currentReview.Author,
		currentReview.Source,
		currentReview.Column,
		currentReview.ReviewID,
	)
	if err != nil {
		return err
	}

	return nil
}

func Delete(reviewID int64) error {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlDelete))
	if err != nil {
		return err
	}

	_, err = mysql.ClientDB.Exec(
		string(query),
		reviewID,
	)
	if err != nil {
		return err
	}

	return nil
}

func Get() ([]reviews.Review, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGet))
	if err != nil {
		return nil, err
	}

	resultReview, err := mysql.ClientDB.Query(string(query))
	if err != nil {
		return nil, err
	}

	var reviewsSearched []reviews.Review
	for resultReview.Next() {
		var review reviews.Review

		err = resultReview.Scan(
			&review.ReviewID,
			&review.Title,
			&review.Review,
			&review.Author,
			&review.Source,
			&review.Keywords,
			&review.Column,
			&review.DateCreated,
		)
		if err != nil {
			return nil, err
		}

		reviewsSearched = append(reviewsSearched, review)
	}

	return reviewsSearched, nil
}

func GetByID(reviewID int64) (reviews.Review, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetByID))
	if err != nil {
		return reviews.Review{}, err
	}

	resultReview, err := mysql.ClientDB.Query(string(query), reviewID)
	if err != nil {
		return reviews.Review{}, err
	}

	var review reviews.Review
	for resultReview.Next() {
		err = resultReview.Scan(
			&review.ReviewID,
			&review.Title,
			&review.Review,
			&review.Author,
			&review.Source,
			&review.Keywords,
			&review.Column,
			&review.DateCreated,
		)
		if err != nil {
			return reviews.Review{}, err
		}
	}

	return review, nil
}

func GetByTitle(title string) ([]reviews.Review, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetByTitle))
	if err != nil {
		return nil, err
	}

	resultReview, err := mysql.ClientDB.Query(string(query), "%"+title+"%")
	if err != nil {
		return nil, err
	}

	var reviewsSearched []reviews.Review
	for resultReview.Next() {
		var review reviews.Review

		err = resultReview.Scan(
			&review.ReviewID,
			&review.Title,
			&review.Review,
			&review.Author,
			&review.Source,
			&review.Keywords,
			&review.Column,
			&review.DateCreated,
		)
		if err != nil {
			return nil, err
		}

		reviewsSearched = append(reviewsSearched, review)
	}

	return reviewsSearched, nil
}
