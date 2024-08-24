package employee_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	departments_domain "shop_erp_mono/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"shop_erp_mono/repository/human_resource_management/employee/validate"
	"sync"
	"time"
)

type employeeRepository struct {
	database             *mongo.Database
	collectionEmployee   string
	collectionDepartment string
	collectionRole       string
	collectionSalary     string
}

func NewEmployeeRepository(db *mongo.Database, collectionEmployee string, collectionDepartment string, collectionRole string, collectionSalary string) employeesdomain.IEmployeeRepository {
	return &employeeRepository{database: db, collectionEmployee: collectionEmployee, collectionDepartment: collectionDepartment, collectionRole: collectionRole, collectionSalary: collectionSalary}
}

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func (e employeeRepository) CreateOne(ctx context.Context, employee *employeesdomain.Input) error {
	collectionEmployee := e.database.Collection(e.collectionEmployee)
	collectionDepartment := e.database.Collection(e.collectionDepartment)
	collectionRole := e.database.Collection(e.collectionRole)
	collectionSalary := e.database.Collection(e.collectionSalary)

	if err := validate.IsNilEmployee(employee); err != nil {
		return err
	}

	filterDepart := bson.M{"name": employee.Department}
	var department departments_domain.Department
	if err := collectionDepartment.FindOne(ctx, filterDepart).Decode(&department); err != nil {
		return err
	}

	filterRole := bson.M{"title": employee.Role}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filterRole).Decode(&role); err != nil {
		return err
	}

	filterSalary := bson.M{"role_id": role.ID}
	var salary salarydomain.Salary
	if err := collectionSalary.FindOne(ctx, filterSalary).Decode(&salary); err != nil {
		return err
	}

	employeeData := employeesdomain.Employee{
		ID:           primitive.NewObjectID(),
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		Gender:       employee.Gender,
		Email:        employee.Email,
		Phone:        employee.Phone,
		Address:      employee.Address,
		AvatarURL:    employee.AvatarURL,
		DateOfBirth:  employee.DateOfBirth,
		DayOfWork:    employee.DayOfWork,
		DepartmentID: department.ID,
		SalaryID:     salary.ID,
		RoleID:       role.ID,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	mutex.Lock()
	_, err := collectionEmployee.InsertOne(ctx, employeeData)
	if err != nil {
		return err
	}
	mutex.Unlock()

	return nil
}

func (e employeeRepository) DeleteOne(ctx context.Context, id string) error {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid employee ID format")
	}
	if employeeID == primitive.NilObjectID {
		return errors.New("employee ID cannot be null")
	}

	filter := bson.M{"_id": employeeID}

	// Sử dụng defer để đảm bảo mutex luôn được mở khóa
	mutex.Lock()
	defer mutex.Unlock()

	_, err = collectionEmployee.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (e employeeRepository) UpdateOne(ctx context.Context, employee *employeesdomain.Input) error {
	collectionEmployee := e.database.Collection(e.collectionEmployee)
	collectionDepartment := e.database.Collection(e.collectionDepartment)
	collectionRole := e.database.Collection(e.collectionRole)
	collectionSalary := e.database.Collection(e.collectionSalary)

	if err := validate.IsNilEmployee(employee); err != nil {
		return err
	}

	filterDepart := bson.M{"name": employee.Department}
	var department departments_domain.Department
	if err := collectionDepartment.FindOne(ctx, filterDepart).Decode(&department); err != nil {
		return err
	}

	filterRole := bson.M{"title": employee.Role}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filterRole).Decode(&role); err != nil {
		return err
	}

	filterSalary := bson.M{"role_id": role.ID}
	var salary salarydomain.Salary
	if err := collectionSalary.FindOne(ctx, filterSalary).Decode(&salary); err != nil {
		return err
	}

	employeeData := employeesdomain.Employee{
		ID:           primitive.NewObjectID(),
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		Gender:       employee.Gender,
		Email:        employee.Email,
		Phone:        employee.Phone,
		Address:      employee.Address,
		AvatarURL:    employee.AvatarURL,
		DateOfBirth:  employee.DateOfBirth,
		DayOfWork:    employee.DayOfWork,
		DepartmentID: department.ID,
		SalaryID:     salary.ID,
		RoleID:       role.ID,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	filter := bson.M{"_id": employeeData.ID}
	update := bson.M{"$set": bson.M{
		"first_name":    employeeData.FirstName,
		"last_name":     employeeData.LastName,
		"gender":        employeeData.Gender,
		"address":       employeeData.Address,
		"date_of_birth": employeeData.DateOfBirth,
		"day_of_work":   employeeData.DayOfWork,
		"department_id": employeeData.DepartmentID,
		"role_id":       employeeData.RoleID,
		"salary_id":     employeeData.SalaryID,
		"updated_at":    employeeData.UpdatedAt,
		"is_active":     employeeData.IsActive,
	}}

	mutex.Lock()
	_, err := collectionEmployee.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating role's information into database ")
	}
	mutex.Unlock()

	return nil
}

