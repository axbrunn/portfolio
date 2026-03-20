package blog

import "time"

type BlogPost struct {
	ID          uint
	Title       string
	Slug        string
	Excerpt     string
	Body        string
	Published   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt *time.Time
}
