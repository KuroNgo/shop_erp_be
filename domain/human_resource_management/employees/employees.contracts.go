package employees_domain

import "context"

type IEmployeeRepository interface {
	CreateOne(ctx context.Context, employee *Employee) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, employee *Employee) error
	GetOneByID(ctx context.Context, id string) (Employee, error)
	GetOneByName(ctx context.Context, name string) (Employee, error)
	GetAll(ctx context.Context) ([]Employee, error)
}

type IEmployeeUseCase interface {
	CreateOne(ctx context.Context, employee *Employee) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, employee *Employee) error
	GetOneByID(ctx context.Context, id string) (Employee, error)
	GetOneByName(ctx context.Context, name string) (Employee, error)
	GetAll(ctx context.Context) ([]Employee, error)
}
