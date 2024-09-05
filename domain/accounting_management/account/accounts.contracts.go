package account_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IAccountRepository interface {
	CreateAccount(ctx context.Context, account *Accounts) error                                   // Tạo tài khoản mới
	GetAccountByID(ctx context.Context, id primitive.ObjectID) (Accounts, error)                  // Lấy tài khoản theo ID
	GetAccountByName(ctx context.Context, name string) (Accounts, error)                          // Lấy tài khoản theo tên
	UpdateAccount(ctx context.Context, account *Accounts) error                                   // Cập nhật tài khoản
	DeleteAccount(ctx context.Context, id primitive.ObjectID) error                               // Xóa tài khoản theo ID
	ListAccounts(ctx context.Context) ([]Accounts, error)                                         // Lấy danh sách tất cả tài khoản
	GetAccountsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Accounts, error) // Lấy tài khoản theo khoảng thời gian
	GetTotalAccountBalance(ctx context.Context) (float64, error)                                  // Lấy tổng số dư của tất cả tài khoản
	DeactivateAccount(ctx context.Context, id primitive.ObjectID) error                           // Đánh dấu tài khoản là không hoạt động
	ReactivateAccount(ctx context.Context, id primitive.ObjectID) error                           // Kích hoạt lại tài khoản
}

type IAccountUseCase interface {
	CreateAccount(ctx context.Context, input *Input) error                                            // Tạo tài khoản từ đầu vào của người dùng
	GetAccountByID(ctx context.Context, id string) (AccountResponse, error)                           // Lấy tài khoản theo ID
	GetAccountByName(ctx context.Context, name string) (AccountResponse, error)                       // Lấy tài khoản theo tên
	UpdateAccount(ctx context.Context, id string, input *Input) error                                 // Cập nhật tài khoản từ đầu vào
	DeleteAccount(ctx context.Context, id string) error                                               // Xóa tài khoản theo ID
	ListAccounts(ctx context.Context) ([]AccountResponse, error)                                      // Lấy danh sách tất cả tài khoản
	GetAccountsByDateRange(ctx context.Context, startDate, endDate string) ([]AccountResponse, error) // Lấy tài khoản theo khoảng thời gian
	GetTotalAccountBalance(ctx context.Context) (float64, error)                                      // Lấy tổng số dư của tất cả tài khoản
	DeactivateAccount(ctx context.Context, id string) error                                           // Đánh dấu tài khoản là không hoạt động
	ReactivateAccount(ctx context.Context, id string) error                                           // Kích hoạt lại tài khoản
}
