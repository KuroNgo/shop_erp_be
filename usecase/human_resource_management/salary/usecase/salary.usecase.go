package salary_usecase

import (
	"context"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"time"
)

type salaryUseCase struct {
	contextTimeout   time.Duration
	salaryRepository salarydomain.ISalaryRepository
}

func NewSalaryUseCase(contextTimout time.Duration, salaryRepository salarydomain.ISalaryRepository) salarydomain.ISalaryUseCase {
	return &salaryUseCase{contextTimeout: contextTimout, salaryRepository: salaryRepository}
}

func (s salaryUseCase) CreateOne(ctx context.Context, salary *salarydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	err := s.salaryRepository.CreateOne(ctx, salary)
	if err != nil {
		return err
	}

	return nil
}

func (s salaryUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	err := s.salaryRepository.DeleteOne(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s salaryUseCase) UpdateOne(ctx context.Context, salary *salarydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	err := s.salaryRepository.UpdateOne(ctx, salary)
	if err != nil {
		return err
	}

	return nil
}

func (s salaryUseCase) GetOneByID(ctx context.Context, id string) (salarydomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	data, err := s.salaryRepository.GetOneByID(ctx, id)
	if err != nil {
		return salarydomain.Output{}, err
	}

	return data, nil
}

func (s salaryUseCase) GetOneByRole(ctx context.Context, role string) (salarydomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	data, err := s.salaryRepository.GetOneByRole(ctx, role)
	if err != nil {
		return salarydomain.Output{}, err
	}

	return data, nil
}

func (s salaryUseCase) GetAll(ctx context.Context) ([]salarydomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	data, err := s.salaryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
