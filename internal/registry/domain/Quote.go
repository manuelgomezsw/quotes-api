package domain

type Quote struct {
	QuoteID     string   `json:"quote_id"`
	Author      string   `json:"author"`
	Work        string   `json:"work"`
	Phrase      string   `json:"phrase"`
	Tags        []string `json:"tags"`
	DateCreated string   `json:"date_created"`
}
