package user

import (
	"context"

	"github.com/22Fariz22/trueconf/internal/user/entity"
)

type Repo interface {
	CreateUser(ctx context.Context, newU entity.User) error
	DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (*entity.UserStore, error)
	UpdateUser(ctx context.Context, id string, updUser entity.User) error
	SearchUsers(ctx context.Context) (*entity.UserStore, error)
}
