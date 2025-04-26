package user_entity

import (
	"context"

	"github.com/pimentafm/auction-go/internal/internal_error"
)

type User struct {
	Id   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, id string) (*User, *internal_error.InternalError)
}
