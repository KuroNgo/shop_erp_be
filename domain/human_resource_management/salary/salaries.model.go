package salary_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"time"
)

const (
	CollectionSalary = "salary"
)

// Salary represents the salary information of an employee.
type Salary struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RoleID       primitive.ObjectID `bson:"role_id" json:"role_id"`
	UnitCurrency string             `bson:"unit_currency" json:"unit_currency"`
	BaseSalary   float64            `bson:"base_salary" json:"base_salary"`
	Bonus        float64            `bson:"bonus" json:"bonus"`
	Deductions   float64            `bson:"deductions" json:"deductions"`
	NetSalary    float64            `bson:"net_salary" json:"net_salary"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	Role         string  `bson:"role" json:"role" exmaple:"Admin"`
	BaseSalary   float64 `bson:"base_salary" json:"base_salary" exmaple:"1500"`
	UnitCurrency string  `bson:"unit_currency" json:"unit_currency" exmaple:"USD"`
	Bonus        float64 `bson:"bonus" json:"bonus" exmaple:"200"`
	Deductions   float64 `bson:"deductions" json:"deductions" exmaple:"100"`
}

type Output struct {
	Salary Salary          `bson:"salary" json:"salary"`
	Role   roledomain.Role `bson:"role" json:"role"`
}
