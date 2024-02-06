package usecase

import (
	"context"

	"github.com/22Fariz22/mycloud/server/internal/entity"
	"github.com/22Fariz22/mycloud/server/internal/user"
	"github.com/22Fariz22/mycloud/server/pkg/grpcerrors"
	"github.com/22Fariz22/mycloud/server/pkg/logger"
	userService "github.com/22Fariz22/passbook/server/proto"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	userByIDCacheDuration = 3600
)

// User UseCase
type userUseCase struct {
	logger     logger.Logger
	userPgRepo user.UserPGRepository
	redisRepo  user.UserRedisRepository
}

// NewUserUseCase New User UseCase
func NewUserUseCase(
	logger logger.Logger, userRepo user.UserPGRepository, redisRepo user.UserRedisRepository,
) *userUseCase {
	return &userUseCase{logger: logger, userPgRepo: userRepo, redisRepo: redisRepo}
}

// Register new user
func (u *userUseCase) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	existsUser, err := u.userPgRepo.FindByLogin(ctx, user.Login)
	if existsUser != nil || err == nil {
		return nil, grpcerrors.ErrEmailExists
	}

	return u.userPgRepo.Create(ctx, user)
}

// FindByLogin Find use by email address
func (u *userUseCase) FindByLogin(ctx context.Context, login string) (*entity.User, error) {
	findByLogin, err := u.userPgRepo.FindByLogin(ctx, login)
	if err != nil {
		return nil, errors.Wrap(err, "userPgRepo.FindByLogin")
	}

	findByLogin.SanitizePassword()

	return findByLogin, nil
}

// FindByID Find use by uuid
func (u *userUseCase) FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	cachedUser, err := u.redisRepo.GetByIDCtx(ctx, userID.String())
	if err != nil && !errors.Is(err, redis.Nil) {
		u.logger.Errorf("redisRepo.GetByIDCtx", err)
	}
	if cachedUser != nil {
		return cachedUser, nil
	}

	foundUser, err := u.userPgRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "userPgRepo.FindByID")
	}

	if err := u.redisRepo.SetUserCtx(ctx, foundUser.UserID.String(), userByIDCacheDuration, foundUser); err != nil {
		u.logger.Errorf("redisRepo.SetUserCtx", err)
	}

	return foundUser, nil
}

// Login user with email and password
func (u *userUseCase) Login(ctx context.Context, login string, password string) (*entity.User, error) {
	foundUser, err := u.userPgRepo.FindByLogin(ctx, login)
	if err != nil {
		return nil, errors.Wrap(err, "userPgRepo.FindByLogin")
	}

	if err := foundUser.ComparePasswords(password); err != nil {
		return nil, errors.Wrap(err, "user.ComparePasswords")
	}

	return foundUser, err
}

// AddBinary add binary data
func (u *userUseCase) AddBinary(ctx context.Context, userID string, request *userService.AddBinaryRequest) error {
	return u.userPgRepo.AddBinary(ctx, userID, request)
}

// GetByTitle find data by title
func (u *userUseCase) GetByTitle(ctx context.Context, userID string, request *userService.GetByTitleRequest) ([]string, error) {
	return u.userPgRepo.GetByTitle(ctx, userID, request)
}

// GetFullList find all type of data
func (u *userUseCase) GetFullList(ctx context.Context, userID string) ([]string, error) {
	return u.userPgRepo.GetFullList(ctx, userID)
}