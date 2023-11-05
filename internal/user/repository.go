package user

import (
	"context"

	"github.com/22Fariz22/trueconf/internal/user/entity"
)

type Repo interface {
	CreateUser(ctx context.Context, newU entity.User) error
	DeleteUser(ctx context.Context)
	GetUser(ctx context.Context, id int) (*entity.UserStore, error)
	UpdateUser(ctx context.Context)
	SearchUsers(ctx context.Context) (*entity.UserStore, error)
}
