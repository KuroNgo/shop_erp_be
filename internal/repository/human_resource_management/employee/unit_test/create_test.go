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

func TestCreateOneEmployee(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	// Function to clear the employee collection before each test case
	clearEmployeeCollection := func() {
		err := database.Collection(staff).Drop(context.Background())
		if err != nil {
			t.Fatalf("Failed to clear employee collection: %v", err)
		}
	}

	mockEmployee := &employeesdomain.Employee{
		FirstName:   "Ngô",
		LastName:    "Hoài Phong",
		Gender:      "Nam",
		Email:       "hoaiphong01012002@gmail.com",
		Phone:       "0329245971",
		Address:     "Bình Thuận",
		AvatarURL:   "https://example.com/avatar.jpg",
		DateOfBirth: time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
		DayOfWork:   time.Date(2024, 8, 30, 0, 0, 0, 0, time.UTC),
	}

	t.Run("success", func(t *testing.T) {
		clearEmployeeCollection() // Clear the collection
		ur := employeerepository.NewEmployeeRepository(database, staff)
		err := ur.CreateOne(context.Background(), mockEmployee)
		assert.Nil(t, err)
	})
}
