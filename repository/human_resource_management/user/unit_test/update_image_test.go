package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/infrastructor"
	user_repository "shop_erp_mono/repository/human_resource_management/user/repository"
	"testing"
)

func TestUpdateImageURLUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	user := &userdomain.User{
		Email: "admin@admin.com",
	}
	ur := user_repository.NewUserRepository(database, individual)
	userData, _ := ur.GetByEmail(context.Background(), user.Email)

	mockUser := &userdomain.User{
		ID:        userData.ID,
		AvatarURL: "abc",
	}
	mockEmptyUser := &userdomain.User{}

	t.Run("success", func(t *testing.T) {
		err := ur.UpdateImage(context.Background(), mockUser.ID.Hex(), mockUser.AvatarURL)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err := ur.UpdateImage(context.Background(), mockEmptyUser.ID.Hex(), mockEmptyUser.AvatarURL)
		assert.Error(t, err)
	})
}
