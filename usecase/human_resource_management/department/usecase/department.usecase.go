package department_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/repository/human_resource_management/department/validate"
	"time"
)

type departmentUseCase struct {
	contextTimeout       time.Duration
	departmentRepository departmentsdomain.IDepartmentRepository
	employeeRepository   employeesdomain.IEmployeeRepository
}

func NewDepartmentUseCase(contextTimeout time.Duration, departmentRepository departmentsdomain.IDepartmentRepository,
	employeeRepository employeesdomain.IEmployeeRepository) departmentsdomain.IDepartmentUseCase {
	return &departmentUseCase{contextTimeout: contextTimeout, departmentRepository: departmentRepository, employeeRepository: employeeRepository}
}

func (d *departmentUseCase) CreateOne(ctx context.Context, input *departmentsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	if err := validate.ValidateDepartment(input); err != nil {
		return err
	}

	managerData, err := d.employeeRepository.GetOneByEmail(ctx, input.ManagerEmail)
	if err != nil {
		return err
	}

	department := &departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		ManagerID:   managerData.ID,
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return d.departmentRepository.CreateOne(ctx, department)
}

func (d *departmentUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	departmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return d.departmentRepository.DeleteOne(ctx, departmentID)
}

func (d *departmentUseCase) UpdateOne(ctx context.Context, id string, input *departmentsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	if err := validate.ValidateDepartment(input); err != nil {
		return err
	}

	departmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	department := &departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		Name:        input.Name,
		Description: input.Description,
		UpdatedAt:   time.Now(),
	}

	return d.departmentRepository.UpdateOne(ctx, departmentID, department)
}

func (d *departmentUseCase) GetOneByID(ctx context.Context, id string) (departmentsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	departmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	departmentData, err := d.departmentRepository.GetOneByID(ctx, departmentID)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	output := departmentsdomain.Output{
		Department: departmentData,
	}

	return output, nil
}

func (d *departmentUseCase) GetOneByName(ctx context.Context, name string) (departmentsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	departmentData, err := d.departmentRepository.GetOneByName(ctx, name)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	output := departmentsdomain.Output{
		Department: departmentData,
	}
	return output, nil
}

func (d *departmentUseCase) GetAll(ctx context.Context) ([]departmentsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	departmentsData, err := d.departmentRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []departmentsdomain.Output
	outputs = make([]departmentsdomain.Output, 0, len(departmentsData))
	for _, departmentData := range departmentsData {
		managerData, err := d.employeeRepository.GetOneByID(ctx, departmentData.ManagerID)
		if err != nil {
			return nil, err
		}

		output := departmentsdomain.Output{
			Department: departmentData,
			Manager:    managerData,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}
