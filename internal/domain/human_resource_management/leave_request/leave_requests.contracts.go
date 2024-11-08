package leave_request_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ILeaveRequestRepository interface {
	CreateOne(ctx context.Context, leaveRequest *LeaveRequest) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, leaveRequest *LeaveRequest) error
	GetByID(ctx context.Context, id primitive.ObjectID) (LeaveRequest, error)
	GetByEmployeeID(ctx context.Context, employeeID primitive.ObjectID) (LeaveRequest, error)
	GetAll(ctx context.Context) ([]LeaveRequest, error)
}

type ILeaveRequestUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetByID(ctx context.Context, id string) (Output, error)
	GetByEmailEmployee(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
