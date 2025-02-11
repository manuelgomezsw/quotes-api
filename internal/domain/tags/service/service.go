package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"quotes-api/internal/domain/quotes"
	"quotes-api/internal/infraestructure/client/firestore"
	"quotes-api/internal/infraestructure/client/secretmanager"
	"quotes-api/internal/util/constant"
)

func GetTagsAI(text string) (string, error) {
	client := &http.Client{}
	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{"role": "user", "content": buildPrompt(text)},
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
	prompt := "Extrae los 3 tags clave que mejor representan la siguiente frase. Devuélvelas como una lista, separadas por comas y sin caracteres especiales: " + quote

	return prompt
}
