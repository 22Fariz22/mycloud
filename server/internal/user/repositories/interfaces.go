package repositories

import "context"


type UserSource interface{
	SignUp(ctx context.Context)error
	SignIn(ctx context.Context)error
	Logout(ctx context.Context)error
}