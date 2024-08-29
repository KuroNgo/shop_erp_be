package role_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionRole = "role"
)

// Role struct represents a role or job role.
type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type Input struct {
	Title       string `bson:"title" example:"Admin"`
	Description string `bson:"description" example:"The admin role has full access and control over the system."`
}

// Output struct represents a role or job role.
type Output struct {
	Role Role `bson:"role"`
}
