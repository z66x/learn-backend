package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/z66x/url-shortener/internal/models"
)

type URLRepository struct {
	db *sqlx.DB
}

func NewURLRepository(db *sqlx.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(ctx context.Context, original string, shortCode string) (models.URL, error) {
	var url models.URL
	query := `
		INSERT INTO urls (original_url, short_code)
		VALUES ($1, $2)
		RETURNING *
	`
	err := r.db.QueryRowxContext(ctx, query, original, shortCode).StructScan(&url)
	return url, err
}

func (r *URLRepository) GetByShortCode(ctx context.Context, shortCode string) (models.URL, error) {
	var url models.URL
	query := `SELECT * FROM urls WHERE short_code = $1`
	err := r.db.QueryRowxContext(ctx, query, shortCode).StructScan(&url)
	return url, err
}