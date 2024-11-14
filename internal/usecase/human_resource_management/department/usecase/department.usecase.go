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
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	"shop_erp_mono/internal/usecase/human_resource_management/department/validate"
	"shop_erp_mono/pkg/shared/constant"
	"strconv"
	"sync"
	"time"
)

type departmentUseCase struct {
	contextTimeout       time.Duration
	departmentRepository departmentsdomain.IDepartmentRepository
	employeeRepository   employeesdomain.IEmployeeRepository
	userRepository       userdomain.IUserRepository
	client               *mongo.Client
	cache                *bigcache.BigCache
}

func NewDepartmentUseCase(contextTimeout time.Duration, departmentRepository departmentsdomain.IDepartmentRepository,
	employeeRepository employeesdomain.IEmployeeRepository, userRepository userdomain.IUserRepository, cacheTTL time.Duration, client *mongo.Client) departmentsdomain.IDepartmentUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &departmentUseCase{contextTimeout: contextTimeout, cache: cache, departmentRepository: departmentRepository, userRepository: userRepository, employeeRepository: employeeRepository, client: client}
}

func (d *departmentUseCase) CreateOne(ctx context.Context, input *departmentsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	if err := validate.Department(input); err != nil {
		return err
	}

	// không được trùng tên department
	count, err := d.departmentRepository.CountDepartmentWithName(ctx, input.Name)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("name of department is exist")
	}

	if input.Level == 1 {
		if input.ParentID != primitive.NilObjectID {
			return errors.New("a department with Level 1 cannot have a ParentID")
		}
	} else {
		// Nếu có ParentID, kiểm tra tính hợp lệ
		if input.ParentID != primitive.NilObjectID {
			parentDept, err := d.departmentRepository.GetByID(ctx, input.ParentID)
			if err != nil {
				return errors.New("parent department not found")
			}

			// Kiểm tra điều kiện hợp lệ cho Level
			if parentDept.Level != input.Level-1 {
				return errors.New("invalid level: parent department must have a level of " + strconv.Itoa(input.Level-1))
			}
		} else {
			return errors.New("a department with Level " + strconv.Itoa(input.Level) + " requires a valid ParentID")
		}
	}

	department := &departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		Name:        input.Name,
		Description: input.Description,
		ParentID:    input.ParentID,
		Level:       input.Level,
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
		// Không được trùng tên department
		count, err := d.departmentRepository.CountDepartmentWithName(sessionCtx, departmentInput.Name)
		if err != nil {
			return nil, err
		}

		if count > 0 {
			return nil, errors.New("name of department is exist")
		}

		if count > 0 {
			return nil, errors.New("name of department is exist")
		}

		if departmentInput.Level == 1 {
			if departmentInput.ParentID != primitive.NilObjectID {
				return nil, errors.New("a department with Level 1 cannot have a ParentID")
			}
		} else {
			// Nếu có ParentID, kiểm tra tính hợp lệ
			if departmentInput.ParentID != primitive.NilObjectID {
				parentDept, err := d.departmentRepository.GetByID(ctx, departmentInput.ParentID)
				if err != nil {
					return nil, errors.New("parent department not found")
				}

				// Kiểm tra điều kiện hợp lệ cho Level
				if parentDept.Level != departmentInput.Level-1 {
					return nil, errors.New("invalid level: parent department must have a level of " + strconv.Itoa(departmentInput.Level-1))
				}
			} else {
				return nil, errors.New("a department with Level " + strconv.Itoa(departmentInput.Level) + " requires a valid ParentID")
			}
		}

		// Step 1: Create department
		department := &departmentsdomain.Department{
			ID:          primitive.NewObjectID(),
			Name:        departmentInput.Name,
			Description: departmentInput.Description,
			ParentID:    departmentInput.ParentID,
			Level:       departmentInput.Level,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
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

func (d *departmentUseCase) DeleteOne(ctx context.Context, id string, userid string) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	departmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid department ID format")
	}

	userID, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	userData, err := d.userRepository.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	if userData.Role != constant.RoleSuperAdmin {
		return errors.New("permission denied: only users with the highest role can delete departments")
	}

	departmentData, err := d.departmentRepository.GetByID(ctx, departmentID)
	if err != nil {
		return errors.New("department not found")
	}

	if departmentData.ManagerID == primitive.NilObjectID {
		return errors.New("cannot delete department with null manager")
	}

	if err = d.departmentRepository.DeleteOne(ctx, departmentID); err != nil {
		return err
	}

	_ = d.cache.Delete("departments")
	return nil
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
	return session.CommitTransaction(ctx)
}

func (d *departmentUseCase) UpdateManager(ctx context.Context, id string, managerID string) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	departmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid department ID format")
	}

	idManager, err := primitive.ObjectIDFromHex(managerID)
	if err != nil {
		return errors.New("invalid employee ID format")
	}

	managerData, err := d.employeeRepository.GetByID(ctx, idManager)
	if err != nil {
		return err
	}

	return d.departmentRepository.UpdateManager(ctx, departmentID, managerData.ID)
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
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	errCh := make(chan error, 1)
	var wg sync.WaitGroup
	var mutex sync.Mutex // Mutex để đồng bộ thao tác với slice outputs

	departmentsData, err := d.departmentRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []departmentsdomain.Output
	outputs = make([]departmentsdomain.Output, 0, len(departmentsData))

	for _, departmentData := range departmentsData {
		wg.Add(1)
		go func(departmentData departmentsdomain.Department) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
				managerData, err := d.employeeRepository.GetByID(ctx, departmentData.ManagerID)
				if err != nil {
					errCh <- err
					return
				}

				output := departmentsdomain.Output{
					Department: departmentData,
					Manager:    *managerData,
				}

				mutex.Lock()
				outputs = append(outputs, output)
				mutex.Unlock()
			}
		}(departmentData)
	}

	wg.Wait()
	close(errCh)

	select {
	case err = <-errCh:
		if err != nil {
			return nil, err
		}
	default:
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
