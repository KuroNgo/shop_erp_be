package activity_log_usecase

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	activitylogdomain "shop_erp_mono/internal/domain/activity_log"
	employees_domain "shop_erp_mono/internal/domain/human_resource_management/employees"
	"time"
)

type activityLogUseCase struct {
	contextTimeout        time.Duration
	activityLogRepository activitylogdomain.ILogRepository
	employeeRepository    employees_domain.IEmployeeRepository
	client                *mongo.Client
	cache                 *bigcache.BigCache
}

func NewActivityLogUseCase(contextTimeout time.Duration, activityLogRepository activitylogdomain.ILogRepository,
	employeeRepository employees_domain.IEmployeeRepository, cacheTTL time.Duration, client *mongo.Client) activitylogdomain.ILogUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &activityLogUseCase{contextTimeout: contextTimeout, cache: cache, client: client,
		activityLogRepository: activityLogRepository, employeeRepository: employeeRepository}
}

func (a *activityLogUseCase) CreateOne(ctx context.Context, activityLog *activitylogdomain.ActivityLog) error {
	//TODO implement me
	panic("implement me")
}

func (a *activityLogUseCase) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (a *activityLogUseCase) UpdateOne(ctx context.Context, activityLog *activitylogdomain.ActivityLog) error {
	//TODO implement me
	panic("implement me")
}

func (a *activityLogUseCase) GetByID(ctx context.Context, id primitive.ObjectID) (activitylogdomain.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (a *activityLogUseCase) GetByEmployeeID(ctx context.Context, idEmployee primitive.ObjectID) ([]activitylogdomain.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (a *activityLogUseCase) GetAll(ctx context.Context) ([]activitylogdomain.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (a *activityLogUseCase) LifeCycle(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
