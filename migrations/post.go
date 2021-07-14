package migrations

import "time"

type Post struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Body      string    `json:"body" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
