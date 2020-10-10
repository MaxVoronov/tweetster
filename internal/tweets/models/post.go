package models

import "time"

type Post struct {
	ID        uint64
	AuthorID  uint64
	Content   string
	CreatedAt time.Time
}
