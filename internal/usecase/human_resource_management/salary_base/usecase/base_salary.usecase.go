package base_salary_usecase

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	basesalarydomain "shop_erp_mono/internal/domain/human_resource_management/salary_base"
	"shop_erp_mono/internal/usecase/human_resource_management/salary_base/validate"
	"time"
)

type baseSalaryUseCase struct {
	contextTimeout       time.Duration
	baseSalaryRepository basesalarydomain.ISalaryRepository
	roleRepository       roledomain.IRoleRepository
	cache                *bigcache.BigCache
}

func NewBaseSalaryUseCase(contextTimout time.Duration, baseSalaryRepository basesalarydomain.ISalaryRepository,
	roleRepository roledomain.IRoleRepository, cacheTTL time.Duration) basesalarydomain.ISalaryUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &baseSalaryUseCase{contextTimeout: contextTimout, cache: cache, baseSalaryRepository: baseSalaryRepository, roleRepository: roleRepository}
}

func (b *baseSalaryUseCase) CreateOne(ctx context.Context, salary *basesalarydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	if err := validate.BaseSalary(salary); err != nil {
		return err
	}

	baseSalary := &basesalarydomain.BaseSalary{
		ID:           primitive.NewObjectID(),
		RoleID:       salary.RoleID,
		UnitCurrency: salary.UnitCurrency,
		BaseSalaries: salary.BaseSalary,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := b.baseSalaryRepository.CreateOne(ctx, baseSalary)
	if err != nil {
		return err
	}

	return nil
}

func (b *baseSalaryUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	baseID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = b.baseSalaryRepository.DeleteOne(ctx, baseID)
	if err != nil {
		return err
	}

	return nil
}

func (b *baseSalaryUseCase) UpdateOne(ctx context.Context, id string, salary *basesalarydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	baseID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err := validate.BaseSalary(salary); err != nil {
		return err
	}

	baseSalary := &basesalarydomain.BaseSalary{
		ID:           primitive.NewObjectID(),
		RoleID:       salary.RoleID,
		UnitCurrency: salary.UnitCurrency,
		BaseSalaries: salary.BaseSalary,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = b.baseSalaryRepository.UpdateOne(ctx, baseID, baseSalary)
	if err != nil {
		return err
	}

	return nil
}

func (b *baseSalaryUseCase) GetByID(ctx context.Context, id string) (basesalarydomain.BaseSalary, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	baseID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return basesalarydomain.BaseSalary{}, err
	}

	data, err := b.baseSalaryRepository.GetByID(ctx, baseID)
	if err != nil {
		return basesalarydomain.BaseSalary{}, err
	}

	return data, nil
}

func (b *baseSalaryUseCase) GetByRoleID(ctx context.Context, roleID string) (basesalarydomain.BaseSalary, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(roleID)
	if err != nil {
		return basesalarydomain.BaseSalary{}, err
	}

	data, err := b.baseSalaryRepository.GetByRoleID(ctx, id)
	if err != nil {
		return basesalarydomain.BaseSalary{}, err
	}

	return data, nil
}

func (b *baseSalaryUseCase) GetAll(ctx context.Context) ([]basesalarydomain.BaseSalary, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	salaryData, err := b.baseSalaryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []basesalarydomain.BaseSalary
	outputs = make([]basesalarydomain.BaseSalary, 0, len(salaryData))
	for _, salary := range salaryData {
		outputs = append(outputs, salary)
	}

	return outputs, nil
}
