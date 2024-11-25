package employee_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	"sync"
	"time"
)

type employeeRepository struct {
	database           *mongo.Database
	collectionEmployee string
}

func NewEmployeeRepository(db *mongo.Database, collectionEmployee string) employeesdomain.IEmployeeRepository {
	return &employeeRepository{database: db, collectionEmployee: collectionEmployee}
}

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func (e *employeeRepository) CreateOne(ctx context.Context, employee *employeesdomain.Employee) error {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	// Sử dụng defer để đảm bảo mutex luôn được mở khóa
	mutex.Lock()
	defer mutex.Unlock()

	_, err := collectionEmployee.InsertOne(ctx, employee)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	if id == primitive.NilObjectID {
		return errors.New("employee ID cannot be null")
	}
	filter := bson.M{"_id": id}

	// Sử dụng defer để đảm bảo mutex luôn được mở khóa
	mutex.Lock()
	defer mutex.Unlock()

	_, err := collectionEmployee.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, employee *employeesdomain.Employee) error {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	if id == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"first_name":    employee.FirstName,
		"last_name":     employee.LastName,
		"gender":        employee.Gender,
		"address":       employee.Address,
		"avatar_url":    employee.AvatarURL,
		"date_of_birth": employee.DateOfBirth,
		"day_of_work":   employee.DayOfWork,
		"department_id": employee.DepartmentID,
		"role_id":       employee.RoleID,
		"salary_id":     employee.SalaryID,
		"updated_at":    employee.UpdatedAt,
		"is_active":     employee.IsActive,
	}}

	// Sử dụng defer để đảm bảo mutex luôn được mở khóa
	mutex.Lock()
	defer mutex.Unlock()

	_, err := collectionEmployee.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating role's information into database ")
	}

	return nil
}

func (e *employeeRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, isActive bool) error {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	if id == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"updated_at": time.Now(),
		"is_active":  isActive,
	}}

	// Sử dụng defer để đảm bảo mutex luôn được mở khóa
	mutex.Lock()
	defer mutex.Unlock()

	_, err := collectionEmployee.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating role's information into database ")
	}

	return nil
}

func (e *employeeRepository) GetByID(ctx context.Context, id primitive.ObjectID) (employeesdomain.Employee, error) {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	var employee employeesdomain.Employee
	filter := bson.M{"_id": id}
	if err := collectionEmployee.FindOne(ctx, filter).Decode(&employee); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return employeesdomain.Employee{}, nil
		}
		return employeesdomain.Employee{}, errors.New("error finding employee's information in the database")
	}

	return employee, nil
}

func (e *employeeRepository) GetByName(ctx context.Context, name string) (employeesdomain.Employee, error) {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	var employee employeesdomain.Employee
	filter := bson.M{"last_name": name}
	if err := collectionEmployee.FindOne(ctx, filter).Decode(&employee); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return employeesdomain.Employee{}, nil
		}
		return employeesdomain.Employee{}, errors.New("error finding employee information in database")
	}

	return employee, nil
}

func (e *employeeRepository) GetByEmail(ctx context.Context, email string) (employeesdomain.Employee, error) {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	filter := bson.M{"email": email}
	var employee employeesdomain.Employee
	err := collectionEmployee.FindOne(ctx, filter).Decode(&employee)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return employeesdomain.Employee{}, nil
		}
		return employeesdomain.Employee{}, err
	}

	return employee, nil
}

func (e *employeeRepository) GetAll(ctx context.Context) ([]employeesdomain.Employee, error) {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	filter := bson.M{}
	cursor, err := collectionEmployee.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var employees []employeesdomain.Employee
	employees = make([]employeesdomain.Employee, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var employee employeesdomain.Employee
		if err = cursor.Decode(&employee); err != nil {
			return nil, errors.New("error decoding employee information from database")
		}

		employees = append(employees, employee)
	}

	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *employeeRepository) CountEmployeeByEmail(ctx context.Context, email string) (int64, error) {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	filter := bson.M{"email": email}
	count, err := collectionEmployee.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (e *employeeRepository) CountEmployeeByDepartmentID(ctx context.Context, departmentID primitive.ObjectID) (int64, error) {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	filter := bson.M{"department_id": departmentID}
	count, err := collectionEmployee.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (e *employeeRepository) CountEmployee(ctx context.Context) (int64, error) {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	filter := bson.M{}
	count, err := collectionEmployee.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (e *employeeRepository) CountEmployeeByRoleID(ctx context.Context, roleID primitive.ObjectID) (int64, error) {
	collectionEmployee := e.database.Collection(e.collectionEmployee)

	filter := bson.M{"role_id": roleID}
	count, err := collectionEmployee.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
