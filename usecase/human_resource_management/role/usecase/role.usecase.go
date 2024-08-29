package role_usecase

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

func (r roleUseCase) CreateOneRole(ctx context.Context, input *roledomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	err := r.roleRepository.CreateOneRole(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (r roleUseCase) GetByTitleRole(ctx context.Context, title string) (roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.roleRepository.GetByTitleRole(ctx, title)
	if err != nil {
		return roledomain.Output{}, err
	}

	return data, nil
}

func (r roleUseCase) GetByIDRole(ctx context.Context, id string) (roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.roleRepository.GetByIDRole(ctx, id)
	if err != nil {
		return roledomain.Output{}, err
	}

	return data, nil
}

func (r roleUseCase) GetAllRole(ctx context.Context) ([]roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.roleRepository.GetAllRole(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r roleUseCase) UpdateOneRole(ctx context.Context, id string, input *roledomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	err := r.roleRepository.UpdateOneRole(ctx, id, input)
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
