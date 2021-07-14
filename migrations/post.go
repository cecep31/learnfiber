package migrations

import "time"

type Post struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Createdat time.Time `json:"createdat"`
}
