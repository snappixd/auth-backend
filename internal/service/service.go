package service

import (
	"auth-back/internal/repository"
	"auth-back/pkg/auth"
	"auth-back/pkg/hash"
	"context"
	"time"
)

type UserSignUpInput struct {
	Name     string
	Email    string
	Password string
}

type UserSignInInput struct {
	Email    string
	Password string
}

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (string, error)
}

type Services struct {
	Users Users
}

type Deps struct {
	Repos        *repository.Repositories
	Hasher       hash.PasswordHasher
	TokenManager auth.TokenManager
	TokenTTL     time.Duration
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users, deps.Hasher, deps.TokenManager, deps.TokenTTL)

	return &Services{
		Users: usersService,
	}
}
