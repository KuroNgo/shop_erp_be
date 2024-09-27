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
	EmployeeID       primitive.ObjectID `bson:"employee_id" json:"employee_id"`
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
	ID               string `bson:"_id" json:"id,omitempty"`
	EmployeeID       string `bson:"employee_id" json:"employee_id"`
	Username         string `bson:"username" json:"username"`
	PasswordHash     string `bson:"password_hash" json:"password_hash"` // Hash of the password
	AvatarURL        string `bson:"avatar_url"  json:"avatar_url"`
	Email            string `bson:"email" json:"email"`
	Phone            string `json:"phone" bson:"phone"`
	Verified         bool   `bson:"verify"   json:"verify"`
	VerificationCode string `bson:"verification_code" json:"verification_code"`
}

type SignIn struct {
	Email    string `bson:"email" json:"email" example:"admin@admin.com" `
	Password string `bson:"password_hash" json:"password_hash" example:"12345"`
}

type VerificationInput struct {
	VerificationCode string `json:"verification_code" binding:"required"`
}

type ChangePasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordCompare string `json:"password_compare" binding:"required"`
}

type ForgetPassword struct {
	Email string `json:"email" bson:"email"`
}

type Output struct {
	User User `bson:"user" json:"user"`
}

type OutputLogin struct {
	RefreshToken string `bson:"refresh_token"`
	AccessToken  string `bson:"access_token"`
	IsLogged     string `bson:"is_logged"`
}

type OutputLoginGoogle struct {
	RefreshToken string `bson:"refresh_token"`
	AccessToken  string `bson:"access_token"`
	IsLogged     string `bson:"is_logged"`
	SignedToken  string `bson:"signed_token"`
}
