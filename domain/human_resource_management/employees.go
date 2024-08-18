package human_resource_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Employee struct represents an employee in the HR system.
type Employee struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	FirstName    string             `bson:"first_name"`
	LastName     string             `bson:"last_name"`
	Email        string             `bson:"email"`
	Phone        string             `bson:"phone"`
	Address      string             `bson:"address"`
	DateOfBirth  time.Time          `bson:"date_of_birth"`
	HireDate     time.Time          `bson:"hire_date"`
	DepartmentID primitive.ObjectID `bson:"department_id"`
	RoleID       primitive.ObjectID `bson:"role_id"`
	Salary       float64            `bson:"salary"`
	IsActive     bool               `bson:"is_active"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}
