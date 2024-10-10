package storage

import (
	"database/sql"
)

type Storage struct {
	Posts PostRepo
	Users UserRepo
}


func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Posts: &PostStore{db: db},
		Users: &UserStore{db: db},
	}
}
