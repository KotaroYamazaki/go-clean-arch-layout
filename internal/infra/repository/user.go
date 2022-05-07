package repository

import (
	"context"

	entities "github.com/KotaroYamazaki/go-clean-arch-layout/internal/entities"
)

type UserRepository interface {
	Get(ctx context.Context, id int) (*entities.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (repo *userRepository) Get(ctx context.Context, id int) (*entities.User, error) {
	return nil, nil
}
