package leave_request_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	leaverequestdomain "shop_erp_mono/domain/human_resource_management/leave_request"
)

type leaveRequestRepository struct {
	database               *mongo.Database
	collectionLeaveRequest string
	collectionEmployee     string
}

func NewLeaveRequestRepository(db *mongo.Database, collectionLeaveRequest string, collectionEmployee string) leaverequestdomain.ILeaveRequestRepository {
	return &leaveRequestRepository{database: db, collectionLeaveRequest: collectionLeaveRequest, collectionEmployee: collectionEmployee}
}

func (l leaveRequestRepository) CreateOne(ctx context.Context, leaveRequest *leaverequestdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (l leaveRequestRepository) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (l leaveRequestRepository) UpdateOne(ctx context.Context, leaveRequest *leaverequestdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (l leaveRequestRepository) GetOneByID(ctx context.Context, id string) (leaverequestdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (l leaveRequestRepository) GetOneByName(ctx context.Context, name string) (leaverequestdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (l leaveRequestRepository) GetAllByEmployeeID(ctx context.Context, employeeID string) ([]leaverequestdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}
