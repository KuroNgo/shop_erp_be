package attendance_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
)

type attendanceRepository struct {
	database             *mongo.Database
	collectionAttendance string
}

func NewAttendanceRepository(db *mongo.Database, collectionAttendance string) attendancedomain.IAttendanceRepository {
	return &attendanceRepository{database: db, collectionAttendance: collectionAttendance}
}

func (a *attendanceRepository) CreateOne(ctx context.Context, attendance *attendancedomain.Attendance) error {
	collectionAttendance := a.database.Collection(a.collectionAttendance)

	_, err := collectionAttendance.InsertOne(ctx, attendance)
	if err != nil {
		return err
	}

	return nil
}

func (a *attendanceRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	collectionAttendance := a.database.Collection(a.collectionAttendance)

	if id == primitive.NilObjectID {
		return errors.New("error the id do not nil")
	}
	filter := bson.M{"_id": id}

	_, err := collectionAttendance.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (a *attendanceRepository) UpdateOne(ctx context.Context, attendance *attendancedomain.Attendance) error {
	collectionAttendance := a.database.Collection(a.collectionAttendance)

	if attendance.ID == primitive.NilObjectID {
		return errors.New("error the id do not nil")
	}
	filter := bson.M{"_id": attendance.ID}
	update := bson.M{"$set": attendance}

	_, err := collectionAttendance.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (a *attendanceRepository) GetOneByID(ctx context.Context, id primitive.ObjectID) (attendancedomain.Attendance, error) {
	collectionAttendance := a.database.Collection(a.collectionAttendance)

	if id == primitive.NilObjectID {
		return attendancedomain.Attendance{}, errors.New("the attendance ID do not null")
	}
	filter := bson.M{"_id": id}

	var attendance attendancedomain.Attendance
	if err := collectionAttendance.FindOne(ctx, filter).Decode(&attendance); err != nil {
		return attendancedomain.Attendance{}, err
	}

	return attendance, nil
}

func (a *attendanceRepository) GetOneByEmployeeID(ctx context.Context, idEmployee primitive.ObjectID) (attendancedomain.Attendance, error) {
	collectionAttendance := a.database.Collection(a.collectionAttendance)

	if idEmployee == primitive.NilObjectID {
		return attendancedomain.Attendance{}, errors.New("the employee ID do not null")
	}
	filter := bson.M{"employee_id": idEmployee}
	var attendance attendancedomain.Attendance
	if err := collectionAttendance.FindOne(ctx, filter).Decode(&attendance); err != nil {
		return attendancedomain.Attendance{}, err
	}

	return attendance, nil
}

func (a *attendanceRepository) GetAll(ctx context.Context) ([]attendancedomain.Attendance, error) {
	collectionAttendance := a.database.Collection(a.collectionAttendance)

	filter := bson.M{}
	cursor, err := collectionAttendance.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var attendances []attendancedomain.Attendance
	attendances = make([]attendancedomain.Attendance, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var attendance attendancedomain.Attendance
		if err = cursor.Decode(&attendance); err != nil {
			return nil, err
		}

		attendances = append(attendances, attendance)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}
