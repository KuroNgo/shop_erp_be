package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	infrastructor "shop_erp_mono/internal/infrastructor/mongo"
	department_repository "shop_erp_mono/internal/repository/human_resource_management/department/repository"
	"testing"
)

func TestGetOneName(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	departmentName := "marketing"

	t.Run("success", func(t *testing.T) {
		// Khởi tạo repository
		ur := department_repository.NewDepartmentRepository(database, Department)

		// Thử lấy department với tên hợp lệ
		department, err := ur.GetByName(context.Background(), departmentName)
		assert.Nil(t, err, "Expected no error when retrieving department with valid name")
		assert.NotNil(t, department, "Expected department to be found")
		assert.Equal(t, departmentName, department.Name, "Department name should match the input name")
	})
}
