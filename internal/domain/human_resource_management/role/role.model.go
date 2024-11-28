package role_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionRole = "role"
)

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Level       int                `bson:"level" json:"level"`
	Status      string             `bson:"status" json:"status"`
	Enable      int                `bson:"enable" json:"enable"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
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
