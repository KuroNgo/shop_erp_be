package salary_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"shop_erp_mono/repository/human_resource_management/salary/validate"
	"time"
)

type salaryUseCase struct {
	contextTimeout   time.Duration
	salaryRepository salarydomain.ISalaryRepository
	roleRepository   roledomain.IRoleRepository
}

func NewSalaryUseCase(contextTimout time.Duration, salaryRepository salarydomain.ISalaryRepository, roleRepository roledomain.IRoleRepository) salarydomain.ISalaryUseCase {
	return &salaryUseCase{contextTimeout: contextTimout, salaryRepository: salaryRepository, roleRepository: roleRepository}
}

func (s salaryUseCase) CreateOne(ctx context.Context, input *salarydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.IsNilSalary(input); err != nil {
		return err
	}

	roleData, err := s.roleRepository.GetByTitleRole(ctx, input.Role)
	if err != nil {
		return err
	}

	salaryData := &salarydomain.Salary{
		ID:           primitive.NewObjectID(),
		RoleID:       roleData.ID,
		UnitCurrency: input.UnitCurrency,
		BaseSalary:   input.BaseSalary,
		Bonus:        input.Bonus,
		Deductions:   input.Deductions,
		NetSalary:    input.NetSalary,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return s.salaryRepository.CreateOne(ctx, salaryData)
}

func (s salaryUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salaryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.salaryRepository.DeleteOne(ctx, salaryID)
}

func (s salaryUseCase) UpdateOne(ctx context.Context, id string, input *salarydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.IsNilSalary(input); err != nil {
		return err
	}

	salaryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	roleData, err := s.roleRepository.GetByTitleRole(ctx, input.Role)
	if err != nil {
		return err
	}

	salaryData := &salarydomain.Salary{
		RoleID:       roleData.ID,
		UnitCurrency: input.UnitCurrency,
		BaseSalary:   input.BaseSalary,
		Bonus:        input.Bonus,
		Deductions:   input.Deductions,
		NetSalary:    input.NetSalary,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return s.salaryRepository.UpdateOne(ctx, salaryID, salaryData)
}

func (s salaryUseCase) GetOneByID(ctx context.Context, id string) (salarydomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salaryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return salarydomain.Output{}, err
	}

	salaryData, err := s.salaryRepository.GetOneByID(ctx, salaryID)
	if err != nil {
		return salarydomain.Output{}, err
	}

	roleData, err := s.roleRepository.GetByIDRole(ctx, salaryData.RoleID)
	if err != nil {
		return salarydomain.Output{}, err
	}

	output := salarydomain.Output{
		Salary: salaryData,
		Role:   roleData,
	}

	return output, nil
}

func (s salaryUseCase) GetOneByRoleTitle(ctx context.Context, roleTitle string) (salarydomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	roleData, err := s.roleRepository.GetByTitleRole(ctx, roleTitle)
	if err != nil {
		return salarydomain.Output{}, err
	}

	salaryData, err := s.salaryRepository.GetOneByRoleID(ctx, roleData.ID)
	if err != nil {
		return salarydomain.Output{}, err
	}

	output := salarydomain.Output{
		Salary: salaryData,
		Role:   roleData,
	}

	return output, nil
}

func (s salaryUseCase) GetAll(ctx context.Context) ([]salarydomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salaryData, err := s.salaryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []salarydomain.Output
	outputs = make([]salarydomain.Output, 0, len(salaryData))
	for _, salary := range salaryData {
		roleData, err := s.roleRepository.GetByIDRole(ctx, salary.RoleID)
		if err != nil {
			return nil, err
		}
		output := salarydomain.Output{
			Salary: salary,
			Role:   roleData,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}
