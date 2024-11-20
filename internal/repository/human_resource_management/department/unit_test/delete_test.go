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

func TestDeleteOneDepartment(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	// Thiết lập tên department cần xóa
	nameDepartment := "marketing"
	ur := department_repository.NewDepartmentRepository(database, Department)

	// Drop collection "departments" để xóa toàn bộ dữ liệu cũ
	err := database.Collection("department").Drop(context.Background())
	if err != nil {
		t.Fatalf("Error dropping departments collection: %v", err)
	}

	// Thêm một vài department mẫu vào cơ sở dữ liệu
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

	// Thêm departments vào database
	for _, dept := range departments {
		_, err = database.Collection("department").InsertOne(context.Background(), dept)
		if err != nil {
			t.Fatalf("Error inserting department: %v", err)
		}
	}

	// Lấy thông tin department "marketing" để kiểm tra trước khi xóa
	departmentData, err := ur.GetByName(context.Background(), nameDepartment)
	if err != nil {
		t.Fatalf("Error retrieving department: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		// Thử xóa department đã tồn tại
		err = ur.DeleteOne(context.Background(), departmentData.ID)
		assert.Nil(t, err, "Expected no error when deleting an existing department")
	})

	t.Run("error with invalid ID", func(t *testing.T) {
		// Xử lý trường hợp ID không hợp lệ (ví dụ nil hoặc rỗng)
		invalidID := primitive.NilObjectID
		err := ur.DeleteOne(context.Background(), invalidID)
		assert.Error(t, err, "Expected an error when trying to delete a department with an invalid ID")
	})
}
