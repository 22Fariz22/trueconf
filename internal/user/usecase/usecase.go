package usecase

import (
	"context"
	"fmt"
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

func (u *UseCase) DeleteUser(ctx context.Context, id string) error {
	err := u.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) GetUser(ctx context.Context, id string) (*entity.UserStore, error) {
	fmt.Println("UC GetUser().")
	data, err := u.GetUser(ctx, id)
	if err != nil {
		l.Errorf(err)
		return &entity.UserStore{}, err
	}

	return data, nil
}

func (u *UseCase) UpdateUser(ctx context.Context, id string, updUser entity.User) error {
	err := u.repo.UpdateUser(ctx, id, updUser)
	if err != nil {
		l.Errorf(err)
		return err
	}

	return nil
}

func (u *UseCase) SearchUsers(ctx context.Context) (*entity.UserStore, error) {
	data, err := u.repo.SearchUsers(ctx)
	if err != nil {
		l.Errorf(err)
		return nil, err
	}

	return data, nil
}
