package user

import "context"

type RepositoryUser interface{
	CreateUser(ctx context.Context)
	DeleteUser(ctx context.Context)
	GetUser(ctx context.Context)
	UpdateUser(ctx context.Context)
	SearchUser(ctx context.Context)
}
