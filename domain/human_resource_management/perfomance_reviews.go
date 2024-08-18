package human_resource_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// PerformanceReview represents performance evaluations of an employee.
type PerformanceReview struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID       primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	ReviewDate       time.Time          `bson:"review_date" json:"review_date"`
	ReviewerID       primitive.ObjectID `bson:"reviewer_id" json:"reviewer_id"`
	PerformanceScore int                `bson:"performance_score" json:"performance_score"`
	Comments         string             `bson:"comments" json:"comments"`
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at"`
}
