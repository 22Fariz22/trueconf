package usecase

import (
	"context"
	"refactoring/internal/user"
)

type useCaseUser struct{
	repo user.RepositoryUser
}

func NewUseCaseUser(repo user.RepositoryUser)*useCaseUser{
	return &useCaseUser{repo: repo}
}

func(u *useCaseUser)CreateUser(ctx context.Context){}
func(u *useCaseUser)DeleteUser(ctx context.Context){}
func(u *useCaseUser)GetUser(ctx context.Context){}
func(u *useCaseUser)UpdateUser(ctx context.Context){}
func(u *useCaseUser)SearchUser(ctx context.Context){}