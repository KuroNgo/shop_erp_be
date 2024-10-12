package departments_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDepartmentRepository interface {
	CreateOne(ctx context.Context, department *Department) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, department *Department) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Department, error)
	GetByName(ctx context.Context, name string) (Department, error)
	GetAll(ctx context.Context) ([]Department, error)
	CountManagerExist(ctx context.Context, managerID primitive.ObjectID) (int64, error)
	CountDepartment(ctx context.Context) (int64, error)
}

type IDepartmentUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetByID(ctx context.Context, id string) (Output, error)
	GetByName(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
	CountManagerExist(ctx context.Context, managerID primitive.ObjectID) (int64, error)
	CountDepartment(ctx context.Context) (int64, error)
}
