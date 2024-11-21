package store

import (
	"context"
	"database/sql"

	"github.com/LuckGets/go-echo-rest/domain"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *domain.Post) error
	}
	Users interface {
		Create(context.Context, *domain.User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db},
		Users: &UsersStore{db},
	}
}
