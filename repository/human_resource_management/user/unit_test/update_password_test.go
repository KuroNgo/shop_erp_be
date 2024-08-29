package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/infrastructor"
	"shop_erp_mono/pkg/password"
	user_repository "shop_erp_mono/repository/human_resource_management/user/repository"
	"testing"
)

func TestUpdatePasswordUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	user := &userdomain.User{
		Email: "admin@admin.com",
	}
	ur := user_repository.NewUserRepository(database, individual)
	userData, err := ur.GetByEmail(context.Background(), user.Email)
	assert.Nil(t, err)

	mockUser := &userdomain.User{
		ID:           userData.ID,
		PasswordHash: "abc",
	}
	mockEmptyPasswordUser := &userdomain.User{
		ID: userData.ID,
	}
	mockEmptyUser := &userdomain.User{}
	mockUser.PasswordHash, _ = password.HashPassword(mockUser.PasswordHash)

	t.Run("success", func(t *testing.T) {
		err = ur.UpdatePassword(context.Background(), mockUser)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.UpdatePassword(context.Background(), mockEmptyPasswordUser)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.UpdatePassword(context.Background(), mockEmptyUser)
		assert.Error(t, err)
	})
}
