package leave_request_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	leaverequestdomain "shop_erp_mono/domain/human_resource_management/leave_request"
)

type leaveRequestRepository struct {
	database               *mongo.Database
	collectionLeaveRequest string
}

func NewLeaveRequestRepository(db *mongo.Database, collectionLeaveRequest string) leaverequestdomain.ILeaveRequestRepository {
	return &leaveRequestRepository{database: db, collectionLeaveRequest: collectionLeaveRequest}
}

func (l *leaveRequestRepository) CreateOne(ctx context.Context, leaveRequest *leaverequestdomain.LeaveRequest) error {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)

	_, err := collectionLeaveRequest.InsertOne(ctx, leaveRequest)
	if err != nil {
		return err
	}

	return nil
}

func (l *leaveRequestRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)

	if id == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	filter := bson.M{"_id": id}
	_, err := collectionLeaveRequest.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (l *leaveRequestRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, leaveRequest *leaverequestdomain.LeaveRequest) error {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"leave_type": leaveRequest.LeaveType,
		"start_date": leaveRequest.StartDate,
		"end_date":   leaveRequest.EndDate,
		"status":     leaveRequest.Status,
	}}

	_, err := collectionLeaveRequest.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (l *leaveRequestRepository) GetOneByID(ctx context.Context, id primitive.ObjectID) (leaverequestdomain.LeaveRequest, error) {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)

	if id == primitive.NilObjectID {
		return leaverequestdomain.LeaveRequest{}, errors.New("id do not nil")
	}

	filter := bson.M{"_id": id}
	var leaveRequest leaverequestdomain.LeaveRequest
	if err := collectionLeaveRequest.FindOne(ctx, filter).Decode(&leaveRequest); err != nil {
		return leaverequestdomain.LeaveRequest{}, err
	}

	return leaveRequest, nil
}

func (l *leaveRequestRepository) GetOneByEmployeeID(ctx context.Context, employeeID primitive.ObjectID) (leaverequestdomain.LeaveRequest, error) {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)

	var leaveRequest leaverequestdomain.LeaveRequest
	filter := bson.M{"employee_id": employeeID}
	if err := collectionLeaveRequest.FindOne(ctx, filter).Decode(&leaveRequest); err != nil {
		return leaverequestdomain.LeaveRequest{}, err
	}

	return leaveRequest, nil
}

func (l *leaveRequestRepository) GetAll(ctx context.Context) ([]leaverequestdomain.LeaveRequest, error) {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)

	filter := bson.M{}
	cursor, err := collectionLeaveRequest.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var leaveRequests []leaverequestdomain.LeaveRequest
	leaveRequests = make([]leaverequestdomain.LeaveRequest, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var leaveRequest leaverequestdomain.LeaveRequest
		if err = cursor.Decode(&leaveRequest); err != nil {
			return nil, err
		}

		leaveRequests = append(leaveRequests, leaveRequest)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return leaveRequests, nil
}
