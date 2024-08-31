package leave_request_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	leaverequestdomain "shop_erp_mono/domain/human_resource_management/leave_request"
	"shop_erp_mono/repository/human_resource_management/leave_request/validate"
	"time"
)

type leaveRequestRepository struct {
	database               *mongo.Database
	collectionLeaveRequest string
	collectionEmployee     string
}

func NewLeaveRequestRepository(db *mongo.Database, collectionLeaveRequest string, collectionEmployee string) leaverequestdomain.ILeaveRequestRepository {
	return &leaveRequestRepository{database: db, collectionLeaveRequest: collectionLeaveRequest, collectionEmployee: collectionEmployee}
}

func (l leaveRequestRepository) CreateOne(ctx context.Context, input *leaverequestdomain.Input) error {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)
	collectionEmployee := l.database.Collection(l.collectionEmployee)

	if err := validate.IsNilLeaveRequest(input); err != nil {
		return err
	}

	var employee employeesdomain.Employee
	filterEmployee := bson.M{"email": input.EmployeeEmail}
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return err
	}

	leaveRequest := leaverequestdomain.LeaveRequest{
		ID:         primitive.NewObjectID(),
		EmployeeID: employee.ID,
		LeaveType:  input.LeaveType,
		StartDate:  input.StartDate,
		EndDate:    input.EndDate,
		Status:     input.Status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, err := collectionLeaveRequest.InsertOne(ctx, leaveRequest)
	if err != nil {
		return err
	}

	return nil
}

func (l leaveRequestRepository) DeleteOne(ctx context.Context, id string) error {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)

	benefitID, _ := primitive.ObjectIDFromHex(id)
	if benefitID == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	filter := bson.M{"_id": benefitID}
	_, err := collectionLeaveRequest.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (l leaveRequestRepository) UpdateOne(ctx context.Context, input *leaverequestdomain.Input) error {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)
	collectionEmployee := l.database.Collection(l.collectionEmployee)

	if err := validate.IsNilLeaveRequest(input); err != nil {
		return err
	}

	filterEmployee := bson.M{"email": input.EmployeeEmail}
	var employee employeesdomain.Employee
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return err
	}

	filter := bson.M{"employee_id": input.EmployeeEmail}
	update := bson.M{"$set": bson.M{
		"leave_type": input.LeaveType,
		"start_date": input.StartDate,
		"end_date":   input.EndDate,
		"status":     input.Status,
	}}

	_, err := collectionLeaveRequest.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (l leaveRequestRepository) GetOneByID(ctx context.Context, id string) (leaverequestdomain.Output, error) {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)
	collectionEmployee := l.database.Collection(l.collectionEmployee)

	leaveRequestID, _ := primitive.ObjectIDFromHex(id)
	if leaveRequestID == primitive.NilObjectID {
		return leaverequestdomain.Output{}, errors.New("id do not nil")
	}

	filter := bson.M{"_id": leaveRequestID}
	var leaveRequest leaverequestdomain.LeaveRequest
	if err := collectionLeaveRequest.FindOne(ctx, filter).Decode(&leaveRequest); err != nil {
		return leaverequestdomain.Output{}, err
	}

	var employee employeesdomain.Employee
	filterEmployee := bson.M{"_id": leaveRequest.EmployeeID}
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return leaverequestdomain.Output{}, err
	}

	output := leaverequestdomain.Output{
		LeaveRequest: leaveRequest,
		NameEmployee: employee.LastName,
	}
	return output, nil
}

func (l leaveRequestRepository) GetOneByEmailEmployee(ctx context.Context, email string) (leaverequestdomain.Output, error) {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)
	collectionEmployee := l.database.Collection(l.collectionEmployee)

	var employee employeesdomain.Employee
	filterEmployee := bson.M{"email": email}
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return leaverequestdomain.Output{}, err
	}

	var leaveRequest leaverequestdomain.LeaveRequest
	filter := bson.M{"employee_id": employee.ID}
	if err := collectionLeaveRequest.FindOne(ctx, filter).Decode(&leaveRequest); err != nil {
		return leaverequestdomain.Output{}, err
	}

	output := leaverequestdomain.Output{
		LeaveRequest: leaveRequest,
		NameEmployee: employee.LastName,
	}

	return output, nil
}

func (l leaveRequestRepository) GetAll(ctx context.Context) ([]leaverequestdomain.Output, error) {
	collectionLeaveRequest := l.database.Collection(l.collectionLeaveRequest)
	collectionEmployee := l.database.Collection(l.collectionEmployee)

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

	var leaveRequests []leaverequestdomain.Output
	leaveRequests = make([]leaverequestdomain.Output, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var leaveRequest leaverequestdomain.LeaveRequest
		if err = cursor.Decode(&leaveRequest); err != nil {
			return nil, err
		}

		var employee employeesdomain.Employee
		filterEmployee := bson.M{"employee_id": leaveRequest.EmployeeID}
		if err = collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
			return nil, err
		}

		output := leaverequestdomain.Output{
			LeaveRequest: leaveRequest,
			NameEmployee: employee.LastName,
		}

		leaveRequests = append(leaveRequests, output)
	}

	return leaveRequests, nil
}
