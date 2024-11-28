package salary_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ISalaryRepository interface {
	CreateOne(ctx context.Context, salary *Salary) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, salary *Salary) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Salary, error)
	GetByStatus(ctx context.Context, status string) ([]Salary, error)
	GetByRoleID(ctx context.Context, roleID primitive.ObjectID) (Salary, error)
	GetAll(ctx context.Context) ([]Salary, error)
	CountSalary(ctx context.Context) (int64, error)
}

type ISalaryUseCase interface {
	CreateOne(ctx context.Context, salary *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, salary *Input) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	GetByID(ctx context.Context, id string) (Output, error)
	GetByRoleTitle(ctx context.Context, roleTitle string) (Output, error)
	GetByStatus(ctx context.Context, status string) ([]Output, error)
	GetAll(ctx context.Context) ([]Output, error)
	CountSalary(ctx context.Context) (int64, error)
}
