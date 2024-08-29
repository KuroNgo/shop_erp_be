package attendance_domain

import "context"

type IAttendanceRepository interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByEmail(ctx context.Context, email string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}

type IAttendanceUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByEmail(ctx context.Context, email string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
