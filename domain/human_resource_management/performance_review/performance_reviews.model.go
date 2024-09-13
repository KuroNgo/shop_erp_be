package performance_review_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	employees_domain "shop_erp_mono/domain/human_resource_management/employees"
	"time"
)

const (
	CollectionPerformanceReview = "performance_review"
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

type Input1 struct {
	EmployeeEmail    string    `bson:"employee" json:"employee"`
	ReviewerEmail    string    `bson:"reviewer" json:"reviewer"`
	ReviewDate       time.Time `bson:"review_date" json:"review_date"`
	PerformanceScore int       `bson:"performance_score" json:"performance_score"`
	Comments         string    `bson:"comments" json:"comments"`
}

type Input2 struct {
	EmployeeID       string    `bson:"employee_id" json:"employee_id"`
	ReviewerID       string    `bson:"reviewer_id" json:"reviewer_id"`
	ReviewDate       time.Time `bson:"review_date" json:"review_date"`
	PerformanceScore int       `bson:"performance_score" json:"performance_score"`
	Comments         string    `bson:"comments" json:"comments"`
}

type Output struct {
	PerformanceReview PerformanceReview         `bson:"performance_review"`
	Reviewer          employees_domain.Employee `bson:"reviewer" json:"reviewer"`
	Employee          employees_domain.Employee `bson:"employee" json:"employee"`
}
