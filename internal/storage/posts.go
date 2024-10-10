package storage

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// PostStore is a struct that contains a database connection.
type PostStore struct {
	db *sql.DB
}

// Post is a struct that represents a post.
type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int      `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func (p *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at, updated_at
	`
	err := p.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags)).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
