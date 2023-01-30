package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"quotes-api/api/domain"
)

type QuoteRepository interface {
	GetLatestQuotes(limit int) ([]domain.Quote, error)
	GetQuotesByTag(c context.Context, valueToFind string) ([]domain.Quote, error)
}

var (
	mongoClient      *mongo.Client
	errorMongoClient error
)

func GetQuotesByTag(c context.Context, valueToFind string) ([]domain.Quote, error) {
	var quotes []domain.Quote
	filter := bson.M{"$or": []bson.M{bson.M{"tags": valueToFind}, bson.M{"author": valueToFind}}}

	mongoClient, errorMongoClient = GetMongoClient()
	if errorMongoClient != nil {
		return nil, errorMongoClient
	}

	quotesCollection := mongoClient.Database(DB).Collection(COLLECTION)

	cursorQuotes, findErr := quotesCollection.Find(c, filter)
	if findErr != nil {
		return nil, findErr
	}

	for cursorQuotes.Next(c) {
		quote := domain.Quote{}

		err := cursorQuotes.Decode(&quote)
		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}
	defer cursorQuotes.Close(c)

	return quotes, nil
}

func GetLatestQuotes(limitQuery int) ([]domain.Quote, error) {
	var quotes []domain.Quote

	mongoClient, errorMongoClient = GetMongoClient()
	if errorMongoClient != nil {
		return nil, errorMongoClient
	}

	quotesCollection := mongoClient.Database(DB).Collection(COLLECTION)

	var optionsQuery *options.FindOptions
	if limitQuery > 0 {
		optionsQuery = options.Find().SetLimit(int64(limitQuery)).SetSort(bson.D{{"date_created", -1}})
	}

	cursorQuotes, findErr := quotesCollection.Find(context.TODO(), bson.D{{}}, optionsQuery)
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
	defer cursorQuotes.Close(context.TODO())

	return quotes, nil
}
