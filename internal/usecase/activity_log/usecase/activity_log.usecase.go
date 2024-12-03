package activity_log_usecase

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	activitylogdomain "shop_erp_mono/internal/domain/activity_log"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	"time"
)

type activityLogUseCase struct {
	contextTimeout        time.Duration
	activityLogRepository activitylogdomain.ILogRepository
	userRepository        userdomain.IUserRepository
	employeeRepository    employeesdomain.IEmployeeRepository
	client                *mongo.Client
	cache                 *bigcache.BigCache
}

func NewActivityLogUseCase(contextTimeout time.Duration, activityLogRepository activitylogdomain.ILogRepository,
	employeeRepository employeesdomain.IEmployeeRepository, userRepository userdomain.IUserRepository,
	cacheTTL time.Duration, client *mongo.Client) activitylogdomain.ILogUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &activityLogUseCase{contextTimeout: contextTimeout, cache: cache, client: client,
		activityLogRepository: activityLogRepository, employeeRepository: employeeRepository, userRepository: userRepository}
}

func (a *activityLogUseCase) CreateOne(ctx context.Context, activityLog *activitylogdomain.ActivityLog) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err := a.activityLogRepository.CreateOne(ctx, activityLog)
	if err != nil {
		return err
	}

	return nil
}

func (a *activityLogUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	logID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = a.activityLogRepository.DeleteOne(ctx, logID)
	if err != nil {
		return err
	}

	return nil
}

func (a *activityLogUseCase) GetByID(ctx context.Context, id string) (activitylogdomain.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	logID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return activitylogdomain.Response{}, err
	}

	data, err := a.activityLogRepository.GetByID(ctx, logID)
	if err != nil {
		return activitylogdomain.Response{}, err
	}

	var users []string
	var employees []string
	for _, userID := range data.UserID {
		userData, err := a.userRepository.GetByID(ctx, userID)
		if err != nil {
			return activitylogdomain.Response{}, err
		}

		employeeData, err := a.employeeRepository.GetByID(ctx, userData.EmployeeID)
		if err != nil {
			return activitylogdomain.Response{}, err
		}

		users = append(users, userData.Username)
		employees = append(employees, employeeData.FirstName+employeeData.LastName)
	}

	response := activitylogdomain.Response{
		ActivityLog: data,
		Username:    users,
		Employee:    employees,
	}

	return response, nil
}

func (a *activityLogUseCase) GetByEmployeeID(ctx context.Context, idEmployee string) ([]activitylogdomain.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	data, err := a.cache.Get(idEmployee)
	if err != nil {
		log.Printf("failed to get departments cache: %v", err)
	}

	if data != nil {
		var response []activitylogdomain.Response
		err = json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	id, err := primitive.ObjectIDFromHex(idEmployee)
	if err != nil {
		return nil, err
	}

	activityData, err := a.activityLogRepository.GetByEmployeeID(ctx, id)
	if err != nil {
		return nil, err
	}

	var outputs []activitylogdomain.Response
	var users []string
	var employees []string
	outputs = make([]activitylogdomain.Response, 0, len(activityData))
	for _, activity := range activityData {
		for _, userID := range activity.UserID {
			userData, err := a.userRepository.GetByID(ctx, userID)
			if err != nil {
				return nil, err
			}

			employeeData, err := a.employeeRepository.GetByID(ctx, userData.EmployeeID)
			if err != nil {
				return nil, err
			}

			users = append(users, userData.Username)
			employees = append(employees, employeeData.FirstName+employeeData.LastName)
		}

		response := activitylogdomain.Response{
			ActivityLog: activity,
			Username:    users,
			Employee:    employees,
		}

		outputs = append(outputs, response)
	}

	data, err = json.Marshal(outputs)
	if err != nil {
		return nil, err
	}

	err = a.cache.Set("departments_deleted", data)
	if err != nil {
		log.Printf("failed to set departments cache: %v", err)
	}

	return outputs, nil
}

func (a *activityLogUseCase) GetAll(ctx context.Context) ([]activitylogdomain.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	data, err := a.cache.Get("logs")
	if err != nil {
		log.Printf("failed to get departments cache: %v", err)
	}

	if data != nil {
		var response []activitylogdomain.Response
		err = json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	activityData, err := a.activityLogRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []activitylogdomain.Response
	var users []string
	var employees []string
	outputs = make([]activitylogdomain.Response, 0, len(activityData))
	for _, activity := range activityData {
		for _, userID := range activity.UserID {
			userData, err := a.userRepository.GetByID(ctx, userID)
			if err != nil {
				return nil, err
			}

			employeeData, err := a.employeeRepository.GetByID(ctx, userData.EmployeeID)
			if err != nil {
				return nil, err
			}

			users = append(users, userData.Username)
			employees = append(employees, employeeData.FirstName+employeeData.LastName)
		}

		response := activitylogdomain.Response{
			ActivityLog: activity,
			Username:    users,
			Employee:    employees,
		}

		outputs = append(outputs, response)
	}

	data, err = json.Marshal(outputs)
	if err != nil {
		return nil, err
	}

	err = a.cache.Set("departments_deleted", data)
	if err != nil {
		log.Printf("failed to set departments cache: %v", err)
	}

	return outputs, nil
}

func (a *activityLogUseCase) PrintLog(ctx context.Context, mos ...int) error {
	//TODO implement me
	panic("implement me")
}

func (a *activityLogUseCase) LifeCycle(ctx context.Context) error {
	return nil
}