func (e employeeRepository) GetOneByID(ctx context.Context, id string) (employeesdomain.Output, error) {
	errCh := make(chan error, 1) // Chỉ cần một lỗi duy nhất
	departmentCh := make(chan departments_domain.Department, 1)
	roleCh := make(chan roledomain.Role, 1)
	salaryCh := make(chan salarydomain.Salary, 1)

	collectionEmployee := e.database.Collection(e.collectionEmployee)
	collectionDepartment := e.database.Collection(e.collectionDepartment)
	collectionRole := e.database.Collection(e.collectionRole)
	collectionSalary := e.database.Collection(e.collectionSalary)

	// Chuyển đổi ID và kiểm tra lỗi
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return employeesdomain.Output{}, errors.New("invalid employee ID format")
	}

	var employee employeesdomain.Employee
	filter := bson.M{"_id": employeeID}
	if err := collectionEmployee.FindOne(ctx, filter).Decode(&employee); err != nil {
		return employeesdomain.Output{}, errors.New("error finding employee's information in the database")
	}

	wg.Add(3)

	// Goroutine để lấy thông tin phòng ban
	go func() {
		defer wg.Done()
		var department departments_domain.Department
		filterDepartment := bson.M{"_id": employee.DepartmentID}
		if err := collectionDepartment.FindOne(ctx, filterDepartment).Decode(&department); err != nil {
			errCh <- errors.New("error finding department information in the database")
			return
		}
		departmentCh <- department
	}()

	// Goroutine để lấy thông tin vai trò
	go func() {
		defer wg.Done()
		var role roledomain.Role
		filterRole := bson.M{"_id": employee.RoleID}
		if err := collectionRole.FindOne(ctx, filterRole).Decode(&role); err != nil {
			errCh <- errors.New("error finding role information in the database")
			return
		}
		roleCh <- role
	}()

	// Goroutine để lấy thông tin lương
	go func() {
		defer wg.Done()
		var salary salarydomain.Salary
		filterSalary := bson.M{"_id": employee.SalaryID}
		if err := collectionSalary.FindOne(ctx, filterSalary).Decode(&salary); err != nil {
			errCh <- errors.New("error finding salary information in the database")
			return
		}
		salaryCh <- salary
	}()

	// Goroutine chờ đợi tất cả các goroutine hoàn thành
	go func() {
		wg.Wait()
		close(departmentCh)
		close(roleCh)
		close(salaryCh)
		close(errCh)
	}()

	// Lắng nghe kênh lỗi trước tiên
	select {
	case err = <-errCh:
		return employeesdomain.Output{}, err
	default:
		// Đảm bảo nhận đủ dữ liệu từ tất cả các kênh
		department := <-departmentCh
		role := <-roleCh
		salary := <-salaryCh

		employeeOutput := employeesdomain.Output{
			Employee:     employee,
			DepartmentID: department.Name,
			RoleID:       role.Title,
			Salary:       salary,
		}

		return employeeOutput, nil
	}
}

