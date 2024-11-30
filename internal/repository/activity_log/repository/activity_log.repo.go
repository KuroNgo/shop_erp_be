package log_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	activitylogdomain "shop_erp_mono/internal/domain/activity_log"
)

type logRepository struct {
	logCollection string
	database      *mongo.Database
}

func NewLogRepository(logCollection string, database *mongo.Database) activitylogdomain.ILogRepository {
	return &logRepository{logCollection: logCollection, database: database}
}

func (l *logRepository) CreateOne(ctx context.Context, activityLog *activitylogdomain.ActivityLog) error {
	logCollection := l.database.Collection(l.logCollection)

	_, err := logCollection.InsertOne(ctx, activityLog)
	if err != nil {
		return err
	}

	return nil
}

func (l *logRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	logCollection := l.database.Collection(l.logCollection)

	filter := bson.M{"_id": id}
	_, err := logCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (l *logRepository) GetByID(ctx context.Context, id primitive.ObjectID) (activitylogdomain.ActivityLog, error) {
	logCollection := l.database.Collection(l.logCollection)

	filter := bson.M{"_id": id}
	var log activitylogdomain.ActivityLog
	err := logCollection.FindOne(ctx, filter).Decode(&log)
	if err != nil {
		return activitylogdomain.ActivityLog{}, err
	}

	return log, nil
}

func (l *logRepository) GetByEmployeeID(ctx context.Context, idEmployee primitive.ObjectID) ([]activitylogdomain.ActivityLog, error) {
	logCollection := l.database.Collection(l.logCollection)

	filter := bson.M{"employee_id": idEmployee}
	cursor, err := logCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var logs []activitylogdomain.ActivityLog
	logs = make([]activitylogdomain.ActivityLog, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var log activitylogdomain.ActivityLog
		if err = cursor.Decode(&log); err != nil {
			return nil, errors.New("error decoding activity log information from database")
		}

		logs = append(logs, log)
	}

	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func (l *logRepository) GetAll(ctx context.Context) ([]activitylogdomain.ActivityLog, error) {
	logCollection := l.database.Collection(l.logCollection)

	filter := bson.M{}
	cursor, err := logCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var logs []activitylogdomain.ActivityLog
	logs = make([]activitylogdomain.ActivityLog, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var log activitylogdomain.ActivityLog
		if err = cursor.Decode(&log); err != nil {
			return nil, errors.New("error decoding activity log information from database")
		}

		logs = append(logs, log)
	}

	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}
