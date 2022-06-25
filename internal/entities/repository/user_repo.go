package repository

import (
	"context"

	entities "github.com/KotaroYamazaki/go-clean-arch-layout/internal/entities"
)

type UserRepository interface {
	Get(ctx context.Context, id int) (*entities.User, error)
}
