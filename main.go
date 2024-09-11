package main

import (
	"github.com/gin-gonic/gin"
	"shop_erp_mono/api/routers"
	"shop_erp_mono/infrastructor"
	"time"
)

// @title Shop E-commerce ERP
// @version 1.0
// @description This is a server for Kuro API

// @contact.name API Support
// @contact.url
// @contact.email hoaiphong01012002@gmail.com

// @tag.name Attendance
// @tag.description represents the attendance information of an employee.
// @tag.name Benefit
// @tag.description represents the benefits an employee receives.
// @tag.name Contract
// @tag.description represents an employment contract of an employee.
// @tag.name Department
// @tag.description struct represents a department within the company.
// @tag.name Employee
// @tag.description struct represents an employee in the HR system.
// @tag.name Leave Request
// @tag.description represents a leave request by an employee.
// @tag.name Performance Review
// @tag.description represents performance evaluations of an employee.
// @tag.name Role
// @tag.description struct represents a role or job role.
// @tag.name Salary
// @tag.description represents the salary information of an employee.
// @tag.name User
// @tag.description represents a user in the system.
// @tag.name Account
// @tag.description Stores information about financial accounts like bank accounts or cash wallets, tracking balance and account type.
// @tag.name Budget
// @tag.description Manages financial transactions for accounts, tracking income and expenses, amounts, and transaction dates.
// @tag.name Product
// @tag.description

// @host localhost:8080
// @BasePath /api/v1
func main() {

	app := infrastructor.App()

	env := app.Env

	db := app.MongoDB.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	_gin := gin.Default()

	routers.SetUp(env, timeout, db, _gin)
	err := _gin.Run(env.ServerAddress)
	if err != nil {
		return
	}

}
