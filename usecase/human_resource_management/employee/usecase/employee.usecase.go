package employee_usecase

import (
	"context"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"time"
)

type employeeUseCase struct {
	employeeRepository employeesdomain.IEmployeeRepository
	contextTimeout     time.Duration
}

func NewEmployeeUseCase(contextTimout time.Duration, employeeRepository employeesdomain.IEmployeeRepository) employeesdomain.IEmployeeUseCase {
	return &employeeUseCase{contextTimeout: contextTimout, employeeRepository: employeeRepository}
}

func (e employeeUseCase) CreateOne(ctx context.Context, employee *employeesdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	err := e.employeeRepository.CreateOne(ctx, employee)
	if err != nil {
		return err
	}

	return nil
}

func (e employeeUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	err := e.employeeRepository.DeleteOne(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (e employeeUseCase) UpdateOne(ctx context.Context, id string, employee *employeesdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	err := e.employeeRepository.UpdateOne(ctx, id, employee)
	if err != nil {
		return err
	}

	return nil
}

func (e employeeUseCase) GetOneByID(ctx context.Context, id string) (employeesdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	data, err := e.employeeRepository.GetOneByID(ctx, id)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	return data, nil
}

func (e employeeUseCase) GetOneByName(ctx context.Context, name string) (employeesdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	data, err := e.employeeRepository.GetOneByName(ctx, name)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	return data, nil
}

func (e employeeUseCase) GetAll(ctx context.Context) ([]employeesdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	data, err := e.employeeRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
