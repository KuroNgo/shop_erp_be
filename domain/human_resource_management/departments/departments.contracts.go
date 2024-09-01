package departments_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDepartmentRepository interface {
	CreateOne(ctx context.Context, department *Department) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, department *Department) error
	GetOneByID(ctx context.Context, id primitive.ObjectID) (Department, error)
	GetOneByName(ctx context.Context, name string) (Department, error)
	GetAll(ctx context.Context) ([]Department, error)
}

type IDepartmentUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByName(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
