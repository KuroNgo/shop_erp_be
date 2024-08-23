package departments_domain

import "context"

type IDepartmentRepository interface {
	CreateOne(ctx context.Context, department *Department) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, department *Department) error
	GetOneByID(ctx context.Context, id string) (Department, error)
	GetOneByName(ctx context.Context, name string) (Department, error)
	GetAll(ctx context.Context) ([]Department, error)
}

type IDepartmentUseCase interface {
	CreateOne(ctx context.Context, department *Department) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, department *Department) error
	GetOneByID(ctx context.Context, id string) (Department, error)
	GetOneByName(ctx context.Context, name string) (Department, error)
	GetAll(ctx context.Context) ([]Department, error)
}
