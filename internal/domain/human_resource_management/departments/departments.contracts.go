package departments_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	employees_domain "shop_erp_mono/internal/domain/human_resource_management/employees"
)

type IDepartmentRepository interface {
	CreateOne(ctx context.Context, department *Department) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	DeleteSoftOne(ctx context.Context, id primitive.ObjectID, whoDeleted primitive.ObjectID) error

	UpdateOne(ctx context.Context, id primitive.ObjectID, department *Department) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	UpdateManager(ctx context.Context, id primitive.ObjectID, managerID primitive.ObjectID) error

	GetByID(ctx context.Context, id primitive.ObjectID) (Department, error)
	GetByStatus(ctx context.Context, status string) ([]Department, error)
	GetByName(ctx context.Context, name string) (Department, error)
	GetAll(ctx context.Context) ([]Department, error)
	GetAllSoftDelete(ctx context.Context) ([]Department, error)
	GetAllDepartmentAlmostExpire(ctx context.Context) ([]Department, error)
	
	CountManagerExist(ctx context.Context, managerID primitive.ObjectID) (int64, error)
	CountDepartment(ctx context.Context) (int64, error)
	CountDepartmentWithName(ctx context.Context, name string) (int64, error)
}

type IDepartmentUseCase interface {
	CreateOne(ctx context.Context, input *Input, idUser string) error
	CreateDepartmentWithManager(ctx context.Context, departmentInput *Input, employeeInput *employees_domain.Input, idUser string) error

	DeleteOne(ctx context.Context, id string, userID string) error
	DeleteSoftOne(ctx context.Context, id string, userID string) error

	UpdateOne(ctx context.Context, id string, input *Input, idUser string) error
	UpdateStatus(ctx context.Context, id string, status string, idUser string) error
	UpdateManager(ctx context.Context, id string, managerID string, idUser string) error

	GetByID(ctx context.Context, id string) (Output, error)
	GetByName(ctx context.Context, name string) (Output, error)
	GetByStatus(ctx context.Context, status string) ([]Output, error)
	GetAll(ctx context.Context, idUser string) ([]Output, error)
	GetAllSoftDelete(ctx context.Context) ([]Output, error)

	CountManagerExist(ctx context.Context, managerID primitive.ObjectID) (int64, error)
	CountDepartment(ctx context.Context) (int64, error)
	LifecycleDepartment(ctx context.Context) error
}
