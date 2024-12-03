package candidate_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionCandidate = "candidate"
)

type Candidate struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Name       string             `bson:"name" json:"name"`
	Address    string             `bson:"address" json:"address"`
	Gender     string             `bson:"gender" json:"gender"`
	Email      string             `bson:"email" json:"email"`
	Phone      string             `bson:"phone" json:"phone"`
	Resume     string             `bson:"resume" json:"resume"`
	Skills     []string           `bson:"skills" json:"skills"`
	ImageURL   string             `bson:"image_url" json:"image_url"`
	RoleHire   primitive.ObjectID `bson:"role" json:"role"`
	Experience []Experience       `bson:"experience" json:"experience"`
	Education  []Education        `bson:"education" json:"education"`
	Status     string             `bson:"status" json:"status"`             // Applied Interviewing Hired Rejected, On Hold, Offer Sent, Offer Accepted, Offer Declined
	IsSentMail bool               `bson:"is_sent_mail" json:"is_sent_mail"` // use mail custom or system
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type Experience struct {
	CompanyName string    `bson:"company_name" json:"company_name"`
	Role        string    `bson:"role" json:"role"`
	StartDate   time.Time `bson:"start_date" json:"start_date"`
	EndDate     time.Time `bson:"end_date" json:"end_date"`
}

type Education struct {
	SchoolName     string `bson:"school_name" json:"school_name"`
	Degree         string `bson:"degree" json:"degree"`
	Major          string `bson:"major" json:"major"`
	GraduationYear int    `bson:"graduation_year" json:"graduation_year"`
}
