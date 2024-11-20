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

func TestGetOneByID(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	ur := employeerepository.NewEmployeeRepository(database, staff)

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

	err := ur.CreateOne(context.Background(), mockEmployee)
	assert.Nil(t, err)

	employeeData, err := ur.GetByEmail(context.Background(), "admin@admin.com")
	if err != nil {
		t.Fatalf("error setting up test data: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		employee, err := ur.GetByID(context.Background(), employeeData.ID)
		assert.Nil(t, err)
		assert.NotNil(t, employee)
		assert.Equal(t, employeeData.ID, employee.ID)
	})
}
