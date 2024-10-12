package role_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IRoleRepository interface {
	CreateOne(ctx context.Context, role *Role) error
	GetByTitle(ctx context.Context, title string) (Role, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (Role, error)
	GetAll(ctx context.Context) ([]Role, error)
	UpdateOne(ctx context.Context, id primitive.ObjectID, role *Role) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	CountRole(ctx context.Context) (int64, error)
}

type IRoleUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByTitle(ctx context.Context, title string) (Output, error)
	GetByID(ctx context.Context, id string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	CountRole(ctx context.Context) (int64, error)
}
