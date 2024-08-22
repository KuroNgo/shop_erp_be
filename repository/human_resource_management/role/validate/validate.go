package validate

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsNilTitle(title string) error {
	if title == "" {
		return errors.New("the role's information do not nil")
	}

	return nil
}

func IsNilDescription(description string) error {
	if description == "" {
		return errors.New("the role's information do not nil")
	}

	return nil
}

func IsNilID(id string) error {
	data, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if data == primitive.NilObjectID {
		return errors.New("the role's information do not nil")
	}

	return nil
}
