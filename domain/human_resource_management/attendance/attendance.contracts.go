package attendance_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAttendanceRepository interface {
	CreateOne(ctx context.Context, attendance *Attendance) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, attendance *Attendance) error
	GetOneByID(ctx context.Context, id primitive.ObjectID) (Attendance, error)
	GetOneByEmployeeID(ctx context.Context, idEmployee primitive.ObjectID) (Attendance, error)
	GetAll(ctx context.Context) ([]Attendance, error)
}

type IAttendanceUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByEmail(ctx context.Context, email string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
