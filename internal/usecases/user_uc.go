package usecases

import (
	"context"

	entities "github.com/KotaroYamazaki/go-clean-arch-layout/internal/entities"
	"github.com/KotaroYamazaki/go-clean-arch-layout/internal/infra/repository"
)

type UserUsecase interface {
	Get(ctx context.Context, id int) (*entities.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (uc *userUsecase) Get(ctx context.Context, id int) (*entities.User, error) {
	return uc.repo.Get(ctx, id)
}
