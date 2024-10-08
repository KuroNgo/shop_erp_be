package departments_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	employees_domain "shop_erp_mono/domain/human_resource_management/employees"
	"time"
)

const (
	CollectionDepartment = "department"
)

// Department struct represents a department within the company.
type Department struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ManagerID   primitive.ObjectID `bson:"manager_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type Input struct {
	Name         string `bson:"name" example:"Human Resources"`
	Description  string `bson:"description" example:"Responsible for managing employee relations, recruitment, and company culture."`
	ManagerEmail string `bson:"manager_email" example:"admin@admin.com"`
}

// Output struct represents a department within the company.
type Output struct {
	Department Department                `bson:"department" json:"department"`
	Manager    employees_domain.Employee `bson:"manager" json:"manager"`
}
