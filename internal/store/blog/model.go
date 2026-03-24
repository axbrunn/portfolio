package blog

import (
	"database/sql"
	"time"
)

// blogPostModel representeert hoe een post uit de database komt.
// Dit type bestaat alleen in de store layer — het kent de quirks van MySQL,
// zoals nullable kolommen (sql.NullTime) die de business layer niet hoeft te weten.
type blogPostModel struct {
	ID        uint
	Title     string
	Slug      string
	Excerpt   string
	Body      string
	Published bool
	CreatedAt time.Time
	UpdatedAt time.Time
	// published_at is nullable in de DB: een post kan bestaan zonder publicatiedatum.
	PublishedAt sql.NullTime
}
