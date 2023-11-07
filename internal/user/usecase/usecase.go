package usecase

import (
	"context"
	"time"

	"github.com/22Fariz22/trueconf/internal/user"
	"github.com/22Fariz22/trueconf/internal/user/entity"
	"github.com/22Fariz22/trueconf/pkg/logger"
)

type UseCase struct {
	Repository user.Repository
}

func NewUseCaseUser(repository user.Repository) *UseCase {
	return &UseCase{Repository: repository}
}

func (u *UseCase) CreateUser(ctx context.Context, l logger.Logger, newU *entity.User) error {
	newU.CreatedAt = time.Now()

	err := u.Repository.CreateUser(ctx, l, newU)
	if err != nil {
		l.Errorf("err %w in usecase CreateUser()->u.repository.CreateUser() with newU.DispalyName:%s", err, newU.DisplayName)
		return err
	}

	return nil
}

func (u *UseCase) DeleteUser(ctx context.Context, l logger.Logger, id string) error {
	err := u.Repository.DeleteUser(ctx, l, id)
	if err != nil {
		l.Errorf("err %w in usecase DeleteUser()->u.repository.DeleteUser() with id:%s", err, id)
		return err
	}

	return nil
}

func (u *UseCase) GetUser(ctx context.Context, l logger.Logger, id string) (*entity.UserStore, error) {
	data, err := u.Repository.GetUser(ctx, l, id)
	if err != nil {
		l.Errorf("err %w in usecase GetUser()->u.repository.GetUser() with id:%s", err, id)
		return &entity.UserStore{}, err
	}

	return data, nil
}

func (u *UseCase) UpdateUser(ctx context.Context, l logger.Logger, id string, updUser *entity.User) error {
	err := u.Repository.UpdateUser(ctx, l, id, updUser)
	if err != nil {
		l.Errorf("err %w in usecase UpdateUser()->u.repository.UpdateUser() with id:%s", err, id)
		return err
	}

	return nil
}

func (u *UseCase) SearchUsers(ctx context.Context, l logger.Logger) (*entity.UserStore, error) {
	data, err := u.Repository.SearchUsers(ctx, l)
	if err != nil {
		l.Errorf("err in usecase SearchUsers()->u.repository.SearchUsers()", err)
		return nil, err
	}

	return data, nil
}
