package employee_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"shop_erp_mono/repository/human_resource_management/employee/validate"
	"time"
)

type employeeUseCase struct {
	contextTimeout       time.Duration
	employeeRepository   employeesdomain.IEmployeeRepository
	departmentRepository departmentsdomain.IDepartmentRepository
	salaryRepository     salarydomain.ISalaryRepository
	roleRepository       roledomain.IRoleRepository
}

func NewEmployeeUseCase(contextTimout time.Duration, employeeRepository employeesdomain.IEmployeeRepository,
	departmentRepository departmentsdomain.IDepartmentRepository, salaryRepository salarydomain.ISalaryRepository,
	roleRepository roledomain.IRoleRepository) employeesdomain.IEmployeeUseCase {
	return &employeeUseCase{contextTimeout: contextTimout, employeeRepository: employeeRepository,
		departmentRepository: departmentRepository, salaryRepository: salaryRepository, roleRepository: roleRepository}
}

func (e employeeUseCase) CreateOne(ctx context.Context, input *employeesdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	if err := validate.IsNilEmployee(input); err != nil {
		return err
	}

	departmentData, err := e.departmentRepository.GetOneByName(ctx, input.Department)
	if err != nil {
		return err
	}

	roleData, err := e.roleRepository.GetByTitleRole(ctx, input.Role)
	if err != nil {
		return err
	}

	salaryData, err := e.salaryRepository.GetOneByRoleID(ctx, roleData.ID)
	if err != nil {
		return err
	}

	employeeData := &employeesdomain.Employee{
		ID:           primitive.NewObjectID(),
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Gender:       input.Gender,
		Email:        input.Email,
		Phone:        input.Phone,
		Address:      input.Address,
		AvatarURL:    input.AvatarURL,
		DateOfBirth:  input.DateOfBirth,
		DayOfWork:    input.DayOfWork,
		DepartmentID: departmentData.ID,
		SalaryID:     salaryData.ID,
		RoleID:       roleData.ID,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return e.employeeRepository.CreateOne(ctx, employeeData)
}

func (e employeeUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return e.employeeRepository.DeleteOne(ctx, employeeID)
}

func (e employeeUseCase) UpdateOne(ctx context.Context, id string, input *employeesdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	departmentData, err := e.departmentRepository.GetOneByName(ctx, input.Department)
	if err != nil {
		return err
	}

	roleData, err := e.roleRepository.GetByTitleRole(ctx, input.Role)
	if err != nil {
		return err
	}

	salaryData, err := e.salaryRepository.GetOneByRoleID(ctx, roleData.ID)
	if err != nil {
		return err
	}

	employee := &employeesdomain.Employee{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Gender:       input.Gender,
		Email:        input.Email,
		Phone:        input.Phone,
		Address:      input.Address,
		AvatarURL:    input.AvatarURL,
		DateOfBirth:  input.DateOfBirth,
		DayOfWork:    input.DayOfWork,
		DepartmentID: departmentData.ID,
		RoleID:       roleData.ID,
		SalaryID:     salaryData.ID,
		UpdatedAt:    time.Now(),
	}

	return e.employeeRepository.UpdateOne(ctx, employeeID, employee)
}

func (e employeeUseCase) GetOneByID(ctx context.Context, id string) (employeesdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	employeeData, err := e.employeeRepository.GetOneByID(ctx, employeeID)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	output := employeesdomain.Output{
		Employee: employeeData,
	}
	return output, nil
}

func (e employeeUseCase) GetOneByEmail(ctx context.Context, name string) (employeesdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	employeeData, err := e.employeeRepository.GetOneByEmail(ctx, name)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	output := employeesdomain.Output{
		Employee: employeeData,
	}
	return output, nil
}

func (e employeeUseCase) GetAll(ctx context.Context) ([]employeesdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	employeeData, err := e.employeeRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []employeesdomain.Output
	outputs = make([]employeesdomain.Output, 0, len(employeeData))
	for _, employee := range employeeData {
		output := employeesdomain.Output{
			Employee: employee,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}
