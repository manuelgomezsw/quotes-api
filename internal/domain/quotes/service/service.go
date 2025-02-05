package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/domain/quotes/repository"
	"quotes-api/internal/infraestructure/client/firestore"
	"quotes-api/internal/infraestructure/client/mailersend"
	"quotes-api/internal/infraestructure/client/secretmanager"
	"quotes-api/internal/util/constant"
	"quotes-api/internal/util/customstrings"
)

func CreateQuoteService(quote *quotes.Quote) error {
	keywords, err := getTags(quote.Phrase)
	if err != nil {
		return err
	}
	formatQuote(quote, keywords)

	if err := repository.CreateQuote(quote); err != nil {
		return err
	}

	return nil
}

func UpdateQuoteService(quoteID int64, currentQuote *quotes.Quote) error {
	keywords, err := getTags(currentQuote.Phrase)
	if err != nil {
		return err
	}
	formatQuote(currentQuote, keywords)

	if err := repository.UpdateQuote(quoteID, currentQuote); err != nil {
		return err
	}

	return nil
}

func DeleteQuoteService(quoteID int64) error {
	if err := repository.DeleteQuote(quoteID); err != nil {
		return err
	}

	return nil
}

func GetQuoteByID(quoteID int64) (quotes.Quote, error) {
	quote, err := repository.GetQuoteByID(quoteID)
	if err != nil {
		return quotes.Quote{}, err
	}

	return quote, nil
}

func GetQuotesByKeyword(keyword string) ([]quotes.Quote, error) {
	quote, err := repository.GetQuotesByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func GetQuotesByAuthor(author string) ([]quotes.Quote, error) {
	quote, err := repository.GetQuotesByAuthor(author)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func GetQuotesByWork(work string) ([]quotes.Quote, error) {
	quote, err := repository.GetQuotesByWork(work)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func GetTopics() ([]string, error) {
	return repository.GetTopics()
}

func GetAuthors() ([]string, error) {
	return repository.GetAuthors()
}

func GetWorks() ([]string, error) {
	return repository.GetWorks()
}

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

func completeDataDailyQuote(quote *quotes.Quote) {
	if quote.Author == "" {
		quote.Author = constant.Desconocido
	}
}

func formatQuote(quote *quotes.Quote, keywords string) {
	quote.Author = customstrings.TrimSpace(quote.Author)
	quote.Phrase = customstrings.TrimSpace(quote.Phrase)
	quote.Work = customstrings.TrimSpace(quote.Work)

	quote.Author = customstrings.RemoveEndPeriod(quote.Author)
	quote.Work = customstrings.RemoveEndPeriod(quote.Work)
	quote.Phrase = customstrings.RemoveEndPeriod(quote.Phrase)

	quote.Tags = customstrings.RemoveSpecialCharacters(keywords)
}

func getTags(quote string) (string, error) {
	client := &http.Client{}
	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{"role": "user", "content": buildPrompt(quote)},
		},
		"temperature": 0.7,
	})

	openAIURL, err := firestore.GetValue(constant.OpenaiAPIURL)
	if err != nil {
		return "", err
	}

	openAIApiKey, err := secretmanager.GetValue(constant.OpenaiApiKey)
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("POST", openAIURL, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIApiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	responseJSON, _ := ioutil.ReadAll(resp.Body)
	var responseOpenAI quotes.OpenAIResponse

	if err := json.Unmarshal(responseJSON, &responseOpenAI); err != nil {
		return "", err
	}

	if len(responseOpenAI.Choices) > 0 {
		tags := responseOpenAI.Choices[0].Message.Content
		return tags, nil
	} else {
		return "", nil
	}
}

func buildPrompt(quote string) string {
	prompt := "Extrae los 4 tags clave que mejor representan la siguiente frase. Devuélvelas como una lista, separadas por comas: " + quote

	return prompt
}
