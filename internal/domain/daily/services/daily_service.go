package services

import (
	"context"
	"quotes-api/internal/domain"
	"quotes-api/internal/domain/daily/repository"
	"quotes-api/internal/infraestructure/client/mailersend"
	"quotes-api/internal/util/constant"
)

func SendDailyQuote(ctx context.Context) (string, error) {
	dailyQuote, err := repository.GetDailyQuote()
	if err != nil {
		return "", err
	}

	completeDataDailyQuote(&dailyQuote)

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
