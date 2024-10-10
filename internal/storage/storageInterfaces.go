package storage

import (
	"context"
)

type PostRepo interface {
	Create(ctx context.Context, post *Post) (error)
}

type UserRepo interface {
	Create(ctx context.Context, user *User) (error)
}