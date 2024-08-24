package employees_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionEmployee = "employee"
)

// Employee struct represents an employee in the HR system.
type Employee struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	FirstName    string             `bson:"first_name"`
	LastName     string             `bson:"last_name"`
	Gender       string             `bson:"gender"`
	Email        string             `bson:"email"`
	Phone        string             `bson:"phone"`
	Address      string             `bson:"address"`
	AvatarURL    string             `bson:"avatar_url"`
	DateOfBirth  time.Time          `bson:"date_of_birth"`
	DayOfWork    time.Time          `bson:"date_of_work"`
	DepartmentID primitive.ObjectID `bson:"department_id"`
	RoleID       primitive.ObjectID `bson:"role_id"`
	SalaryID     primitive.ObjectID `bson:"salary"`
	IsActive     bool               `bson:"is_active"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type Input struct {
	FirstName   string    `bson:"first_name" example:"Ngô"`
	LastName    string    `bson:"last_name" example:"Hoài Phong"`
	Gender      string    `bson:"gender" example:"Nam"`
	Email       string    `bson:"email" example:"hoaiphong01012002@gmail.com"`
	Phone       string    `bson:"phone" example:"0329245971"`
	Address     string    `bson:"address" example:"Bình Thuận"`
	DateOfBirth time.Time `bson:"date_of_birth" example:"01/01/2002"`
	DayOfWork   time.Time `bson:"date_of_work" example:"30/8/2024"`
	Department  string    `bson:"department" example:"marketing"`
	Role        string    `bson:"role" example:"developer"`
	Salary      string    `bson:"salary" example:"21000000"`
}

type Output struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FirstName   string             `bson:"first_name"`
	LastName    string             `bson:"last_name"`
	Gender      string             `bson:"gender"`
	Email       string             `bson:"email"`
	Phone       string             `bson:"phone"`
	Address     string             `bson:"address"`
	DateOfBirth time.Time          `bson:"date_of_birth"`
	DayOfWork   time.Time          `bson:"date_of_work"`
	Department  string             `bson:"department"`
	Role        string             `bson:"role"`
	Salary      string             `bson:"salary"`
	IsActive    bool               `bson:"is_active"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
