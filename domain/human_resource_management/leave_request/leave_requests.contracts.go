package leave_request_domain

import "context"

type ILeaveRequestRepository interface {
	CreateOne(ctx context.Context, leaveRequest *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, leaveRequest *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByEmailEmployee(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}

type ILeaveRequestUseCase interface {
	CreateOne(ctx context.Context, leaveRequest *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, leaveRequest *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByEmailEmployee(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
