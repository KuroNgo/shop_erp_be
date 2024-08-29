package role_domain

import "context"

type IRoleRepository interface {
	CreateOneRole(ctx context.Context, role *Input) error
	GetByTitleRole(ctx context.Context, title string) (Output, error)
	GetByIDRole(ctx context.Context, id string) (Output, error)
	GetAllRole(ctx context.Context) ([]Output, error)
	UpdateOneRole(ctx context.Context, id string, role *Input) error
	DeleteOneRole(ctx context.Context, id string) error
}

type IRoleUseCase interface {
	CreateOneRole(ctx context.Context, role *Input) error
	GetByTitleRole(ctx context.Context, title string) (Output, error)
	GetByIDRole(ctx context.Context, id string) (Output, error)
	GetAllRole(ctx context.Context) ([]Output, error)
	UpdateOneRole(ctx context.Context, id string, role *Input) error
	DeleteOneRole(ctx context.Context, id string) error
}
