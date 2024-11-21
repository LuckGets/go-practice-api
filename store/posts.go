package store

import (
	"context"
	"database/sql"

	"github.com/LuckGets/go-echo-rest/domain"
	"github.com/lib/pq"
)

type PostsStore struct {
	db *sql.DB
}

func (p *PostsStore) Create(ctx context.Context, post *domain.Post) error {
	query := `
	INSERT INTO posts (content, title, user_id,tags)
	VALUES ($1,$2,$3,$4) RETURNING id,created_at,updated_at
	`
	err := p.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags)).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
