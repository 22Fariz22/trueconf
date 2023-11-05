package usecase

import (
	"context"
	"encoding/json"

	"github.com/22Fariz22/trueconf/internal/user"
	"github.com/22Fariz22/trueconf/internal/user/entity"
)

var l logger.Logger

type UseCase struct {
	repo user.Repo
}

func NewUseCaseUser(repo user.Repo) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) CreateUser(ctx context.Context) {}
func (u *UseCase) DeleteUser(ctx context.Context) {}
func (u *UseCase) GetUser(ctx context.Context)    {}
func (u *UseCase) UpdateUser(ctx context.Context) {}

func (u *UseCase) SearchUsers(ctx context.Context) (*entity.UserStore, error) {
	data, err := u.repo.SearchUsers(ctx)
	if err != nil {
		l.Error(err)
		return nil, err
	}

	// res, err := json.Marshal(data)

	return data, nil
}
