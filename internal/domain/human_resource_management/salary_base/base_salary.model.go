package base_salary_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionBaseSalary = "base_salary"
)

type BaseSalary struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RoleID       primitive.ObjectID `bson:"role_id" json:"role_id"`
	UnitCurrency string             `bson:"unit_currency" json:"unit_currency"`
	BaseSalaries float64            `bson:"base_salary" json:"base_salary"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	RoleID       primitive.ObjectID `bson:"role_id" json:"role_id"`
	BaseSalary   float64            `bson:"base_salary" json:"base_salary" exmaple:"1500"`
	UnitCurrency string             `bson:"unit_currency" json:"unit_currency" exmaple:"USD"`
}
