package mailersend

import (
	"context"
	"github.com/mailersend/mailersend-go"
	"os"
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/util/constant"
)

func SendMail(ctx context.Context, quote quotes.Quote) (string, error) {
	ms := mailersend.NewMailersend(os.Getenv(constant.MailersendApiKey))

	subject := constant.SenderSubject
	from := getFromSender()
	recipients := getRecipients()
	variables := getVariables()
	personalization := getPersonalization(quote)

	message := ms.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetTemplateID(os.Getenv(constant.EmailTemplateID))
	message.SetSubstitutions(variables)
	message.SetPersonalization(personalization)

	res, err := ms.Email.Send(ctx, message)
	if err != nil {
		return "", err
	}

	return res.Header.Get(constant.MessageID), nil
}

func getFromSender() mailersend.From {
	from := mailersend.From{}
	from.Name = constant.SenderName
	from.Email = constant.SenderEmail

	return from
}

func getRecipients() []mailersend.Recipient {
	var (
		recipients    []mailersend.Recipient
		manuRecipient mailersend.Recipient
		cataRecipient mailersend.Recipient
	)

	manuRecipient.Name = constant.RecipientManuName
	manuRecipient.Email = constant.RecipientManuEmail
	recipients = append(recipients, manuRecipient)

	cataRecipient.Name = constant.RecipientCataName
	cataRecipient.Email = constant.RecipientCataEmail
	recipients = append(recipients, cataRecipient)

	return recipients
}

func getVariables() []mailersend.Variables {
	return []mailersend.Variables{
		{
			Email: constant.RecipientManuEmail,
			Substitutions: []mailersend.Substitution{
				{
					Var:   "url",
					Value: constant.SenderUrlSite,
				},
			},
		},
		{
			Email: constant.RecipientCataEmail,
			Substitutions: []mailersend.Substitution{
				{
					Var:   "url",
					Value: constant.SenderUrlSite,
				},
			},
		},
	}
}

func getPersonalization(quote quotes.Quote) []mailersend.Personalization {
	var (
		manuDailyQuotePersonalization mailersend.Personalization
		cataDailyQuotePersonalization mailersend.Personalization

		personalizations []mailersend.Personalization
	)

	manuDailyQuotePersonalization.Email = constant.RecipientManuEmail
	manuDailyQuotePersonalization.Data = map[string]interface{}{
		"name":         constant.RecipientManuName,
		"work":         quote.Work,
		"quote":        quote.Phrase,
		"author":       quote.Author,
		"date_created": quote.GetDateCreatedFormatted(),
	}
	personalizations = append(personalizations, manuDailyQuotePersonalization)

	cataDailyQuotePersonalization.Email = constant.RecipientCataEmail
	cataDailyQuotePersonalization.Data = map[string]interface{}{
		"name":         constant.RecipientCataName,
		"work":         quote.Work,
		"quote":        quote.Phrase,
		"author":       quote.Author,
		"date_created": quote.GetDateCreatedFormatted(),
	}
	personalizations = append(personalizations, cataDailyQuotePersonalization)

	return personalizations
}
