package user_validate

import (
	"errors"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
)

func IsInvalidUser(user *userdomain.User) error {
	if user.Username == "" || user.Email == "" || user.Phone == "" || user.PasswordHash == "" || user.Role == "" {
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

func IsNilPhone(user *userdomain.User) error {
	if user.Phone == "" {
		return errors.New("the user's information cannot be empty")
	}
	return nil
}

func IsNilRole(user *userdomain.User) error {
	if user.Role == "" {
		return errors.New("the user's information cannot be empty")
	}
	return nil
}
