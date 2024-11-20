package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	infrastructor "shop_erp_mono/internal/infrastructor/mongo"
	employeerepository "shop_erp_mono/internal/repository/human_resource_management/employee/repository"
	"testing"
	"time"
)

func TestUpdateOneEmployee(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	ur := employeerepository.NewEmployeeRepository(database, staff)

	// Tạo dữ liệu nhân viên mock để test
	mockEmployee := &employeesdomain.Employee{
		FirstName:   "Ngô",
		LastName:    "Hoài Phong",
		Email:       "admin@admin.com",
		Phone:       "0329245971",
		Address:     "Bình Thuận",
		AvatarURL:   "https://example.com/avatar.jpg",
		DateOfBirth: time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
		DayOfWork:   time.Date(2024, 8, 30, 0, 0, 0, 0, time.UTC),
	}

	// Thêm mockEmployee vào database
	err := ur.CreateOne(context.Background(), mockEmployee)
	assert.Nil(t, err)

	// Lấy ID của nhân viên vừa được thêm
	employeeData, err := ur.GetByEmail(context.Background(), "admin@admin.com")
	if err != nil {
		t.Fatalf("error setting up test data: %v", err)
	}

	// Cập nhật thông tin nhân viên
	updatedEmployee := &employeesdomain.Employee{
		FirstName:   "Ngô",
		LastName:    "Hoài Phong",
		Email:       "admin@admin.com",
		Phone:       "0329245971",
		Address:     "TP.HCM",                             // Cập nhật địa chỉ
		AvatarURL:   "https://example.com/new-avatar.jpg", // Cập nhật avatar
		DateOfBirth: time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
		DayOfWork:   time.Date(2024, 8, 30, 0, 0, 0, 0, time.UTC),
	}

	t.Run("success", func(t *testing.T) {
		err = ur.UpdateOne(context.Background(), employeeData.ID, updatedEmployee)
		assert.Nil(t, err)

		updatedEmployeeData, err := ur.GetByID(context.Background(), employeeData.ID)
		assert.Nil(t, err)
		assert.Equal(t, updatedEmployee.Address, updatedEmployeeData.Address)     // Kiểm tra cập nhật địa chỉ
		assert.Equal(t, updatedEmployee.AvatarURL, updatedEmployeeData.AvatarURL) // Kiểm tra cập nhật avatar
	})
}
