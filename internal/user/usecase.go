package user

import (
	"context"

	"github.com/22Fariz22/trueconf/internal/user/entity"
	"github.com/22Fariz22/trueconf/pkg/logger"
)

// UseCase итрерфейс service слоя
type UseCase interface {
	CreateUser(ctx context.Context, l logger.Logger, newU *entity.User) error
	DeleteUser(ctx context.Context, l logger.Logger, id string) error
	GetUser(ctx context.Context, l logger.Logger, id string) (*entity.UserStore, error)
	UpdateUser(ctx context.Context, l logger.Logger, id string, updUser *entity.User) error
	SearchUsers(ctx context.Context, l logger.Logger) (*entity.UserStore, error)
}
