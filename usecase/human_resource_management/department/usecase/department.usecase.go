package department_usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/usecase/human_resource_management/department/validate"
	"sync"
	"time"
)

type departmentUseCase struct {
	contextTimeout       time.Duration
	departmentRepository departmentsdomain.IDepartmentRepository
	employeeRepository   employeesdomain.IEmployeeRepository
	cache                *bigcache.BigCache
}

func NewDepartmentUseCase(contextTimeout time.Duration, departmentRepository departmentsdomain.IDepartmentRepository,
	employeeRepository employeesdomain.IEmployeeRepository, cacheTTL time.Duration) departmentsdomain.IDepartmentUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &departmentUseCase{contextTimeout: contextTimeout, cache: cache, departmentRepository: departmentRepository, employeeRepository: employeeRepository}
}

func (d *departmentUseCase) CreateOne(ctx context.Context, input *departmentsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	if err := validate.Department(input); err != nil {
		return err
	}

	managerData, err := d.employeeRepository.GetByEmail(ctx, input.ManagerEmail)
	if err != nil {
		return err
	}

	count, err := d.departmentRepository.CountManagerExist(ctx, managerData.ID)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("the employee is managing in other department")
	}

	department := &departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		ManagerID:   managerData.ID,
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = d.cache.Delete("departments")
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
		return err
	case <-ctx.Done():
		return ctx.Err()
	default:
		return d.departmentRepository.CreateOne(ctx, department)
	}
}

func (d *departmentUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	departmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err = d.cache.Delete("departments"); err != nil {
			errCh <- err
			return
		}
	}()
	go func() {
		defer wg.Done()
		if err = d.cache.Delete("departments"); err != nil {
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
		return err
	case <-ctx.Done():
		return ctx.Err()
	default:
		return d.departmentRepository.DeleteOne(ctx, departmentID)
	}
}

func (d *departmentUseCase) UpdateOne(ctx context.Context, id string, input *departmentsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, d.contextTimeout)
	defer cancel()

	if err := validate.Department(input); err != nil {
		return err
	}

	managerData, err := d.employeeRepository.GetByEmail(ctx, input.ManagerEmail)
	if err != nil {
		return err
	}

	count, err := d.departmentRepository.CountManagerExist(ctx, managerData.ID)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("the employee is managing in other department")
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

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err = d.cache.Delete("departments"); err != nil {
			errCh <- err
			return
		}
	}()
	go func() {
		defer wg.Done()
		if err = d.cache.Delete("departments"); err != nil {
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
		return err
	case <-ctx.Done():
		return ctx.Err()
	default:
		return d.departmentRepository.UpdateOne(ctx, departmentID, department)
	}
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
			Manager:    managerData,
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
