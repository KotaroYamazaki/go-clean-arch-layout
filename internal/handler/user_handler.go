package handler

import (
	"context"

	entities "github.com/KotaroYamazaki/go-clean-arch-layout/internal/entities"
	"github.com/KotaroYamazaki/go-clean-arch-layout/internal/usecases"
)

type userHandler struct {
	userUC usecases.UserUsecase
}

func NewUserHandler(userUC usecases.UserUsecase) *userHandler {
	return &userHandler{
		userUC: userUC,
	}
}

func (h *userHandler) Get(ctx context.Context, id int) (*entities.User, error) {
	return h.userUC.Get(ctx, id)
}
