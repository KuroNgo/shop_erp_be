package role_domain

import "context"

type IRoleRepository interface {
	CreateOneRole(ctx context.Context, role *Role) error
	GetByTitleRole(ctx context.Context, title string) (Role, error)
	GetByIDRole(ctx context.Context, id string) (Role, error)
	GetAllRole(ctx context.Context) ([]Role, error)
	UpdateOneRole(ctx context.Context, role *Role) error
	DeleteOneRole(ctx context.Context, id string) error
}

type IRoleUseCase interface {
	CreateOneRole(ctx context.Context, role *Role) error
	GetByTitleRole(ctx context.Context, title string) (Role, error)
	GetByIDRole(ctx context.Context, id string) (Role, error)
	GetAllRole(ctx context.Context) ([]Role, error)
	UpdateOneRole(ctx context.Context, role *Role) error
	DeleteOneRole(ctx context.Context, id string) error
}
