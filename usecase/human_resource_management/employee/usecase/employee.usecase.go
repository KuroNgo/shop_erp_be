package employee_usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"shop_erp_mono/usecase/human_resource_management/employee/validate"
	"time"
)

type employeeUseCase struct {
	contextTimeout       time.Duration
	employeeRepository   employeesdomain.IEmployeeRepository
	departmentRepository departmentsdomain.IDepartmentRepository
	salaryRepository     salarydomain.ISalaryRepository
	roleRepository       roledomain.IRoleRepository
	cache                *bigcache.BigCache
}

func NewEmployeeUseCase(contextTimout time.Duration, employeeRepository employeesdomain.IEmployeeRepository,
	departmentRepository departmentsdomain.IDepartmentRepository, salaryRepository salarydomain.ISalaryRepository,
	roleRepository roledomain.IRoleRepository, cacheTTL time.Duration) employeesdomain.IEmployeeUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &employeeUseCase{contextTimeout: contextTimout, employeeRepository: employeeRepository,
		departmentRepository: departmentRepository, cache: cache, salaryRepository: salaryRepository, roleRepository: roleRepository}
}

func (e *employeeUseCase) CreateOne(ctx context.Context, input *employeesdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	if err := validate.Employee(input); err != nil {
		return err
	}

	departmentData, err := e.departmentRepository.GetByName(ctx, input.Department)
	if err != nil {
		return err
	}

	roleData, err := e.roleRepository.GetByTitle(ctx, input.Role)
	if err != nil {
		return err
	}

	salaryData, err := e.salaryRepository.GetByRoleID(ctx, roleData.ID)
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

	countEmployee, err := e.employeeRepository.CountEmployeeByEmail(ctx, employeeData.Email)
	if err != nil {
		return err
	}

	if countEmployee > 0 {
		return errors.New("the employee's data is exist")
	}

	_ = e.cache.Delete("employees")

	return e.employeeRepository.CreateOne(ctx, employeeData)
}

func (e *employeeUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_ = e.cache.Delete(id)
	_ = e.cache.Delete("employees")

	return e.employeeRepository.DeleteOne(ctx, employeeID)
}

func (e *employeeUseCase) UpdateOne(ctx context.Context, id string, input *employeesdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	if err := validate.Employee(input); err != nil {
		return err
	}

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	departmentData, err := e.departmentRepository.GetByName(ctx, input.Department)
	if err != nil {
		return err
	}

	roleData, err := e.roleRepository.GetByTitle(ctx, input.Role)
	if err != nil {
		return err
	}

	salaryData, err := e.salaryRepository.GetByID(ctx, roleData.ID)
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

	_ = e.cache.Delete(id)
	_ = e.cache.Delete("employees")

	return e.employeeRepository.UpdateOne(ctx, employeeID, employee)
}

func (e *employeeUseCase) UpdateStatus(ctx context.Context, id string, isActive bool) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_ = e.cache.Delete(id)
	_ = e.cache.Delete("employees")

	return e.employeeRepository.UpdateStatus(ctx, employeeID, isActive)
}

func (e *employeeUseCase) GetByID(ctx context.Context, id string) (employeesdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	data, _ := e.cache.Get(id)
	if data != nil {
		var response employeesdomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return employeesdomain.Output{}, err
		}
		return response, nil
	}

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	employeeData, err := e.employeeRepository.GetByID(ctx, employeeID)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	output := employeesdomain.Output{
		Employee: *employeeData,
	}

	data, err = json.Marshal(output)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	err = e.cache.Set(id, data)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	return output, nil
}

func (e *employeeUseCase) GetByEmail(ctx context.Context, name string) (employeesdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	data, _ := e.cache.Get(name)
	if data != nil {
		var response employeesdomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return employeesdomain.Output{}, err
		}
		return response, nil
	}

	employeeData, err := e.employeeRepository.GetByEmail(ctx, name)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	output := employeesdomain.Output{
		Employee: *employeeData,
	}

	data, err = json.Marshal(output)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	err = e.cache.Set(name, data)
	if err != nil {
		return employeesdomain.Output{}, err
	}

	return output, nil
}

func (e *employeeUseCase) GetAll(ctx context.Context) ([]employeesdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	data, _ := e.cache.Get("employees")
	if data != nil {
		var response []employeesdomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

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

	data, err = json.Marshal(outputs)
	if err != nil {
		return nil, err
	}

	err = e.cache.Set("employees", data)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}

func (e *employeeUseCase) CountEmployee(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	return e.employeeRepository.CountEmployee(ctx)
}
