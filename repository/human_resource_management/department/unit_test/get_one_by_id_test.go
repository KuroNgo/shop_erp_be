package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	"shop_erp_mono/infrastructor"
	department_repository "shop_erp_mono/repository/human_resource_management/department/repository"
	"testing"
)

func TestGetOneByID(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	ur := department_repository.NewDepartmentRepository(database, Department)
	departmentName := "marketing"
	departmentData, err := ur.GetOneByName(context.Background(), departmentName)
	if err != nil {
		assert.Error(t, err)
	}

	mockDepartmentNil := departmentsdomain.Department{}

	t.Run("success", func(t *testing.T) {
		_, err = ur.GetOneByID(context.Background(), departmentData.ID.Hex())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		_, err = ur.GetOneByID(context.Background(), mockDepartmentNil.ID.Hex())
		assert.Error(t, err)
	})
}
