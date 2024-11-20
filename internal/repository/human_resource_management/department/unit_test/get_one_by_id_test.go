package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	infrastructor "shop_erp_mono/internal/infrastructor/mongo"
	department_repository "shop_erp_mono/internal/repository/human_resource_management/department/repository"
	"testing"
)

func TestGetOneByID(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	ur := department_repository.NewDepartmentRepository(database, Department)
	departmentName := "marketing"
	departmentData, err := ur.GetByName(context.Background(), departmentName)

	// Kiểm tra lỗi khi lấy department
	if err != nil {
		t.Fatalf("Error retrieving department: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		// Kiểm tra nếu ID hợp lệ
		department, err := ur.GetByID(context.Background(), departmentData.ID)
		assert.Nil(t, err)
		assert.NotNil(t, department)                      // Đảm bảo kết quả không nil
		assert.Equal(t, departmentData.ID, department.ID) // Kiểm tra ID trùng khớp
	})
}
