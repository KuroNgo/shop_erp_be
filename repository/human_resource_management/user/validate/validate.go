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
