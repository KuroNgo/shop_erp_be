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
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ManagerID   primitive.ObjectID `bson:"manager_id" json:"manager_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	ParentID    primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id"`
	Level       int                `bson:"level" json:"level"`
	Status      string             `bson:"status" json:"status"`
	WhoDeleted  primitive.ObjectID `bson:"who_deleted" json:"who_deleted"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
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
