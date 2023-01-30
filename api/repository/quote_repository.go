package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"quotes-api/api/domain"
)

type QuoteRepository interface {
	GetLatestQuotes(limit int) string
}

func GetLatestQuotes(limitQuery int) ([]domain.Quote, error) {
	filter := bson.D{{}}
	var quotes []domain.Quote

	mongoClient, err := GetMongoClient()
	if err != nil {
		return nil, err
	}

	quotesCollection := mongoClient.Database(DB).Collection(COLLECTION)

	var optionsQuery *options.FindOptions
	if limitQuery > 0 {
		optionsQuery = options.Find().SetLimit(int64(limitQuery)).SetSort(bson.D{{"date_created", -1}})
	}

	cursorQuotes, findErr := quotesCollection.Find(context.TODO(), filter, optionsQuery)
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
