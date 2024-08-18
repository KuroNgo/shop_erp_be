package human_resource_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Role struct represents a role or job position.
type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
