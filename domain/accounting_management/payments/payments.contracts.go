package payments_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IPaymentsRepository interface {
	CreatePayment(ctx context.Context, payment *Payments) error
	GetPaymentByID(ctx context.Context, id primitive.ObjectID) (Payments, error)
	GetPaymentsByInvoiceID(ctx context.Context, invoiceID primitive.ObjectID) ([]Payments, error)
	GetPaymentsByAccountID(ctx context.Context, accountID primitive.ObjectID) ([]Payments, error)
	UpdatePayment(ctx context.Context, payment *Payments) error
	DeletePayment(ctx context.Context, id primitive.ObjectID) error
	ListPayments(ctx context.Context) ([]Payments, error)
	GetPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Payments, error) // Lấy thanh toán theo khoảng thời gian
	GetTotalPaymentsAmount(ctx context.Context) (int32, error)                                    // Lấy tổng số tiền của tất cả thanh toán
}

type IPaymentsUseCase interface {
	CreatePayment(ctx context.Context, input *Input) error
	GetPaymentByID(ctx context.Context, id string) (PaymentsResponse, error)
	GetPaymentsByInvoiceID(ctx context.Context, invoiceID string) ([]PaymentsResponse, error)
	GetPaymentsByAccountID(ctx context.Context, accountID string) ([]PaymentsResponse, error)
	UpdatePayment(ctx context.Context, id string, input *Input) error
	DeletePayment(ctx context.Context, id string) error
	ListPayments(ctx context.Context) ([]PaymentsResponse, error)
	GetPaymentsByDateRange(ctx context.Context, startDate, endDate string) ([]PaymentsResponse, error) // Lấy thanh toán theo khoảng thời gian
	GetTotalPaymentsAmount(ctx context.Context) (int32, error)                                         // Lấy tổng số tiền của tất cả thanh toán
}
