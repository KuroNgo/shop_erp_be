package salary_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionSalary = "salary"
)

// Salary represents the salary information of an employee.
type Salary struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	BaseSalary float64            `bson:"base_salary" json:"base_salary"`
	Bonus      float64            `bson:"bonus" json:"bonus"`
	Deductions float64            `bson:"deductions" json:"deductions"`
	NetSalary  float64            `bson:"net_salary" json:"net_salary"`
	PayDate    time.Time          `bson:"pay_date" json:"pay_date"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
