package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/infrastructor"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	"testing"
	"time"
)

func TestCreateOneEmployee(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

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
	mockEmptyEmployee := &employeesdomain.Employee{}

	t.Run("success", func(t *testing.T) {
		ur := employeerepository.NewEmployeeRepository(database, staff)
		err := ur.CreateOne(context.Background(), mockEmployee)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := employeerepository.NewEmployeeRepository(database, staff)
		// Trying to insert an empty user, expecting an error
		err := ur.CreateOne(context.Background(), mockEmptyEmployee)
		assert.Error(t, err)
	})

}
