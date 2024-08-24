package employees_domain

import "context"

type IEmployeeRepository interface {
	CreateOne(ctx context.Context, employee *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, employee *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByName(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}

type IEmployeeUseCase interface {
	CreateOne(ctx context.Context, employee *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, employee *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByName(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
