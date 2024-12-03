package activity_log_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ILogRepository interface {
	CreateOne(ctx context.Context, activityLog *ActivityLog) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetByID(ctx context.Context, id primitive.ObjectID) (ActivityLog, error)
	GetByEmployeeID(ctx context.Context, idEmployee primitive.ObjectID) ([]ActivityLog, error)
	GetAll(ctx context.Context) ([]ActivityLog, error)
}

type ILogUseCase interface {
	CreateOne(ctx context.Context, activityLog *ActivityLog) error
	DeleteOne(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (Response, error)
	GetByEmployeeID(ctx context.Context, idEmployee string) ([]Response, error)
	GetAll(ctx context.Context) ([]Response, error)
	PrintLog(ctx context.Context, mos ...int) error
	LifeCycle(ctx context.Context) error
}
