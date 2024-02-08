package user

import (
	"context"
	"github.com/google/uuid"

	userService "github.com/22Fariz22/mycloud/proto"
	"github.com/22Fariz22/mycloud/server/internal/entity"
)

//go:generate mockgen -source pg-repository.go -destination mock/pg_repository.go -package mock

// UserPGRepository User pg repository
type UserPGRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindByLogin(ctx context.Context, login string) (*entity.User, error)
	FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error)

	AddBinary(ctx context.Context, userID string, request *userService.AddBinaryRequest) error
	GetByTitle(ctx context.Context, userID string, request *userService.GetByTitleRequest) ([]string, error)
	GetFullList(ctx context.Context, userID string) ([]string, error)
}
