package employees_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IEmployeeRepository interface {
	CreateOne(ctx context.Context, employee *Employee) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	DeleteSoft(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, employee *Employee) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, active string) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Employee, error)
	GetByName(ctx context.Context, name string) (Employee, error)
	GetByEmail(ctx context.Context, email string) (Employee, error)
	GetByStatus(ctx context.Context, status string) ([]Employee, error)
	GetAll(ctx context.Context) ([]Employee, error)
	CountEmployeeByEmail(ctx context.Context, email string) (int64, error)
	CountEmployee(ctx context.Context) (int64, error)
	CountEmployeeByDepartmentID(ctx context.Context, departmentID primitive.ObjectID) (int64, error)
	CountEmployeeByDepartmentID2(ctx context.Context, departmentID primitive.ObjectID) (int64, error)
	CountEmployeeByRoleID(ctx context.Context, roleID primitive.ObjectID) (int64, error)
}

type IEmployeeUseCase interface {
	CreateOne(ctx context.Context, employee *Input) error
	DeleteOne(ctx context.Context, id string) error
	DeleteSoft(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, employee *Input) error
	UpdateStatus(ctx context.Context, id string, active string) error
	GetByID(ctx context.Context, id string) (Output, error)
	GetByName(ctx context.Context, name string) (Output, error)
	GetByEmail(ctx context.Context, email string) (Output, error)
	GetByStatus(ctx context.Context, status string) ([]Output, error)
	GetAll(ctx context.Context) ([]Output, error)
	CountEmployee(ctx context.Context) (int64, error)
	LifeCycle(ctx context.Context) error
}
