package role

import (
	"context"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"time"
)

type roleUseCase struct {
	contextTimeout time.Duration
	roleRepository roledomain.IRoleRepository
}

func NewRoleUseCase(contextTimeout time.Duration, roleRepository roledomain.IRoleRepository) roledomain.IRoleUseCase {
	return &roleUseCase{contextTimeout: contextTimeout, roleRepository: roleRepository}
}

func (r roleUseCase) CreateOneRole(ctx context.Context, role *roledomain.Role) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	err := r.roleRepository.CreateOneRole(ctx, role)
	if err != nil {
		return err
	}

	return nil
}

func (r roleUseCase) GetByTitleRole(ctx context.Context, title string) (roledomain.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.roleRepository.GetByTitleRole(ctx, title)
	if err != nil {
		return roledomain.Role{}, err
	}

	return data, nil
}

func (r roleUseCase) GetByIDRole(ctx context.Context, id string) (roledomain.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.roleRepository.GetByIDRole(ctx, id)
	if err != nil {
		return roledomain.Role{}, err
	}

	return data, nil
}

func (r roleUseCase) GetAllRole(ctx context.Context) ([]roledomain.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.roleRepository.GetAllRole(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r roleUseCase) UpdateOneRole(ctx context.Context, role *roledomain.Role) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	err := r.roleRepository.UpdateOneRole(ctx, role)
	if err != nil {
		return err
	}

	return nil
}

func (r roleUseCase) DeleteOneRole(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	err := r.roleRepository.DeleteOneRole(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
