package usecase

import (
	"context"

	"github.com/alex-jienexa/velvet-scribes/internal/domain"
)

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Register(ctx context.Context, name, email, password string) (*domain.User, error) {
	// Валидация и бизнес-логика
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	hashedPassword := hashPassword(password)
	user := &domain.User{Name: name, Email: email, Password: hashedPassword}
	if err := uc.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
