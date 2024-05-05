package reviews

import "time"

type Review struct {
	ReviewID    int64     `json:"review_id"`
	Title       string    `json:"title"`
	Review      string    `json:"review"`
	DateCreated time.Time `json:"date_created"`
}
