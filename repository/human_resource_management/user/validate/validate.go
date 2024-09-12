package validate

import (
	"errors"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
)

func IsInvalidUser(user *userdomain.User) error {
	if user.Username == "" || user.Email == "" {
		return errors.New("the user's information cannot be empty")
	}
	return nil
}

func IsNilUsername(user *userdomain.User) error {
	if user.Username == "" {
		return errors.New("the user's information cannot be empty")
	}
	return nil
}

func IsNilEmail(email string) error {
	if email == "" {
		return errors.New("the user's information cannot be empty")
	}
	return nil
}

func IsNilPasswordHash(user *userdomain.User) error {
	if user.PasswordHash == "" {
		return errors.New("the user's information cannot be empty")
	}
	return nil
}

func IsNilImage(avatarUrl string) error {
	if avatarUrl == "" {
		return errors.New("the user's information cannot be empty")
	}
	return nil
}

func IsNilUser(input *userdomain.Input) error {
	if input.PasswordHash == "" {
		return errors.New("the user's information cannot be empty")
	}
	if input.Email == "" {
		return errors.New("the user's information cannot be empty")
	}
	if input.Username == "" {
		return errors.New("the user's information cannot be empty")
	}

	return nil
}
