package salary_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionSalary = "salary"
)

type Salary struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID   primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	UnitCurrency string             `bson:"unit_currency" json:"unit_currency"`
	BaseSalary   float64            `bson:"base_salary" json:"base_salary"`
	Bonus        float64            `bson:"bonus" json:"bonus"`
	Deductions   float64            `bson:"deductions" json:"deductions"`
	NetSalary    float64            `bson:"net_salary" json:"net_salary"`
	Status       string             `bson:"status" json:"status"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	EmployeeID   primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	BaseSalary   float64            `bson:"base_salary" json:"base_salary" exmaple:"1500"`
	UnitCurrency string             `bson:"unit_currency" json:"unit_currency" exmaple:"USD"`
	Bonus        float64            `bson:"bonus" json:"bonus" exmaple:"200"`
	Deductions   float64            `bson:"deductions" json:"deductions" exmaple:"100"`
	Status       string             `bson:"status" json:"status"`
}

type Output struct {
	Salary       Salary `bson:"salary" json:"salary"`
	EmployeeName string `bson:"employee" json:"employee"`
}
