package employees_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
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
}

type Input struct {
	FirstName   string    `bson:"first_name" example:"Ngô"`
	LastName    string    `bson:"last_name" example:"Hoài Phong"`
	Gender      string    `bson:"gender" example:"Nam"`
	Email       string    `bson:"email" example:"hoaiphong01012002@gmail.com"`
	Phone       string    `bson:"phone" example:"0329245971"`
	Address     string    `bson:"address" example:"Bình Thuận"`
	AvatarURL   string    `bson:"avatar_url"`
	DateOfBirth time.Time `bson:"date_of_birth" example:"01/01/2002"`
	DayOfWork   time.Time `bson:"date_of_work" example:"30/8/2024"`
	Department  string    `bson:"department" example:"marketing"`
	Role        string    `bson:"role" example:"developer"`
}

type Output struct {
	Employee       Employee            `bson:"employee"`
	DepartmentName string              `bson:"department"`
	RoleID         roledomain.Role     `bson:"role"`
	Salary         salarydomain.Salary `bson:"salary"`
}
