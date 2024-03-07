package service

import (
	"github.com/22Fariz22/mycloud/server/config"
	"github.com/22Fariz22/mycloud/server/internal/session"
	"github.com/22Fariz22/mycloud/server/internal/user"
	"github.com/22Fariz22/mycloud/server/pkg/logger"
	pb "github.com/22Fariz22/mycloud/server/proto"
)

type usersService struct {
	pb.UnimplementedUserServiceServer
	logger logger.Logger
	cfg    *config.Config
	userUC user.UserUseCase
	sessUC session.SessionUseCase
}

// NewAuthServerGRPC Auth service constructor
func NewAuthServerGRPC(
	logger logger.Logger, cfg *config.Config, userUC user.UserUseCase, sessUC session.SessionUseCase,
) *usersService {
	return &usersService{
		UnimplementedUserServiceServer: pb.UnimplementedUserServiceServer{},
		logger:                         logger,
		cfg:                            cfg,
		userUC:                         userUC,
		sessUC:                         sessUC,
	}
}
