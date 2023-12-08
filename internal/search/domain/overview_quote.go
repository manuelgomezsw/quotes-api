package domain

type OverviewQuote struct {
	QuoteID string `json:"quote_id"`
	Author  string `json:"author"`
	Phrase  string `json:"phrase"`
}
