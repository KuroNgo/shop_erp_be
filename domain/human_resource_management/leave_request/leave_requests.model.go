package leave_request_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionLeaveRequest = "leave_request"
)

// LeaveRequest represents a leave request by an employee.
type LeaveRequest struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	LeaveType  string             `bson:"leave_type" json:"leave_type"` // Example: "Sick Leave", "Annual Leave", "Unpaid Leave"
	StartDate  time.Time          `bson:"start_date" json:"start_date"`
	EndDate    time.Time          `bson:"end_date" json:"end_date"`
	Status     string             `bson:"status" json:"status"` // Example: "Approved", "Pending", "Rejected"
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