func (e employeeRepository) GetOneByName(ctx context.Context, name string) (employeesdomain.Output, error) {
	errCh := make(chan error)
	departmentCh := make(chan departments_domain.Department, 1)
	roleCh := make(chan roledomain.Role, 1)
	salaryCh := make(chan salarydomain.Salary, 1)

	collectionEmployee := e.database.Collection(e.collectionEmployee)
	collectionDepartment := e.database.Collection(e.collectionDepartment) // Sửa thành e.collectionDepartment
	collectionRole := e.database.Collection(e.collectionRole)
	collectionSalary := e.database.Collection(e.collectionSalary)

	var employee employeesdomain.Employee
	filter := bson.M{"last_name": name}
	if err := collectionEmployee.FindOne(ctx, filter).Decode(&employee); err != nil {
		return employeesdomain.Output{}, errors.New("error finding employee information in database")
	}

	wg.Add(3)
	go func(employee employeesdomain.Employee) {
		defer wg.Done()
		var department departments_domain.Department
		filterDepartment := bson.M{"_id": employee.DepartmentID}
		if err := collectionDepartment.FindOne(ctx, filterDepartment).Decode(&department); err != nil {
			errCh <- errors.New("error finding department information in database")
			return
		}
		departmentCh <- department
	}(employee)

	go func(employee employeesdomain.Employee) {
		defer wg.Done()
		var role roledomain.Role
		filterRole := bson.M{"_id": employee.RoleID}
		if err := collectionRole.FindOne(ctx, filterRole).Decode(&role); err != nil {
			errCh <- errors.New("error finding role information in database")
			return
		}
		roleCh <- role
	}(employee)

	go func(employee employeesdomain.Employee) {
		defer wg.Done()
		var salary salarydomain.Salary
		filterSalary := bson.M{"_id": employee.SalaryID} // Sửa thành "salary_id"
		if err := collectionSalary.FindOne(ctx, filterSalary).Decode(&salary); err != nil {
			errCh <- errors.New("error finding salary information in database")
			return
		}
		salaryCh <- salary
	}(employee)

	go func() {
		wg.Wait()
		defer close(departmentCh)
		defer close(roleCh)
		defer close(salaryCh)
	}()

	select {
	case err := <-errCh:
		return employeesdomain.Output{}, err
	default:
		department := <-departmentCh
		role := <-roleCh
		salary := <-salaryCh

		employeeOutput := employeesdomain.Output{
			Employee:     employee,
			DepartmentID: department.Name,
			Salary:       salary,
			RoleID:       role.Title,
		}

		return employeeOutput, nil
	}
}

func (e employeeRepository) GetAll(ctx context.Context) ([]employeesdomain.Output, error) {
	errCh := make(chan error, 1)
	departmentCh := make(chan departments_domain.Department, 1)
	roleCh := make(chan roledomain.Role, 1)
	salaryCh := make(chan salarydomain.Salary, 1)

	collectionEmployee := e.database.Collection(e.collectionEmployee)
	collectionDepartment := e.database.Collection(e.collectionDepartment) // Sửa thành e.collectionDepartment
	collectionRole := e.database.Collection(e.collectionRole)
	collectionSalary := e.database.Collection(e.collectionSalary)

	filter := bson.M{}
	cursor, err := collectionEmployee.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var employees []employeesdomain.Output
	employees = make([]employeesdomain.Output, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var employee employeesdomain.Employee
		if err = cursor.Decode(&employee); err != nil {
			return nil, errors.New("error decoding employee information from database")
		}

		wg.Add(3)
		go func(employee employeesdomain.Employee) {
			defer wg.Done()
			var department departments_domain.Department
			filterDepartment := bson.M{"_id": employee.DepartmentID}
			if err = collectionDepartment.FindOne(ctx, filterDepartment).Decode(&department); err != nil {
				errCh <- errors.New("error finding department information in database")
				return
			}
			departmentCh <- department
		}(employee)

		go func(employee employeesdomain.Employee) {
			defer wg.Done()
			var role roledomain.Role
			filterRole := bson.M{"_id": employee.RoleID}
			if err = collectionRole.FindOne(ctx, filterRole).Decode(&role); err != nil {
				errCh <- errors.New("error finding role information in database")
				return
			}
			roleCh <- role
		}(employee)

		go func(employee employeesdomain.Employee) {
			defer wg.Done()
			var salary salarydomain.Salary
			filterSalary := bson.M{"_id": employee.SalaryID} // Đúng tên field là "salary_id"
			if err = collectionSalary.FindOne(ctx, filterSalary).Decode(&salary); err != nil {
				errCh <- errors.New("error finding salary information in database")
				return
			}
			salaryCh <- salary
		}(employee)

		// Đảm bảo đợi các goroutine kết thúc trước khi tiếp tục xử lý
		go func() {
			wg.Wait()
			close(departmentCh)
			close(roleCh)
			close(salaryCh)
			close(errCh)
		}()

		select {
		case err = <-errCh:
			return nil, err
		default:
			role := <-roleCh
			salary := <-salaryCh
			department := <-departmentCh

			employeeOutput := employeesdomain.Output{
				Employee:     employee,
				DepartmentID: department.Name,
				Salary:       salary,
				RoleID:       role.Title,
			}

			employees = append(employees, employeeOutput)
		}
	}

	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	return employees, nil
}
