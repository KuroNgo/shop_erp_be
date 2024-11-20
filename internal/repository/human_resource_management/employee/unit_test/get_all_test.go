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

func TestGetAll(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	ur := employeerepository.NewEmployeeRepository(database, staff)
	mockEmployees := []employeesdomain.Employee{
		{
			FirstName:   "Ngô",
			LastName:    "Hoài Phong",
			Email:       "hoaiphong01012002@gmail.com",
			Phone:       "0329245971",
			Address:     "Bình Thuận",
			AvatarURL:   "https://example.com/avatar.jpg",
			DateOfBirth: time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
			DayOfWork:   time.Date(2024, 8, 30, 0, 0, 0, 0, time.UTC),
		},
		{
			FirstName:   "John",
			LastName:    "Doe",
			Email:       "john.doe@example.com",
			Phone:       "0123456789",
			Address:     "Hà Nội",
			AvatarURL:   "https://example.com/john_avatar.jpg",
			DateOfBirth: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
			DayOfWork:   time.Date(2024, 5, 20, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, employee := range mockEmployees {
		err := ur.CreateOne(context.Background(), &employee)
		assert.Nil(t, err)
	}

	t.Run("success", func(t *testing.T) {
		employees, err := ur.GetAll(context.Background())
		assert.Nil(t, err)

		assert.Greater(t, len(employees), 0, "Expected at least one employee")
		assert.Equal(t, "Ngô", employees[0].FirstName)
		assert.Equal(t, "hoaiphong01012002@gmail.com", employees[0].Email)
	})
}
