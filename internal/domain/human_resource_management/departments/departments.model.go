package departments_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	employees_domain "shop_erp_mono/internal/domain/human_resource_management/employees"
	"time"
)

const (
	CollectionDepartment = "department"
)

type Department struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ManagerID   primitive.ObjectID `bson:"manager_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	ParentID    primitive.ObjectID `bson:"parent_id,omitempty"`
	Level       int                `bson:"level"`
	Status      string             `bson:"status"`
	Enable      int                `bson:"enable"`
	WhoDeleted  primitive.ObjectID `bson:"who_deleted"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type Input struct {
	Name         string             `bson:"name" example:"Human Resources"`
	Level        int                `bson:"level" example:"1"`
	ParentID     primitive.ObjectID `bson:"parent_id,omitempty"`
	Description  string             `bson:"description" example:"Responsible for managing employee relations, recruitment, and company culture."`
	ManagerEmail string             `bson:"manager_email" example:"admin@admin.com"`
}

type Output struct {
	Department    Department                `json:"department"`
	Manager       employees_domain.Employee `json:"manager"`
	CountEmployee int64                     `json:"count_employee"`
}
