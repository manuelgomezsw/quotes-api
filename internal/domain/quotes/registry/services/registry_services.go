package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/domain/quotes/registry/repository"
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

	req, _ := http.NewRequest("POST", os.Getenv(constant.OpenaiAPIURL), bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv(constant.OpenaiApiKey))

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
