package service

import (
	"auth-back/internal/models"
	"auth-back/internal/repository"
	"auth-back/pkg/auth"
	"auth-back/pkg/hash"
	"context"
	"errors"
	"time"
)

type UsersService struct {
	repo   repository.Users
	hasher hash.PasswordHasher

	tokenManager auth.TokenManager

	TokenTTL time.Duration
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	tokenTL time.Duration) *UsersService {
	return &UsersService{
		repo:   repo,
		hasher: hasher,

		tokenManager: tokenManager,

		TokenTTL: tokenTL,
	}
}

func (s *UsersService) SignUp(ctx context.Context, input UserSignUpInput) error {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     input.Name,
		Password: passwordHash,
		Email:    input.Email,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		if errors.Is(err, models.ErrUserAlreadyExists) {
			return err
		}

		return err
	}

	return nil
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (string, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.repo.GetByCredentials(ctx, input.Email, passwordHash)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return "", err
		}

		return "", err
	}

	token, err := s.tokenManager.NewJWT(user.ID.Hex(), s.TokenTTL)
	if err != nil {
		return token, err
	}

	return token, nil
}
