package base_salary_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ISalaryRepository interface {
	CreateOne(ctx context.Context, baseSalary *BaseSalary) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, baseSalary *BaseSalary) error
	GetByID(ctx context.Context, id primitive.ObjectID) (BaseSalary, error)
	GetByRoleID(ctx context.Context, roleID primitive.ObjectID) (BaseSalary, error)
	GetAll(ctx context.Context) ([]BaseSalary, error)
}

type ISalaryUseCase interface {
	CreateOne(ctx context.Context, salary *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, salary *Input) error
	GetByID(ctx context.Context, id string) (BaseSalary, error)
	GetByRoleID(ctx context.Context, roleID string) (BaseSalary, error)
	GetAll(ctx context.Context) ([]BaseSalary, error)
}
