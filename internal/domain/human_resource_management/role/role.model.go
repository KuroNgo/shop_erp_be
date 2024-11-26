package role_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionRole = "role"
)

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Level       int                `bson:"level"`
	Status      string             `bson:"status"`
	Enable      int                `bson:"enable"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type Input struct {
	Name        string `bson:"name" example:"Admin"`
	Description string `bson:"description" example:"The admin role has full access and control over the system."`
	Level       int    `bson:"level"  example:"1"`
	Enable      int    `bson:"enable" example:"1"`
}

type Output struct {
	Role           Role  `bson:"role"`
	NumberOfPeople int64 `bson:"number_of_people"` // anti-corruption
}
