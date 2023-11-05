package user

import (
	"context"

	"github.com/22Fariz22/trueconf/internal/user/entity"
)

type UseCase interface {
	CreateUser(ctx context.Context)
	DeleteUser(ctx context.Context)
	GetUser(ctx context.Context)
	UpdateUser(ctx context.Context)
	SearchUsers(ctx context.Context) (*entity.UserStore, error)
}
