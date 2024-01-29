package auth

import (
	"context"

	"github.com/22Fariz22/mycloud/server/entity"
)

const CtxUserKey = "user"

type UseCase interface{
	SignUp(ctx context.Context,username,password string)error
	SignIn(ctx context.Context,username,password string)(string,error)
	ParseToken(ctx context.Context,accessToken string)(*entity.User,error)
}