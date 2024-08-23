package department_usecase

import (
	"context"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	"time"
)

type departmentUseCase struct {
	contextTimeout       time.Duration
	departmentRepository departmentsdomain.IDepartmentRepository
}

func NewDepartmentUseCase(contextTimeout time.Duration, departmentRepository departmentsdomain.IDepartmentRepository) departmentsdomain.IDepartmentUseCase {
	return &departmentUseCase{contextTimeout: contextTimeout, departmentRepository: departmentRepository}
}

func (d departmentUseCase) CreateOne(ctx context.Context, department *departmentsdomain.Department) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	err := d.departmentRepository.CreateOne(ctx, department)
	if err != nil {
		return err
	}

	return nil
}

func (d departmentUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	err := d.departmentRepository.DeleteOne(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (d departmentUseCase) UpdateOne(ctx context.Context, department *departmentsdomain.Department) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	err := d.departmentRepository.UpdateOne(ctx, department)
	if err != nil {
		return err
	}

	return nil
}

func (d departmentUseCase) GetOneByID(ctx context.Context, id string) (departmentsdomain.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	data, err := d.departmentRepository.GetOneByID(ctx, id)
	if err != nil {
		return departmentsdomain.Department{}, err
	}

	return data, nil
}

func (d departmentUseCase) GetOneByName(ctx context.Context, name string) (departmentsdomain.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	data, err := d.departmentRepository.GetOneByName(ctx, name)
	if err != nil {
		return departmentsdomain.Department{}, err
	}

	return data, nil
}

func (d departmentUseCase) GetAll(ctx context.Context) ([]departmentsdomain.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	data, err := d.departmentRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
