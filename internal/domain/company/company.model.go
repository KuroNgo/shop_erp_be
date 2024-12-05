package company_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionCompany = "company"
)

type Company struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TaxID       string             `bson:"tax_id" json:"tax_id"`
	Represent   string             `bson:"represent" json:"represent"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	LogoURL     string             `bson:"logo_url" json:"logo_url"`
	Address     string             `bson:"address" json:"address"`
	Verify      bool               `bson:"verify" json:"verify"`
	Status      string             `bson:"status" json:"status"`
	//PaymentStatus string             `bson:"payment_status" json:"payment_status"`
	//LevelPayment  string             `bson:"level_payment" json:"level_payment"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	TaxID       string `bson:"tax_id" json:"tax_id"`
	Represent   string `bson:"represent" json:"represent"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	LogoURL     string `bson:"logo_url" json:"logo_url"`
	Address     string `bson:"address" json:"address"`
}

type CompanyResponse struct {
	Company         []Company `bson:"company" json:"company"`
	CountDepartment int       `bson:"count_department" json:"count_department"`
	CountEmployee   int       `bson:"count_employee" json:"count_employee"`
}
