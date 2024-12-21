package quotes

import "time"

type Quote struct {
	QuoteID     int64     `json:"quote_id"`
	Author      string    `json:"author"`
	Work        string    `json:"work"`
	Phrase      string    `json:"phrase"`
	Tags        string    `json:"Tags"`
	DateCreated time.Time `json:"date_created"`
}

func (d Quote) GetDateCreatedFormatted() string {
	return d.DateCreated.Format("Enero 02, 2006")
}
