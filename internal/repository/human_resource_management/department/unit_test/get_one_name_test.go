package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"shop_erp_mono/infrastructor"
	department_repository "shop_erp_mono/repository/human_resource_management/department/repository"
	"testing"
)

func TestGetOneName(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	departmentName := "marketing"
	departmentName2 := ""

	t.Run("success", func(t *testing.T) {
		ur := department_repository.NewDepartmentRepository(database, Department)
		_, err := ur.GetOneByName(context.Background(), departmentName)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := department_repository.NewDepartmentRepository(database, Department)
		_, err := ur.GetOneByName(context.Background(), departmentName2)
		assert.Error(t, err)
	})
}
