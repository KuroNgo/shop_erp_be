package leave_request_usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	leaverequestdomain "shop_erp_mono/internal/domain/human_resource_management/leave_request"
	"shop_erp_mono/internal/usecase/human_resource_management/leave_request/validate"
	"shop_erp_mono/pkg/shared/mail/handles"
	"sync"
	"time"
)

type leaveRequestUseCase struct {
	contextTimeout         time.Duration
	leaveRequestRepository leaverequestdomain.ILeaveRequestRepository
	employeeRepository     employeesdomain.IEmployeeRepository
	cache                  *bigcache.BigCache
	client                 *mongo.Client
}

func NewLeaveRequestUseCase(contextTimeout time.Duration, leaveRequestRepository leaverequestdomain.ILeaveRequestRepository,
	employeeRepository employeesdomain.IEmployeeRepository, cacheTTL time.Duration, client *mongo.Client) leaverequestdomain.ILeaveRequestUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &leaveRequestUseCase{contextTimeout: contextTimeout, client: client, cache: cache, leaveRequestRepository: leaveRequestRepository, employeeRepository: employeeRepository}
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

	result, err := l.submitLeaveRequest(ctx, employeeData.ID, input)
	if err != nil {
		log.Printf("%s", result)
		return err
	}

	approvesData, err := l.employeeRepository.GetByEmail(ctx, input.ApprovesEmail)
	if err != nil {
		return err
	}

	leaveRequest := &leaverequestdomain.LeaveRequest{
		ID:          primitive.NewObjectID(),
		EmployeeID:  employeeData.ID,
		ApprovesID:  approvesData.ID,
		LeaveType:   input.LeaveType,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		RequestDays: time.Now(),
		Status:      input.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
		ID:         leaveRequestID,
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

func (l *leaveRequestUseCase) UpdateOneWithApproved(ctx context.Context, requestID string) error {
	err := l.approvedLeaveRequest(ctx, requestID)
	if err != nil {
		return err
	}

	return nil
}

func (l *leaveRequestUseCase) UpdateRemainingLeaveDays(ctx context.Context) error {
	const batchSize = 100 // Kích thước batch processing

	employees, err := l.employeeRepository.GetAll(ctx)
	if err != nil {
		return err
	}

	session, err := l.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	for i := 0; i < len(employees); i += batchSize {
		end := i + batchSize
		if end > len(employees) {
			end = len(employees)
		}
		batch := employees[i:end]

		// Transaction callback cho mỗi batch
		callback := func(sessionCtx mongo.SessionContext) (interface{}, error) {
			for _, employee := range batch {
				// Lấy dữ liệu yêu cầu nghỉ phép cho mỗi nhân viên
				leaveRequestData, err := l.leaveRequestRepository.GetByEmployeeID(sessionCtx, employee.ID)
				if err != nil {
					return nil, err
				}

				// Cập nhật số ngày nghỉ còn lại
				leaveRequestData.RemainingDays += leaveRequestData.TotalLeaveDays
				err = l.leaveRequestRepository.UpdateRemainingLeaveDays(sessionCtx, employee.ID, leaveRequestData.RemainingDays)
				if err != nil {
					return nil, err
				}
			}
			return nil, nil
		}

		// Chạy transaction cho batch hiện tại
		if _, err := session.WithTransaction(ctx, callback); err != nil {
			return err // Trả về lỗi ngay khi transaction fail
		}
	}

	return nil
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

func calculatedRequestDays(startDate, endDate time.Time) (int, error) {
	if endDate.Before(startDate) {
		return 0, errors.New("end date cannot before start date")
	}

	days := int(endDate.Sub(startDate).Hours() / 24)
	return days + 1, nil
}

func (uc *leaveRequestUseCase) submitLeaveRequest(ctx context.Context, employeeID primitive.ObjectID, input *leaverequestdomain.Input) (string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	requestedDays, err := calculatedRequestDays(input.StartDate, input.EndDate)
	if err != nil {
		return "", err
	}

	remainingLeaveDays, err := uc.leaveRequestRepository.GetRemainingLeaveDays(ctx, employeeID)
	if err != nil {
		return "", err
	}

	employeeData, err := uc.employeeRepository.GetByID(ctx, employeeID)
	if err != nil {
		return "", err
	}

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	// Check leave balance
	wg.Add(1)
	go func() {
		defer wg.Done()
		if requestedDays > remainingLeaveDays {
			emailData := handles.EmailData{
				FullName: employeeData.FirstName + employeeData.LastName,
				HREmail:  "kurongo.test@gmail.com",
			}
			if err := handles.SendEmail(&emailData, employeeData.Email, "leave_request.warning_remaining_days.html"); err != nil {
				errCh <- err
			} else {
				errCh <- errors.New("insufficient leave balance")
			}
			cancel() // Cancel other goroutines if leave balance is insufficient
		}
	}()

	// Send email based on leave type
	wg.Add(1)
	go func() {
		defer wg.Done()
		leaveTypeMapping := map[string]string{
			"Sick Leave":   "sick",
			"Annual Leave": "annual",
			"Maternity":    "maternity",
			"Unpaid Leave": "unpaid",
		}
		leaveType, exists := leaveTypeMapping[input.LeaveType]
		if exists {
			emailData := handles.EmailData{
				FullName:  employeeData.FirstName + employeeData.LastName,
				HREmail:   "kurongo.test@gmail.com",
				LeaveType: leaveType,
			}
			if err := handles.SendEmail(&emailData, employeeData.Email, "leave_request.leave_type.html"); err != nil {
				errCh <- err
				cancel() // Cancel other goroutines if email sending fails
			}
		}
	}()

	// Close the error channel once all goroutines are done
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Use select to handle errors as soon as they occur
	for {
		select {
		case err, ok := <-errCh:
			if ok && err != nil {
				return "", err
			}
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			if len(errCh) == 0 {
				return "leave request submitted successfully", nil
			}
		}
	}
}

func (l *leaveRequestUseCase) approvedLeaveRequest(ctx context.Context, requestID string) error {
	idRequest, err := primitive.ObjectIDFromHex(requestID)
	if err != nil {
		return err
	}

	session, err := l.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	callback := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		requestData, err := l.leaveRequestRepository.GetByID(sessionCtx, idRequest)
		if err != nil {
			return nil, err
		}

		remainingDays, err := l.leaveRequestRepository.GetRemainingLeaveDays(sessionCtx, requestData.EmployeeID)
		if err != nil {
			return nil, err

		}

		if requestData.RemainingDays > remainingDays {
			return nil, errors.New("insufficient leave balance to approve this request")
		}

		err = l.leaveRequestRepository.UpdateRemainingLeaveDays(sessionCtx,
			requestData.EmployeeID, remainingDays-requestData.RemainingDays)
		if err != nil {
			return nil, err
		}

		err = l.leaveRequestRepository.UpdateStatus(sessionCtx, requestData.EmployeeID, "Approved")
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

	return session.CommitTransaction(ctx)
}
