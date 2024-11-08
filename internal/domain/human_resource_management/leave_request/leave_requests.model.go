package leave_request_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	employees_domain "shop_erp_mono/internal/domain/human_resource_management/employees"
	"time"
)

const (
	CollectionLeaveRequest = "leave_request"
)

// LeaveRequest represents a leave request by an employee.
type LeaveRequest struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID    primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	ApprovesID    primitive.ObjectID `bson:"approves_id" json:"approves_id"`
	LeaveType     string             `bson:"leave_type" json:"leave_type"` // Example: "Sick Leave", "Annual Leave", "Unpaid Leave"
	StartDate     time.Time          `bson:"start_date" json:"start_date"`
	EndDate       time.Time          `bson:"end_date" json:"end_date"`
	RemainingDays int                `bson:"remaining_days" json:"remaining_days"`
	Status        string             `bson:"status" json:"status"` // Example: "Approved", "Pending", "Rejected"
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	EmployeeEmail string    `bson:"employee" json:"employee" example:"admin@admin.com"`
	ApprovesEmail string    `bson:"approves" json:"approves"`
	LeaveType     string    `bson:"leave_type" json:"leave_type" example:"Sick Leave"` // Example: "Sick Leave", "Annual Leave", "Unpaid Leave"
	StartDate     time.Time `bson:"start_date" json:"start_date" example:"20/07/2024"`
	EndDate       time.Time `bson:"end_date" json:"end_date" example:"20/07/2024"`
	RemainingDays int       `bson:"remaining_days" json:"remaining_days"`
	Status        string    `bson:"status" json:"status" example:"Approved"` // Example: "Approved", "Pending", "Rejected"
}

type Output struct {
	LeaveRequest LeaveRequest              `bson:"leave_request"`
	Employee     employees_domain.Employee `bson:"name_employee"`
}
