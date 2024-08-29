package departments_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionDepartment = "department"
)

// Department struct represents a department within the company.
type Department struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type Input struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
}

// Output struct represents a department within the company.
type Output struct {
	Department Department `bson:"department" json:"department"`
}
