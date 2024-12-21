package quotes

type Tag struct {
	ID      int64  `json:"id"`
	QuoteID string `json:"quote_id"`
	Tag     string `json:"tag"`
}
