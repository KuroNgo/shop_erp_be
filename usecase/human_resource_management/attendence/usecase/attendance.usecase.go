package attendance_usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/usecase/human_resource_management/attendence/validate"
	"time"
)

type attendanceUseCase struct {
	contextTimeout       time.Duration
	attendanceRepository attendancedomain.IAttendanceRepository
	employeeRepository   employeesdomain.IEmployeeRepository
	cache                *bigcache.BigCache
}

func NewAttendanceUseCase(contextTimeout time.Duration, attendanceRepository attendancedomain.IAttendanceRepository,
	employeeRepository employeesdomain.IEmployeeRepository, cacheTTL time.Duration) attendancedomain.IAttendanceUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &attendanceUseCase{contextTimeout: contextTimeout, cache: cache, attendanceRepository: attendanceRepository, employeeRepository: employeeRepository}
}

func (a *attendanceUseCase) CreateOne(ctx context.Context, input *attendancedomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	if err := validate.Attendance(input); err != nil {
		return err
	}

	employee, err := a.employeeRepository.GetByEmail(ctx, input.EmailEmployee)
	if err != nil {
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

	_ = a.cache.Delete("attendances")

	return a.attendanceRepository.CreateOne(ctx, &attendance)
}

func (a *attendanceUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	attendanceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_ = a.cache.Delete(id)
	_ = a.cache.Delete("attendances")

	return a.attendanceRepository.DeleteOne(ctx, attendanceID)
}

func (a *attendanceUseCase) UpdateOne(ctx context.Context, id string, input *attendancedomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	if input.CheckInTime.Before(input.CheckOutTime) {
		return errors.New("check-out time cannot be before check-in time")
	}

	if err := validate.Attendance(input); err != nil {
		return err
	}

	attendanceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	employee, err := a.employeeRepository.GetByEmail(ctx, input.EmailEmployee)
	if err != nil {
		return err
	}

	hoursWorked := input.CheckOutTime.Sub(input.CheckInTime)
	attendance := attendancedomain.Attendance{
		ID:           attendanceID,
		EmployeeID:   employee.ID,
		Date:         input.Date,
		CheckInTime:  input.CheckInTime,
		CheckOutTime: input.CheckOutTime,
		HoursWorked:  int8(hoursWorked.Hours()),
		UpdatedAt:    time.Now(),
	}

	_ = a.cache.Delete(id)
	_ = a.cache.Delete("attendances")

	return a.attendanceRepository.UpdateOne(ctx, &attendance)
}

func (a *attendanceUseCase) GetByID(ctx context.Context, id string) (attendancedomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	data, _ := a.cache.Get(id)
	if data != nil {
		var response attendancedomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return attendancedomain.Output{}, err
		}
		return response, nil
	}

	attendanceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	attendanceData, err := a.attendanceRepository.GetByID(ctx, attendanceID)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	employeeData, err := a.employeeRepository.GetByID(ctx, attendanceData.EmployeeID)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	output := attendancedomain.Output{
		Attendance: attendanceData,
		Employee:   employeeData,
	}

	data, err = json.Marshal(output)
	if err != nil {
		return attendancedomain.Output{}, err
	}
	err = a.cache.Set(id, data)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	return output, nil
}

func (a *attendanceUseCase) GetByEmail(ctx context.Context, email string) (attendancedomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	data, _ := a.cache.Get(email)
	if data != nil {
		var response attendancedomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return attendancedomain.Output{}, err
		}
		return response, nil
	}

	employeeData, err := a.employeeRepository.GetByEmail(ctx, email)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	attendanceData, err := a.attendanceRepository.GetByEmployeeID(ctx, employeeData.ID)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	output := attendancedomain.Output{
		Attendance: attendanceData,
		Employee:   employeeData,
	}

	data, err = json.Marshal(output)
	if err != nil {
		return attendancedomain.Output{}, err
	}
	err = a.cache.Set(email, data)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	return output, nil
}

func (a *attendanceUseCase) GetAll(ctx context.Context) ([]attendancedomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	data, _ := a.cache.Get("attendances")
	if data != nil {
		var response []attendancedomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	attendanceData, err := a.attendanceRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []attendancedomain.Output
	outputs = make([]attendancedomain.Output, 0, len(attendanceData))
	var employeesData []employeesdomain.Employee
	employeesData = make([]employeesdomain.Employee, 0, len(attendanceData))
	for _, i := range attendanceData {
		employeeData, err := a.employeeRepository.GetByID(ctx, i.EmployeeID)
		if err != nil {
			return nil, err
		}

		output := attendancedomain.Output{
			Attendance: i,
			Employee:   employeeData,
		}
		employeesData = append(employeesData, employeeData)
		outputs = append(outputs, output)
	}

	data, err = json.Marshal(outputs)
	if err != nil {
		return nil, err
	}
	err = a.cache.Set("attendances", data)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}
