package user_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionUser = "user"
)

// User represents a user in the system.
type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username         string             `bson:"username" json:"username"`
	PasswordHash     string             `bson:"password_hash" json:"password_hash"` // Hash of the password
	Email            string             `bson:"email" json:"email"`
	Phone            string             `json:"phone" bson:"phone"`
	AssetURL         string             `bson:"asset_url"  json:"asset_url"`
	AvatarURL        string             `bson:"avatar_url"  json:"avatar_url"`
	Verified         bool               `bson:"verify"   json:"verify"`
	VerificationCode string             `bson:"verification_code" json:"verification_code"`
	Provider         string             `bson:"provider" json:"provider"`
	Role             string             `bson:"role" json:"role"` // Example: "Admin", "Manager", "Employee"
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	ID           string `bson:"_id" json:"id,omitempty"`
	Username     string `bson:"username" json:"username"`
	PasswordHash string `bson:"password_hash" json:"password_hash"` // Hash of the password
	Email        string `bson:"email" json:"email"`
	Phone        string `json:"phone" bson:"phone"`
	AssetURL     string `bson:"asset_url"  json:"asset_url"`
	AvatarURL    string `bson:"avatar_url"  json:"avatar_url"`
}

type Output struct {
	User User `bson:"user" json:"user"`
}

type OutputLogin struct {
	RefreshToken string `bson:"refresh_token"`
	AccessToken  string `bson:"access_token"`
	IsLogged     bool   `bson:"is_logged"`
}
