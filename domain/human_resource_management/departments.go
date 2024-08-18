package human_resource_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Department struct represents a department within the company.
type Department struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
