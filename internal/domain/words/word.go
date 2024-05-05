package words

import "time"

type Word struct {
	WordID      int64     `json:"word_id"`
	Word        string    `json:"word"`
	Meaning     string    `json:"meaning"`
	DateCreated time.Time `json:"date_created"`
}
