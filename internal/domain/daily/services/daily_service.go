package services

import (
	"context"
	"quotes-api/internal/domain"
	"quotes-api/internal/domain/daily/repository"
	"quotes-api/internal/infraestructure/client/mailersend"
	"quotes-api/internal/util/constant"
	"time"
)

func SendDailyQuote(ctx context.Context) (string, error) {
	dailyQuote, err := repository.GetDailyQuote()
	if err != nil {
		return "", err
	}

	completeDataDailyQuote(&dailyQuote)
	formatDate(&dailyQuote)

	confirmationID, err := mailersend.SendMail(ctx, dailyQuote)
	if err != nil {
		return "", err
	}

	return confirmationID, nil
}

func completeDataDailyQuote(quote *domain.Quote) {
	if quote.Author == "" {
		quote.Author = constant.Desconocido
	}
}

func formatDate(quote *domain.Quote) {
	if quote.DateCreated != "" {
		castDate, err := time.Parse("January 02, 2006", quote.DateCreated)
		if err != nil {
			quote.DateCreated = castDate.String()
		}
	}
}
