package employee_usecase

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"shop_erp_mono/usecase/human_resource_management/employee/validate"
	"sync"
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

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = e.cache.Delete("employees")
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err = <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}
	return e.employeeRepository.CreateOne(ctx, employeeData)
}

func (e *employeeUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	errCh := make(chan error, 1)
	var wg sync.WaitGroup
	var once sync.Once

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Helper function to send error to errCh
	sendError := func(err error) {
		once.Do(func() {
			errCh <- err
		})
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err = e.cache.Delete(id); err != nil {
			sendError(err)
		}
	}()

	go func() {
		defer wg.Done()
		if err = e.cache.Delete("employees"); err != nil {
			sendError(err)
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err = <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}
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

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err = e.cache.Delete(id); err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err = e.cache.Delete("employees"); err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err = <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}
	return e.employeeRepository.UpdateOne(ctx, employeeID, employee)
}

func (e *employeeUseCase) UpdateStatus(ctx context.Context, id string, isActive bool) error {
	ctx, cancel := context.WithTimeout(ctx, e.contextTimeout)
	defer cancel()

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err = e.cache.Delete(id); err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err = e.cache.Delete("employees"); err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err = <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}
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
		Employee: employeeData,
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
		Employee: employeeData,
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
