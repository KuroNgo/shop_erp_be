package attendance_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/repository/human_resource_management/attendance/validate"
	"time"
)

type attendanceUseCase struct {
	contextTimeout       time.Duration
	attendanceRepository attendancedomain.IAttendanceRepository
	employeeRepository   employeesdomain.IEmployeeRepository
}

func NewAttendanceUseCase(contextTimeout time.Duration, attendanceRepository attendancedomain.IAttendanceRepository, employeeRepository employeesdomain.IEmployeeRepository) attendancedomain.IAttendanceUseCase {
	return &attendanceUseCase{contextTimeout: contextTimeout, attendanceRepository: attendanceRepository, employeeRepository: employeeRepository}
}

func (a *attendanceUseCase) CreateOne(ctx context.Context, input *attendancedomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	if err := validate.IsNilAttendance(input); err != nil {
		return err
	}

	employee, err := a.employeeRepository.GetOneByEmail(ctx, input.EmailEmployee)
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

	return a.attendanceRepository.CreateOne(ctx, &attendance)
}

func (a *attendanceUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	attendanceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return a.attendanceRepository.DeleteOne(ctx, attendanceID)
}

func (a *attendanceUseCase) UpdateOne(ctx context.Context, id string, input *attendancedomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	if err := validate.IsNilAttendance(input); err != nil {
		return err
	}

	attendanceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	employee, err := a.employeeRepository.GetOneByEmail(ctx, input.EmailEmployee)
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

	return a.attendanceRepository.UpdateOne(ctx, &attendance)
}

func (a *attendanceUseCase) GetOneByID(ctx context.Context, id string) (attendancedomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	attendanceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	attendanceData, err := a.attendanceRepository.GetOneByID(ctx, attendanceID)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	employeeData, err := a.employeeRepository.GetOneByID(ctx, attendanceData.EmployeeID)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	output := attendancedomain.Output{
		Attendance: attendanceData,
		Employee:   employeeData,
	}

	return output, nil
}

func (a *attendanceUseCase) GetOneByEmail(ctx context.Context, email string) (attendancedomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	employeeData, err := a.employeeRepository.GetOneByEmail(ctx, email)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	attendanceData, err := a.attendanceRepository.GetOneByEmployeeID(ctx, employeeData.ID)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	output := attendancedomain.Output{
		Attendance: attendanceData,
		Employee:   employeeData,
	}

	return output, nil
}

func (a *attendanceUseCase) GetAll(ctx context.Context) ([]attendancedomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	attendanceData, err := a.attendanceRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []attendancedomain.Output
	outputs = make([]attendancedomain.Output, 0, len(attendanceData))
	var employeesData []employeesdomain.Employee
	employeesData = make([]employeesdomain.Employee, 0, len(attendanceData))
	for _, i := range attendanceData {
		employeeData, err := a.employeeRepository.GetOneByID(ctx, i.EmployeeID)
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

	return outputs, nil
}
