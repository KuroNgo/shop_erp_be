package salary_domain

import "context"

type ISalaryRepository interface {
	CreateOne(ctx context.Context, salary *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, salary *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByRole(ctx context.Context, role string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}

type ISalaryUseCase interface {
	CreateOne(ctx context.Context, salary *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, salary *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByRole(ctx context.Context, role string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
