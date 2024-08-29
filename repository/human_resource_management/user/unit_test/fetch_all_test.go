package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"shop_erp_mono/infrastructor"
	user_repository "shop_erp_mono/repository/human_resource_management/user/repository"
	"testing"
)

func TestGetAllUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	t.Run("success", func(t *testing.T) {
		ur := user_repository.NewUserRepository(database, individual)
		_, err := ur.FetchMany(context.Background())
		assert.Nil(t, err)
	})
}
