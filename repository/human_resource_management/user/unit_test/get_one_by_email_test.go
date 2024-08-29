package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/infrastructor"
	user_repository "shop_erp_mono/repository/human_resource_management/user/repository"
	"testing"
)

func TestGetByEmailUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	mockUser := &userdomain.User{
		Email: "admin@admin.com",
	}
	mockEmptyUser := &userdomain.User{}

	t.Run("success", func(t *testing.T) {
		ur := user_repository.NewUserRepository(database, individual)
		_, err := ur.GetByEmail(context.Background(), mockUser.Email)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := user_repository.NewUserRepository(database, individual)

		// Trying to insert an empty user, expecting an error
		_, err := ur.GetByEmail(context.Background(), mockEmptyUser.Email)
		assert.Error(t, err)
	})
}
