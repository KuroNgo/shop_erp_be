package contracts_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Contract represents an employment contract of an employee.
type Contract struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID   primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	ContractType string             `bson:"contract_type" json:"contract_type"` // Example: "Full-Time", "Part-Time", "Temporary"
	StartDate    time.Time          `bson:"start_date" json:"start_date"`
	EndDate      time.Time          `bson:"end_date,omitempty" json:"end_date,omitempty"`
	Salary       float64            `bson:"salary" json:"salary"`
	Status       string             `bson:"status" json:"status"` // Example: "Active", "Expired"
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}