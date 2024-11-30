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
	return a.activityLogRepository.CreateOne(ctx, activityLog)
}

func (a *activityLogUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	logID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return a.activityLogRepository.DeleteOne(ctx, logID)
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

	userData, err := a.userRepository.GetByID(ctx, data.UserID)
	if err != nil {
		return activitylogdomain.Response{}, err
	}

	employeeData, err := a.employeeRepository.GetByID(ctx, userData.EmployeeID)
	if err != nil {
		return activitylogdomain.Response{}, err
	}

	response := activitylogdomain.Response{
		ActivityLog: data,
		Username:    userData.Username,
		Employee:    employeeData.LastName + employeeData.FirstName,
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
	outputs = make([]activitylogdomain.Response, 0, len(activityData))
	for _, activity := range activityData {
		userData, err := a.userRepository.GetByID(ctx, activity.UserID)
		if err != nil {
			return nil, err
		}

		employeeData, err := a.employeeRepository.GetByID(ctx, userData.EmployeeID)
		if err != nil {
			return nil, err
		}

		response := activitylogdomain.Response{
			ActivityLog: activity,
			Username:    userData.Username,
			Employee:    employeeData.LastName + employeeData.FirstName,
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
	outputs = make([]activitylogdomain.Response, 0, len(activityData))
	for _, activity := range activityData {
		userData, err := a.userRepository.GetByID(ctx, activity.UserID)
		if err != nil {
			return nil, err
		}

		employeeData, err := a.employeeRepository.GetByID(ctx, userData.EmployeeID)
		if err != nil {
			return nil, err
		}

		response := activitylogdomain.Response{
			ActivityLog: activity,
			Username:    userData.Username,
			Employee:    employeeData.LastName + employeeData.FirstName,
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

func (a *activityLogUseCase) LifeCycle(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	return nil
}
