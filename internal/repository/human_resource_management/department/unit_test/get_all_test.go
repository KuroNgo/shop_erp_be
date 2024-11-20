package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	departmentsdomain "shop_erp_mono/internal/domain/human_resource_management/departments"
	infrastructor "shop_erp_mono/internal/infrastructor/mongo"
	department_repository "shop_erp_mono/internal/repository/human_resource_management/department/repository"
	"testing"
	"time"
)

func TestGetAll(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	// Drop collection "departments" để xóa toàn bộ dữ liệu cũ
	err := database.Collection("department").Drop(context.Background())
	if err != nil {
		t.Fatalf("Error dropping departments collection: %v", err)
	}

	// Thêm một vài department mẫu vào cơ sở dữ liệu trước khi kiểm tra
	departments := []departmentsdomain.Department{
		{
			ID:          primitive.NewObjectID(),
			Name:        "marketing",
			Description: "Marketing Department",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "sales",
			Description: "Sales Department",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// Insert departments vào database
	for _, dept := range departments {
		_, err := database.Collection("department").InsertOne(context.Background(), dept)
		if err != nil {
			t.Fatalf("Error inserting department: %v", err)
		}
	}

	// Kiểm tra GetAll
	t.Run("success", func(t *testing.T) {
		ur := department_repository.NewDepartmentRepository(database, Department)
		deptList, err := ur.GetAll(context.Background())
		assert.Nil(t, err)
		assert.NotNil(t, deptList) // Kiểm tra kết quả trả về không phải nil
		assert.Len(t, deptList, len(departments), "Expected the number of departments to match the inserted departments")
	})
}
