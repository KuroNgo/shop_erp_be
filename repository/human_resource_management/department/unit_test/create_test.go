package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	"shop_erp_mono/infrastructor"
	departmentrepository "shop_erp_mono/repository/human_resource_management/department/repository"
	"testing"
)

func TestCreateOneDepartment(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	mockDepartment := &departmentsdomain.Input{
		Name:        "marketing",
		Description: "abc",
	}
	mockEmptyDepartment := &departmentsdomain.Input{}

	t.Run("success", func(t *testing.T) {
		ur := departmentrepository.NewDepartmentRepository(database, Department)
		err := ur.CreateOne(context.Background(), mockDepartment)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := departmentrepository.NewDepartmentRepository(database, Department)
		// Trying to insert an empty user, expecting an error
		err := ur.CreateOne(context.Background(), mockEmptyDepartment)
		assert.Error(t, err)
	})

}
