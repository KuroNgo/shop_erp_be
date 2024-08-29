package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/infrastructor"
	user_repository "shop_erp_mono/repository/human_resource_management/user/repository"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	user := &userdomain.User{
		Email: "admin@admin.com",
	}
	ur := user_repository.NewUserRepository(database, individual)
	userData, err := ur.GetByEmail(context.Background(), user.Email)
	assert.Nil(t, err)

	mockUser := &userdomain.User{
		ID: userData.ID,
	}
	mockEmptyUser := &userdomain.User{}

	t.Run("success", func(t *testing.T) {
		err = ur.DeleteOne(context.Background(), mockUser.ID.Hex())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.DeleteOne(context.Background(), mockEmptyUser.ID.Hex())
		assert.Nil(t, err)
	})
}
