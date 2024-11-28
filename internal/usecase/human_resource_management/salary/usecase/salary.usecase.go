package salary_usecase

import (
	"context"
	"errors"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	employees_domain "shop_erp_mono/internal/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/internal/domain/human_resource_management/salary"
	"shop_erp_mono/internal/usecase/human_resource_management/salary/validate"
	"time"
)

type salaryUseCase struct {
	contextTimeout     time.Duration
	salaryRepository   salarydomain.ISalaryRepository
	roleRepository     roledomain.IRoleRepository
	employeeRepository employees_domain.IEmployeeRepository
	cache              *bigcache.BigCache
}

func NewSalaryUseCase(contextTimout time.Duration, salaryRepository salarydomain.ISalaryRepository,
	roleRepository roledomain.IRoleRepository, employeeRepository employees_domain.IEmployeeRepository, cacheTTL time.Duration) salarydomain.ISalaryUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &salaryUseCase{contextTimeout: contextTimout, cache: cache, salaryRepository: salaryRepository,
		roleRepository: roleRepository, employeeRepository: employeeRepository}
}

func (s *salaryUseCase) CreateOne(ctx context.Context, input *salarydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.Salary(input); err != nil {
		return err
	}

	// Calculate net salary
	netSalary := input.BaseSalary + input.Bonus - input.Deductions

	salaryData := &salarydomain.Salary{
		ID:           primitive.NewObjectID(),
		EmployeeID:   input.EmployeeID,
		UnitCurrency: input.UnitCurrency,
		BaseSalary:   input.BaseSalary,
		Bonus:        input.Bonus,
		Deductions:   input.Deductions,
		NetSalary:    netSalary,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := s.salaryRepository.CreateOne(ctx, salaryData)
	if err != nil {
		return err
	}

	return nil
}

func (s *salaryUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salaryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = s.salaryRepository.DeleteOne(ctx, salaryID)
	if err != nil {
		return err
	}

	return nil
}

func (s *salaryUseCase) UpdateOne(ctx context.Context, id string, input *salarydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salaryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	salary, err := s.salaryRepository.GetByID(ctx, salaryID)
	if err != nil {
		return err
	}

	if salary.Status == "paid" {
		return errors.New("the salary has been paid and is therefore not eligible for updates")
	}

	if err := validate.Salary(input); err != nil {
		return err
	}

	// Calculate net salary
	netSalary := input.BaseSalary + input.Bonus - input.Deductions

	salaryData := &salarydomain.Salary{
		EmployeeID:   input.EmployeeID,
		UnitCurrency: input.UnitCurrency,
		BaseSalary:   input.BaseSalary,
		Bonus:        input.Bonus,
		Deductions:   input.Deductions,
		NetSalary:    netSalary,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = s.salaryRepository.UpdateOne(ctx, salaryID, salaryData)
	if err != nil {
		return err
	}

	return nil
}

func (s *salaryUseCase) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salary, err := s.salaryRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if salary.Status == "paid" {
		return errors.New("the salary has been paid and is therefore not eligible for updates")
	}

	salaryData := &salarydomain.Salary{
		Status:    status,
		UpdatedAt: time.Now(),
	}

	err = s.salaryRepository.UpdateOne(ctx, id, salaryData)
	if err != nil {
		return err
	}

	return nil
}

func (s *salaryUseCase) GetByID(ctx context.Context, id string) (salarydomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salaryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return salarydomain.Output{}, err
	}

	salaryData, err := s.salaryRepository.GetByID(ctx, salaryID)
	if err != nil {
		return salarydomain.Output{}, err
	}

	output := salarydomain.Output{
		Salary: salaryData,
	}

	return output, nil
}

func (s *salaryUseCase) GetByStatus(ctx context.Context, status string) ([]salarydomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salaryData, err := s.salaryRepository.GetByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var outputs []salarydomain.Output
	outputs = make([]salarydomain.Output, 0, len(salaryData))
	for _, salary := range salaryData {
		output := salarydomain.Output{
			Salary: salary,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}

func (s *salaryUseCase) GetByRoleTitle(ctx context.Context, roleTitle string) (salarydomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	roleData, err := s.roleRepository.GetByName(ctx, roleTitle)
	if err != nil {
		return salarydomain.Output{}, err
	}

	salaryData, err := s.salaryRepository.GetByRoleID(ctx, roleData.ID)
	if err != nil {
		return salarydomain.Output{}, err
	}

	output := salarydomain.Output{
		Salary: salaryData,
	}

	return output, nil
}

func (s *salaryUseCase) GetAll(ctx context.Context) ([]salarydomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salaryData, err := s.salaryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []salarydomain.Output
	outputs = make([]salarydomain.Output, 0, len(salaryData))
	for _, salary := range salaryData {
		output := salarydomain.Output{
			Salary: salary,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}

func (s *salaryUseCase) CountSalary(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	return s.salaryRepository.CountSalary(ctx)
}
