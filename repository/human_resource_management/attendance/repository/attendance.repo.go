package attendance_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/repository/human_resource_management/attendance/validate"
	"time"
)

type attendanceRepository struct {
	database             *mongo.Database
	collectionAttendance string
	collectionEmployee   string
}

func NewAttendanceRepository(db *mongo.Database, collectionAttendance string, collectionEmployee string) attendancedomain.IAttendanceRepository {
	return &attendanceRepository{database: db, collectionAttendance: collectionAttendance, collectionEmployee: collectionEmployee}
}

func (a attendanceRepository) CreateOne(ctx context.Context, input *attendancedomain.Input) error {
	collectionAttendance := a.database.Collection(a.collectionAttendance)
	collectionEmployee := a.database.Collection(a.collectionEmployee)

	if err := validate.IsNilAttendance(input); err != nil {
		return err
	}

	filterEmployee := bson.M{"email": input.EmailEmployee}
	var employee employeesdomain.Employee
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return err
	}

	hoursWorked := input.CheckOutTime.Sub(input.CheckInTime)

	attendance := attendancedomain.Attendance{
		ID:           primitive.NewObjectID(),
		EmployeeID:   employee.ID,
		Date:         input.Date,
		CheckInTime:  input.CheckInTime,
		CheckOutTime: input.CheckOutTime,
		HoursWorked:  int8(hoursWorked.Hours()),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err := collectionAttendance.InsertOne(ctx, attendance)
	if err != nil {
		return err
	}

	return nil
}

func (a attendanceRepository) DeleteOne(ctx context.Context, id string) error {
	collectionAttendance := a.database.Collection(a.collectionAttendance)

	attendanceID, _ := primitive.ObjectIDFromHex(id)
	if attendanceID == primitive.NilObjectID {
		return errors.New("error the id do not nil")
	}

	_, err := collectionAttendance.DeleteOne(ctx, attendanceID)
	if err != nil {
		return err
	}

	return nil
}

func (a attendanceRepository) UpdateOne(ctx context.Context, id string, input *attendancedomain.Input) error {
	collectionAttendance := a.database.Collection(a.collectionAttendance)
	collectionEmployee := a.database.Collection(a.collectionEmployee)

	if err := validate.IsNilAttendance(input); err != nil {
		return err
	}

	attendanceID, _ := primitive.ObjectIDFromHex(id)
	if attendanceID == primitive.NilObjectID {
		return errors.New("the id do not nil")
	}

	filterEmployee := bson.M{"email": input.EmailEmployee}
	var employee employeesdomain.Employee
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return err
	}

	hoursWorked := input.CheckOutTime.Sub(input.CheckInTime)

	filter := bson.M{"_id": attendanceID}
	update := bson.M{"$set": bson.M{
		"employee_id":    employee,
		"date":           input.Date,
		"check_in_time":  input.CheckInTime,
		"check_out_time": input.CheckOutTime,
		"hours_worked":   int8(hoursWorked.Hours()),
		"status":         input.Status,
	}}

	_, err := collectionAttendance.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (a attendanceRepository) GetOneByID(ctx context.Context, id string) (attendancedomain.Output, error) {
	collectionAttendance := a.database.Collection(a.collectionAttendance)
	collectionEmployee := a.database.Collection(a.collectionEmployee)

	attendanceID, _ := primitive.ObjectIDFromHex(id)
	if attendanceID == primitive.NilObjectID {
		return attendancedomain.Output{}, errors.New("the id do not nil")
	}

	filter := bson.M{"_id": attendanceID}
	var attendance attendancedomain.Attendance
	if err := collectionAttendance.FindOne(ctx, filter).Decode(&attendance); err != nil {
		return attendancedomain.Output{}, err
	}

	filterEmployee := bson.M{"_id": attendance.EmployeeID}
	var employee employeesdomain.Employee
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return attendancedomain.Output{}, err
	}

	employeeFullName := employee.FirstName + employee.LastName

	output := attendancedomain.Output{
		Attendance: attendance,
		Employee:   employeeFullName,
	}

	return output, nil
}

func (a attendanceRepository) GetOneByEmail(ctx context.Context, email string) (attendancedomain.Output, error) {
	collectionAttendance := a.database.Collection(a.collectionAttendance)
	collectionEmployee := a.database.Collection(a.collectionEmployee)

	if err := validate.IsNilEmailEmployee(email); err != nil {
		return attendancedomain.Output{}, err
	}

	filterEmployee := bson.M{"email": email}
	var employee employeesdomain.Employee
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return attendancedomain.Output{}, err
	}

	filter := bson.M{"employee_id": employee.ID}
	var attendance attendancedomain.Attendance
	if err := collectionAttendance.FindOne(ctx, filter).Decode(&attendance); err != nil {
		return attendancedomain.Output{}, err
	}

	employeeFullName := employee.FirstName + employee.LastName

	output := attendancedomain.Output{
		Attendance: attendance,
		Employee:   employeeFullName,
	}

	return output, nil
}

func (a attendanceRepository) GetAll(ctx context.Context) ([]attendancedomain.Output, error) {
	collectionAttendance := a.database.Collection(a.collectionAttendance)
	collectionEmployee := a.database.Collection(a.collectionEmployee)

	filter := bson.M{}
	cursor, err := collectionAttendance.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var attendances []attendancedomain.Output
	attendances = make([]attendancedomain.Output, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var attendance attendancedomain.Attendance
		if err = cursor.Decode(&attendance); err != nil {
			return nil, err
		}

		var employee employeesdomain.Employee
		filterEmployee := bson.M{"_id": attendance.EmployeeID}
		if err = collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
			return nil, err
		}

		employeeFullName := employee.FirstName + employee.LastName
		output := attendancedomain.Output{
			Attendance: attendance,
			Employee:   employeeFullName,
		}

		attendances = append(attendances, output)
	}

	return attendances, nil
}
