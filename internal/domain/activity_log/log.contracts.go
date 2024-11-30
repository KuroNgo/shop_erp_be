package activity_log_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ILogRepository interface {
	CreateOne(ctx context.Context, activityLog *ActivityLog) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, activityLog *ActivityLog) error
	GetByID(ctx context.Context, id primitive.ObjectID) (ActivityLog, error)
	GetByEmployeeID(ctx context.Context, idEmployee primitive.ObjectID) ([]ActivityLog, error)
	GetAll(ctx context.Context) ([]ActivityLog, error)
}

type ILogUseCase interface {
	CreateOne(ctx context.Context, activityLog *ActivityLog) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, activityLog *ActivityLog) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Response, error)
	GetByEmployeeID(ctx context.Context, idEmployee primitive.ObjectID) ([]Response, error)
	GetAll(ctx context.Context) ([]Response, error)
	LifeCycle(ctx context.Context) error
}
