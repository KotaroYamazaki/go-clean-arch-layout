package repository

import (
	"context"

	"github.com/KotaroYamazaki/go-clean-arch-layout/internal/entities/repository"

	entities "github.com/KotaroYamazaki/go-clean-arch-layout/internal/entities"
	"github.com/KotaroYamazaki/go-clean-arch-layout/internal/infra/db"
	orm "github.com/KotaroYamazaki/go-clean-arch-layout/pkg/orm"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (repo *userRepository) Get(ctx context.Context, id int) (*entities.User, error) {
	dbUser, err := orm.FindUser(ctx, db.GetContextExecutor(ctx), id)
	if err != nil {
		return nil, err
	}
	return &entities.User{
		User: *dbUser,
	}, nil
}
