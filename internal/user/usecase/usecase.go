package usecase

import (
	"context"
	"time"

	"github.com/22Fariz22/trueconf/internal/user"
	"github.com/22Fariz22/trueconf/internal/user/entity"
	"github.com/22Fariz22/trueconf/pkg/logger"
)

var l logger.Logger

type UseCase struct {
	repo user.Repo
}

func NewUseCaseUser(repo user.Repo) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) CreateUser(ctx context.Context, newU entity.User) error {
	newU.CreatedAt = time.Now()

	err := u.repo.CreateUser(ctx, newU)
	if err != nil {
		l.Errorf(err)
		return err
	}
	return nil
}

func (u *UseCase) DeleteUser(ctx context.Context) {}

func (u *UseCase) GetUser(ctx context.Context, id int) (*entity.UserStore, error) {
	data, err := u.GetUser(ctx, id)
	if err != nil {
		l.Errorf(err)
		return &entity.UserStore{}, err
	}

	return data, nil
}

func (u *UseCase) UpdateUser(ctx context.Context) {}

func (u *UseCase) SearchUsers(ctx context.Context) (*entity.UserStore, error) {
	data, err := u.repo.SearchUsers(ctx)
	if err != nil {
		l.Errorf(err)
		return nil, err
	}

	// res, err := json.Marshal(data)

	return data, nil
}
