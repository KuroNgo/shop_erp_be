package departments_domain

import "context"

type IDepartmentRepository interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByName(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}

type IDepartmentUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByName(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
