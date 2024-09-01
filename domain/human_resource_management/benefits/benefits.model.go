package benefits_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	employees_domain "shop_erp_mono/domain/human_resource_management/employees"
	"time"
)

const (
	CollectionBenefit = "benefit"
)

// Benefit represents the benefits an employee receives.
type Benefit struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID  primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	BenefitType string             `bson:"benefit_type" json:"benefit_type"` // Example: "Health Insurance", "Social Insurance", "Meal Allowance"
	Amount      float64            `bson:"amount" json:"amount"`
	StartDate   time.Time          `bson:"start_date" json:"start_date"`
	EndDate     time.Time          `bson:"end_date,omitempty" json:"end_date,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	EmployeeEmail string    `bson:"employee_email" json:"employee_email"`
	BenefitType   string    `bson:"benefit_type" json:"benefit_type"` // Example: "Health Insurance", "Social Insurance", "Meal Allowance"
	Amount        float64   `bson:"amount" json:"amount"`
	StartDate     time.Time `bson:"start_date" json:"start_date"`
	EndDate       time.Time `bson:"end_date,omitempty" json:"end_date,omitempty"`
}

type Output struct {
	Benefit  Benefit                   `bson:"benefit" json:"benefit"`
	Employee employees_domain.Employee `bson:"employee" json:"employee"`
}
