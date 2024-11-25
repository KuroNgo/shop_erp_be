package role_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IRoleRepository interface {
	CreateOne(ctx context.Context, role *Role) error
	GetByName(ctx context.Context, name string) (Role, error)
	GetByLevel(ctx context.Context, level int) ([]Role, error)
	GetByEnable(ctx context.Context, enable int) ([]Role, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (Role, error)
	GetAll(ctx context.Context) ([]Role, error)
	UpdateOne(ctx context.Context, id primitive.ObjectID, role *Role) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	DeleteSoft(ctx context.Context, id primitive.ObjectID) error
	CountRole(ctx context.Context) (int64, error)
}

type IRoleUseCase interface {
	CreateOne(ctx context.Context, input *Input, idUser string) error
	UpdateOne(ctx context.Context, id string, input *Input, idUser string) error
	DeleteOne(ctx context.Context, id string, idUser string) error
	DeleteSoft(ctx context.Context, id string, idUser string) error
	CountRole(ctx context.Context) (int64, error)
	GetByEnable(ctx context.Context, enable int) ([]Output, error)
	GetByLevel(ctx context.Context, level int) ([]Output, error)
	GetByName(ctx context.Context, name string) (Output, error)
	GetByID(ctx context.Context, id string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
	Lifecycle(ctx context.Context) error
}
