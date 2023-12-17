package domain

type CompleteQuote struct {
	QuoteID     int64  `json:"quote_id"`
	Author      string `json:"author"`
	Work        string `json:"work"`
	Phrase      string `json:"phrase"`
	DateCreated string `json:"date_created"`
}
