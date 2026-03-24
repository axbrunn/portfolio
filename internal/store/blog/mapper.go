package blog

import (
	"database/sql"

	bizblog "github.com/axbrunn/portfolio/internal/business/blog"
)

// toModel zet een business BlogPost om naar het store-interne model.
// Gebruik dit voordat je iets naar de DB schrijft (Insert, Update).
// Zo hoeft de business layer niets te weten over sql.NullTime of andere DB-types.
func toModel(p bizblog.BlogPost) blogPostModel {
	m := blogPostModel{
		ID:        p.ID,
		Title:     p.Title,
		Slug:      p.Slug,
		Excerpt:   p.Excerpt,
		Body:      p.Body,
		Published: p.Published,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}

	// Converteer *time.Time naar sql.NullTime voor de nullable DB kolom.
	if p.PublishedAt != nil {
		m.PublishedAt = sql.NullTime{Time: *p.PublishedAt, Valid: true}
	}

	return m
}

// toBusiness zet een store-intern model om naar een business BlogPost.
// Gebruik dit nadat je iets uit de DB hebt gelezen (Select*).
// De business layer ontvangt altijd schone Go-types, nooit DB-specifieke types.
func toBusiness(m blogPostModel) bizblog.BlogPost {
	p := bizblog.BlogPost{
		ID:        m.ID,
		Title:     m.Title,
		Slug:      m.Slug,
		Excerpt:   m.Excerpt,
		Body:      m.Body,
		Published: m.Published,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	// Converteer sql.NullTime terug naar *time.Time.
	// Alleen invullen als de waarde daadwerkelijk aanwezig was in de DB.
	if m.PublishedAt.Valid {
		p.PublishedAt = &m.PublishedAt.Time
	}

	return p
}
