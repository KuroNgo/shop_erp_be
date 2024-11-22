package attendance_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAttendanceRepository interface {
	CreateOne(ctx context.Context, attendance *Attendance) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, attendance *Attendance) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Attendance, error)
	GetByStatus(ctx context.Context, status string) ([]Attendance, error)
	GetByEmployeeID(ctx context.Context, idEmployee primitive.ObjectID) (Attendance, error)
	GetAll(ctx context.Context) ([]Attendance, error)
}

type IAttendanceUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetByID(ctx context.Context, id string) (Output, error)
	GetByStatus(ctx context.Context, status string) ([]Output, error)
	GetByEmail(ctx context.Context, email string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
