package budgets_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IBudgetRepository interface {
	CreateBudget(ctx context.Context, budget *Budget) error                                    // Tạo một ngân sách mới
	GetBudgetByID(ctx context.Context, id primitive.ObjectID) (Budget, error)                  // Lấy thông tin ngân sách theo ID
	GetBudgetByName(ctx context.Context, name string) (Budget, error)                          // Lấy ngân sách theo tên
	UpdateBudget(ctx context.Context, budget *Budget) error                                    // Cập nhật thông tin ngân sách
	DeleteBudget(ctx context.Context, id primitive.ObjectID) error                             // Xóa ngân sách theo ID
	ListBudgets(ctx context.Context) ([]Budget, error)                                         // Lấy danh sách tất cả ngân sách
	GetBudgetsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Budget, error) // Lấy ngân sách theo khoảng thời gian
	GetTotalBudgetAmount(ctx context.Context) (float64, error)                                 // Lấy tổng số tiền của tất cả ngân sách
}

type IBudgetUseCase interface {
	CreateBudget(ctx context.Context, input *Input) error                                           // Tạo ngân sách từ đầu vào của người dùng
	GetBudget(ctx context.Context, id string) (BudgetResponse, error)                               // Lấy thông tin ngân sách theo ID
	GetBudgetByName(ctx context.Context, name string) (BudgetResponse, error)                       // Lấy ngân sách theo tên
	UpdateBudget(ctx context.Context, id string, input *Input) error                                // Cập nhật ngân sách từ đầu vào
	DeleteBudget(ctx context.Context, id string) error                                              // Xóa ngân sách theo ID
	ListBudgets(ctx context.Context) ([]BudgetResponse, error)                                      // Lấy danh sách tất cả ngân sách
	GetBudgetsByDateRange(ctx context.Context, startDate, endDate string) ([]BudgetResponse, error) // Lấy ngân sách theo khoảng thời gian
	GetTotalBudgetAmount(ctx context.Context) (float64, error)                                      // Lấy tổng số tiền của tất cả ngân sách
}
