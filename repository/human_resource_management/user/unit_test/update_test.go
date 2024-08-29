package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/infrastructor"
	user_repository "shop_erp_mono/repository/human_resource_management/user/repository"
	"testing"
	"time"
)

func TestUpdateUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	user := &userdomain.User{
		Email: "admin@admin.com",
	}
	ur := user_repository.NewUserRepository(database, individual)
	userData, err := ur.GetByEmail(context.Background(), user.Email)
	assert.Nil(t, err)

	mockUser := &userdomain.UpdateUser{
		ID:        userData.ID,
		Username:  "unit_test",
		AvatarURL: "",
		UpdatedAt: time.Now(),
	}
	mockEmptyUser := &userdomain.UpdateUser{}
	mockEmptyDataUser := &userdomain.UpdateUser{
		ID: userData.ID,
	}

	t.Run("success", func(t *testing.T) {
		err = ur.Update(context.Background(), mockUser)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.Update(context.Background(), mockEmptyDataUser)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.Update(context.Background(), mockEmptyUser)
		assert.Error(t, err)
	})
}
