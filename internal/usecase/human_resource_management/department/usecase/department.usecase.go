package department_usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	departmentsdomain "shop_erp_mono/internal/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	"shop_erp_mono/internal/usecase/human_resource_management/department/validate"
	"time"
)

type departmentUseCase struct {
	contextTimeout       time.Duration
	departmentRepository departmentsdomain.IDepartmentRepository
	employeeRepository   employeesdomain.IEmployeeRepository
	client               *mongo.Client
	cache                *bigcache.BigCache
}

func NewDepartmentUseCase(contextTimeout time.Duration, departmentRepository departmentsdomain.IDepartmentRepository,
	employeeRepository employeesdomain.IEmployeeRepository, cacheTTL time.Duration, client *mongo.Client) departmentsdomain.IDepartmentUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &departmentUseCase{contextTimeout: contextTimeout, cache: cache, departmentRepository: departmentRepository, employeeRepository: employeeRepository, client: client}
}

func (d *departmentUseCase) CreateOne(ctx context.Context, input *departmentsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	if err := validate.Department(input); err != nil {
		return err
	}

	count, err := d.departmentRepository.CountDepartmentWithName(ctx, input.Name)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("name of department is exist")
	}

	department := &departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_ = d.cache.Delete("departments")

	return d.departmentRepository.CreateOne(ctx, department)
}

func (d *departmentUseCase) CreateDepartmentWithManager(ctx context.Context, departmentInput *departmentsdomain.Input,
	employeeInput *employeesdomain.Input) error {

	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	session, err := d.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		// Step 1: Create department
		department := &departmentsdomain.Department{
			ID:          primitive.NewObjectID(),
			Name:        departmentInput.Name,
			Description: departmentInput.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		count, err := d.departmentRepository.CountDepartmentWithName(sessionCtx, departmentInput.Name)
		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, errors.New("name of department is exist")
		}

		err = d.departmentRepository.CreateOne(sessionCtx, department)
		if err != nil {
			return nil, err
		}

		// Step 2: Check if employee exists
		managerData, err := d.employeeRepository.GetByEmail(sessionCtx, employeeInput.Email)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}

		var employeeID primitive.ObjectID

		if managerData != nil {
			// If employee exists, check if they are managing another department
			count, err := d.departmentRepository.CountManagerExist(sessionCtx, managerData.ID)
			if err != nil {
				return nil, err
			}

			if count > 0 {
				return nil, errors.New("the employee is managing in another department")
			}

			employeeID = managerData.ID
		} else {
			// If employee does not exist, create a new employee
			employee := employeesdomain.Employee{
				ID:           primitive.NewObjectID(),
				FirstName:    employeeInput.FirstName,
				LastName:     employeeInput.LastName,
				Gender:       employeeInput.Gender,
				Email:        employeeInput.Email,
				Phone:        employeeInput.Phone,
				Address:      employeeInput.Address,
				AvatarURL:    employeeInput.AvatarURL,
				DateOfBirth:  employeeInput.DateOfBirth,
				DayOfWork:    employeeInput.DayOfWork,
				DepartmentID: department.ID,
				IsActive:     true,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			err = d.employeeRepository.CreateOne(sessionCtx, &employee)
			if err != nil {
				return nil, err
			}
			employeeID = employee.ID
		}

		// Step 3: Update department with ManagerID
		err = d.departmentRepository.UpdateManager(sessionCtx, department.ID, employeeID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	// Run the transaction
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	_ = d.cache.Delete("departments")

	return session.CommitTransaction(ctx)
}

func (d *departmentUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	departmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_ = d.cache.Delete("departments")
	_ = d.cache.Delete("departments")

	return d.departmentRepository.DeleteOne(ctx, departmentID)
}

func (d *departmentUseCase) UpdateOne(ctx context.Context, id string, input *departmentsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	if err := validate.Department(input); err != nil {
		return err
	}

	session, err := d.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		managerData, err := d.employeeRepository.GetByEmail(ctx, input.ManagerEmail)
		if err != nil {
			return nil, err
		}

		count, err := d.departmentRepository.CountManagerExist(ctx, managerData.ID)
		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, errors.New("the employee is managing in other department")
		}

		departmentID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}

		department := &departmentsdomain.Department{
			ID:          primitive.NewObjectID(),
			Name:        input.Name,
			Description: input.Description,
			UpdatedAt:   time.Now(),
		}

		err = d.departmentRepository.UpdateOne(ctx, departmentID, department)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	// Run the transaction
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	_ = d.cache.Delete("departments")
	_ = d.cache.Delete("departments")

	return session.CommitTransaction(ctx)
}

func (d *departmentUseCase) GetByID(ctx context.Context, id string) (departmentsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	data, err := d.cache.Get(id)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	if data != nil {
		var response departmentsdomain.Output
		err = json.Unmarshal(data, &response)
		return response, nil
	}

	departmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	departmentData, err := d.departmentRepository.GetByID(ctx, departmentID)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	output := departmentsdomain.Output{
		Department: departmentData,
	}

	data, err = json.Marshal(output)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	err = d.cache.Set(id, data)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	return output, nil
}

func (d *departmentUseCase) GetByName(ctx context.Context, name string) (departmentsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	data, err := d.cache.Get(name)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	if data != nil {
		var response departmentsdomain.Output
		err = json.Unmarshal(data, &response)
		return response, nil
	}

	departmentData, err := d.departmentRepository.GetByName(ctx, name)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	output := departmentsdomain.Output{
		Department: departmentData,
	}

	data, err = json.Marshal(output)
	if err != nil {
		return departmentsdomain.Output{}, err
	}

	err = d.cache.Set(name, data)
	if err != nil {
		return departmentsdomain.Output{}, err
	}
	return output, nil
}

// GetAll không cần phân trang vì không có công ty nào có tới 1000 phòng ban
func (d *departmentUseCase) GetAll(ctx context.Context) ([]departmentsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	data, err := d.cache.Get("departments")
	if data != nil {
		var response []departmentsdomain.Output
		err = json.Unmarshal(data, &response)
		return response, nil
	}

	departmentsData, err := d.departmentRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []departmentsdomain.Output
	outputs = make([]departmentsdomain.Output, 0, len(departmentsData))
	for _, departmentData := range departmentsData {
		managerData, err := d.employeeRepository.GetByID(ctx, departmentData.ManagerID)
		if err != nil {
			return nil, err
		}

		output := departmentsdomain.Output{
			Department: departmentData,
			Manager:    *managerData,
		}

		outputs = append(outputs, output)
	}

	data, err = json.Marshal(outputs)
	if err != nil {
		return nil, err
	}

	err = d.cache.Set("departments", data)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}

func (d *departmentUseCase) CountManagerExist(ctx context.Context, managerID primitive.ObjectID) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	return d.departmentRepository.CountManagerExist(ctx, managerID)
}

func (d *departmentUseCase) CountDepartment(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	return d.departmentRepository.CountDepartment(ctx)
}
