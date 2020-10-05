package models

import "time"

type Post struct {
	Id        uint64
	AuthorId  uint64
	Content   string
	CreatedAt time.Time
}
