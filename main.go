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

// @tag.name Account
// @tag.description Stores information about financial accounts like bank accounts or cash wallets, tracking balance and account type.
// @tag.name Budget
// @tag.description Manages financial transactions for accounts, tracking income and expenses, amounts, and transaction dates.
// @tag.name Budget
// @tag.description Manages financial transactions for accounts, tracking income and expenses, amounts, and transaction dates.
// @tag.name FinancialReport
// @tag.description Generates financial reports over a specified date range for analysis and tracking.
// @tag.name Invoices
// @tag.description Tracks issued invoices, including amounts, due dates, and payment status.
// @tag.name Payments
// @tag.description Records payments made towards invoices, linking them to accounts and payment methods.
// @tag.name Taxes
// @tag.description Defines tax types and rates applicable to invoices, transactions, or other financial elements.
// @tag.name TransactionCategory
// @tag.description Categorizes transactions into types like income or expenses for better organization.
// @tag.name Transaction
// @tag.description Manages financial transactions for accounts, tracking income and expenses, amounts, and transaction dates.

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

// @tag.name Customer
// @tag.description Stores information about customers, including their contact details (name, email, phone, address) and basic information for future reference and order management.
// @tag.name Invoices
// @tag.description Contains details of the products available for sale, such as product name, description, price, and stock quantity. This helps track inventory and product offerings.
// @tag.name OrderDetails
// @tag.description Contains information about individual items in each order. It tracks the products, quantities, and unit prices associated with a specific order to provide more detailed order history.
// @tag.name Payments
// @tag.descriptionManages payment details for orders, including the payment method, status (paid or unpaid), and amount paid, ensuring accurate financial records.
// @tag.name SaleOrders
// @tag.description Tracks customer orders, including customer details, shipping information, total amount, and the status of the order (processing, shipped, or canceled).
// @tag.name SaleReport
// @tag.description Manages financial transactions for accounts, tracking income and expenses, amounts, and transaction dates.
// @tag.name Shipping
// @tag.description Records shipment details, including the shipping method, dates, tracking information, and the status of the delivery to monitor the progress of the shipping process.

// @tag.name Inventory
// @tag.description Tracks the stock levels of products in different warehouses. It helps monitor how much stock is available and where it is located.
// @tag.name Product
// @tag.description Contains details of products such as name, description, price, and associated categories. It is central to managing what items are available for sale and for inventory control.
// @tag.name ProductCategory
// @tag.description Organizes products into different categories for better product classification and easier management. It allows for filtering and categorizing similar products.
// @tag.name PurchaseOrder
// @tag.description Manages purchase orders sent to suppliers. It includes order numbers, supplier references, order dates, and status, ensuring that product orders are properly tracked.
// @tag.name PurchaseOrderDetail
// @tag.description Provides detailed information about the individual products in each purchase order, including quantities and prices. It links the products ordered from suppliers to their respective purchase orders.
// @tag.name StockAdjustment
// @tag.description Keeps records of stock adjustments, such as increases or decreases in inventory due to reasons like damage, returns, or reallocation. This ensures accurate stock levels are maintained.
// @tag.name StockMovement
// @tag.description Records the movements of products in and out of warehouses, such as when items are received (Nhập kho) or dispatched (Xuất kho). This helps track changes in inventory levels.
// @tag.name Supplier
// @tag.description Stores supplier information, including contact details. This table is crucial for managing relationships with vendors and tracking where products are sourced.
// @tag.name Warehouse
// @tag.description This table stores information about different warehouses, including their location, capacity, and basic details. It helps manage the distribution and storage of products across different facilities.

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
