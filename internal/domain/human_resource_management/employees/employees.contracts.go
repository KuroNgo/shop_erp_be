package employees_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IEmployeeRepository interface {
	CreateOne(ctx context.Context, employee *Employee) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, employee *Employee) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, isActive bool) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Employee, error)
	GetByEmail(ctx context.Context, email string) (Employee, error)
	GetAll(ctx context.Context) ([]Employee, error)
	CountEmployeeByEmail(ctx context.Context, email string) (int64, error)
	CountEmployee(ctx context.Context) (int64, error)
	CountEmployeeByDepartmentID(ctx context.Context, departmentID primitive.ObjectID) (int64, error)
}

type IEmployeeUseCase interface {
	CreateOne(ctx context.Context, employee *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, employee *Input) error
	UpdateStatus(ctx context.Context, id string, isActive bool) error
	GetByID(ctx context.Context, id string) (Output, error)
	GetByEmail(ctx context.Context, email string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
	CountEmployee(ctx context.Context) (int64, error)
}
