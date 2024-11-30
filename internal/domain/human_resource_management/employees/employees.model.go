package employees_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionEmployee = "employee"
)

type Employee struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DepartmentID primitive.ObjectID `bson:"department_id" json:"department_id"`
	RoleID       primitive.ObjectID `bson:"role_id" json:"role_id"`
	FirstName    string             `bson:"first_name" json:"first_name"`
	LastName     string             `bson:"last_name" json:"last_name"`
	Gender       string             `bson:"gender" json:"gender"`
	Email        string             `bson:"email" json:"email"`
	Phone        string             `bson:"phone" json:"phone"`
	Address      string             `bson:"address" json:"address"`
	AvatarURL    string             `bson:"avatar_url" json:"avatar_url"`
	DateOfBirth  time.Time          `bson:"date_of_birth" json:"date_of_birth"`
	DayOfWork    time.Time          `bson:"date_of_work" json:"day_of_work"`
	StartDate    time.Time          `bson:"start_date" json:"start_date"`
	EndDate      time.Time          `bson:"end_date" json:"end_date"`
	Active       string             `bson:"active" json:"active"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	FirstName   string    `bson:"first_name" example:"Ngô"`
	LastName    string    `bson:"last_name" example:"Hoài Phong"`
	Gender      string    `bson:"gender" example:"Nam"`
	Email       string    `bson:"email" example:"hoaiphong01012002@gmail.com"`
	Phone       string    `bson:"phone" example:"0329245971"`
	Address     string    `bson:"address" example:"Bình Thuận"`
	AvatarURL   string    `bson:"avatar_url"`
	DateOfBirth time.Time `bson:"date_of_birth" json:"date_of_birth" example:"2002-01-01T00:00:00Z"`
	DayOfWork   time.Time `bson:"day_of_work" json:"day_of_work" example:"2024-08-30T00:00:00Z"`
	Department  string    `bson:"department" example:"marketing"`
	Role        string    `bson:"role" example:"developer"`
}

type Output struct {
	Employee   Employee `bson:"employee" json:"employee"`
	Department string   `bson:"department" json:"department"`
	Role       string   `bson:"role" json:"role"`
}
