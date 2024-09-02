package role_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IRoleRepository interface {
	CreateOneRole(ctx context.Context, role *Role) error
	GetByTitleRole(ctx context.Context, title string) (Role, error)
	GetByIDRole(ctx context.Context, id primitive.ObjectID) (Role, error)
	GetAllRole(ctx context.Context) ([]Role, error)
	UpdateOneRole(ctx context.Context, id primitive.ObjectID, role *Role) error
	DeleteOneRole(ctx context.Context, id primitive.ObjectID) error
}

type IRoleUseCase interface {
	CreateOneRole(ctx context.Context, input *Input) error
	GetByTitleRole(ctx context.Context, title string) (Output, error)
	GetByIDRole(ctx context.Context, id string) (Output, error)
	GetAllRole(ctx context.Context) ([]Output, error)
	UpdateOneRole(ctx context.Context, id string, input *Input) error
	DeleteOneRole(ctx context.Context, id string) error
}
