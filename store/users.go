package store

import (
	"context"
	"database/sql"

	"github.com/LuckGets/go-echo-rest/domain"
)

type UsersStore struct {
	db *sql.DB
}

func (us *UsersStore) Create(ctx context.Context, user *domain.User) error {
	query := `
	INSERT INTO users (fullname,lastname,email,password,email)	
	VALUES ($1,$2,$3,$4)
	RETURNING id,created_at
	`

	err := us.db.QueryRowContext(ctx, query, user.Fullname, user.Lastname, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
