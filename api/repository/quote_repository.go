package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"quotes-api/api/domain"
)

type QuoteRepository interface {
	GetLatestQuotes(limit int) string
}

func GetLatestQuotes(limit int) ([]domain.Quote, error) {
	filter := bson.D{{}}
	quotes := []domain.Quote{}

	mongoClient, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	quotesCollection := mongoClient.Database(DB).Collection(COLLECTION)
	cursorQuotes, findErr := quotesCollection.Find(context.TODO(), filter)
	if findErr != nil {
		return nil, findErr
	}

	for cursorQuotes.Next(context.TODO()) {
		quote := domain.Quote{}

		err := cursorQuotes.Decode(&quote)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	cursorQuotes.Close(context.TODO())

	return quotes, nil
}
