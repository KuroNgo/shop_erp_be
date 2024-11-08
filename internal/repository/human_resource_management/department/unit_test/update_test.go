package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	"shop_erp_mono/infrastructor"
	department_repository "shop_erp_mono/repository/human_resource_management/department/repository"
	"testing"
)

func TestUpdateOneDepartment(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	nameDepartment := "marketing"
	ur := department_repository.NewDepartmentRepository(database, Department)
	departmentData, err := ur.GetOneByName(context.Background(), nameDepartment)
	if err != nil {
		assert.Error(t, err)
	}

	mockNilDepartment := departmentsdomain.Department{}
	t.Run("success", func(t *testing.T) {
		err = ur.UpdateOne(context.Background(), departmentData.ID, &departmentData)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.UpdateOne(context.Background(), mockNilDepartment.ID, &mockNilDepartment)
		assert.Error(t, err)
	})
}
