package usecase

import "context"


type UserInteraction interface{
	SignUp(ctx context.Context) error
	SignIn(ctx context.Context)error
	Logout(ctx context.Context)error
}