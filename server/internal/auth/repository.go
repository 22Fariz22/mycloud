package auth

import (
	"context"

	"github.com/22Fariz22/mycloud/server/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, username, password string) (*entity.User, error)
}
