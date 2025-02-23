package reviews

import "time"

type Review struct {
	ReviewID    int       `json:"review_id"`
	Title       string    `json:"title"`
	Review      string    `json:"review"`
	Author      string    `json:"author"`
	Source      string    `json:"source"`
	Tags        []string  `json:"tags"`
	Column      bool      `json:"column"`
	DateCreated time.Time `json:"date_created"`
	Keywords    string    `json:"-"`
}
