package employees_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IEmployeeRepository interface {
	CreateOne(ctx context.Context, employee *Employee) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, employee *Employee) error
	GetOneByID(ctx context.Context, id primitive.ObjectID) (Employee, error)
	GetOneByEmail(ctx context.Context, email string) (Employee, error)
	GetAll(ctx context.Context) ([]Employee, error)
}

type IEmployeeUseCase interface {
	CreateOne(ctx context.Context, employee *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, employee *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByEmail(ctx context.Context, email string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
