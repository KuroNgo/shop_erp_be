package leave_request_usecase

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	leaverequestdomain "shop_erp_mono/domain/human_resource_management/leave_request"
	"shop_erp_mono/usecase/human_resource_management/leave_request/validate"
	"time"
)

type leaveRequestUseCase struct {
	contextTimeout         time.Duration
	leaveRequestRepository leaverequestdomain.ILeaveRequestRepository
	employeeRepository     employeesdomain.IEmployeeRepository
	cache                  *bigcache.BigCache
}

func NewLeaveRequestUseCase(contextTimeout time.Duration, leaveRequestRepository leaverequestdomain.ILeaveRequestRepository,
	employeeRepository employeesdomain.IEmployeeRepository, cacheTTL time.Duration) leaverequestdomain.ILeaveRequestUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &leaveRequestUseCase{contextTimeout: contextTimeout, cache: cache, leaveRequestRepository: leaveRequestRepository, employeeRepository: employeeRepository}
}

func (l *leaveRequestUseCase) CreateOne(ctx context.Context, input *leaverequestdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	if err := validate.LeaveRequest(input); err != nil {
		return err
	}

	employeeData, err := l.employeeRepository.GetByEmail(ctx, input.EmployeeEmail)
	if err != nil {
		return err
	}

	leaveRequest := &leaverequestdomain.LeaveRequest{
		ID:         primitive.NewObjectID(),
		EmployeeID: employeeData.ID,
		LeaveType:  input.LeaveType,
		StartDate:  input.StartDate,
		EndDate:    input.EndDate,
		Status:     input.Status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_ = l.cache.Delete("leaveRequests")

	return l.leaveRequestRepository.CreateOne(ctx, leaveRequest)
}

func (l *leaveRequestUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	leaveRequestID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_ = l.cache.Delete(id)
	_ = l.cache.Delete("leaveRequests")

	return l.leaveRequestRepository.DeleteOne(ctx, leaveRequestID)
}

func (l *leaveRequestUseCase) UpdateOne(ctx context.Context, id string, input *leaverequestdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	leaveRequestID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err = validate.LeaveRequest(input); err != nil {
		return err
	}

	employeeData, err := l.employeeRepository.GetByEmail(ctx, input.EmployeeEmail)
	if err != nil {
		return err
	}

	leaveRequest := &leaverequestdomain.LeaveRequest{
		ID:         primitive.NewObjectID(),
		EmployeeID: employeeData.ID,
		LeaveType:  input.LeaveType,
		StartDate:  input.StartDate,
		EndDate:    input.EndDate,
		Status:     input.Status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_ = l.cache.Delete(id)
	_ = l.cache.Delete("leaveRequests")

	return l.leaveRequestRepository.UpdateOne(ctx, leaveRequestID, leaveRequest)
}

func (l *leaveRequestUseCase) GetByID(ctx context.Context, id string) (leaverequestdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	data, _ := l.cache.Get(id)
	if data != nil {
		var response leaverequestdomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return leaverequestdomain.Output{}, err
		}
		return response, nil
	}

	leaveRequestID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	leaveRequestData, err := l.leaveRequestRepository.GetByID(ctx, leaveRequestID)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	employeeData, err := l.employeeRepository.GetByID(ctx, leaveRequestData.EmployeeID)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	output := leaverequestdomain.Output{
		LeaveRequest: leaveRequestData,
		Employee:     *employeeData,
	}

	data, err = json.Marshal(output)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	err = l.cache.Set(id, data)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	return output, nil
}

func (l *leaveRequestUseCase) GetByEmailEmployee(ctx context.Context, email string) (leaverequestdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	data, _ := l.cache.Get(email)
	if data != nil {
		var response leaverequestdomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return leaverequestdomain.Output{}, err
		}
		return response, nil
	}

	employeeData, err := l.employeeRepository.GetByEmail(ctx, email)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	leaveRequestData, err := l.leaveRequestRepository.GetByEmployeeID(ctx, employeeData.ID)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	output := leaverequestdomain.Output{
		LeaveRequest: leaveRequestData,
		Employee:     *employeeData,
	}

	data, err = json.Marshal(output)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	err = l.cache.Set(email, data)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	return output, nil
}

func (l *leaveRequestUseCase) GetAll(ctx context.Context) ([]leaverequestdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	data, _ := l.cache.Get("leaveRequests")
	if data != nil {
		var response []leaverequestdomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	leaveRequestData, err := l.leaveRequestRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []leaverequestdomain.Output
	outputs = make([]leaverequestdomain.Output, 0, len(leaveRequestData))
	for _, leaveRequest := range leaveRequestData {
		employeeData, err := l.employeeRepository.GetByID(ctx, leaveRequest.EmployeeID)
		if err != nil {
			return nil, err
		}

		output := leaverequestdomain.Output{
			LeaveRequest: leaveRequest,
			Employee:     *employeeData,
		}

		outputs = append(outputs, output)
	}

	data, err = json.Marshal(outputs)
	if err != nil {
		return nil, err
	}

	err = l.cache.Set("leaveRequests", data)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}
