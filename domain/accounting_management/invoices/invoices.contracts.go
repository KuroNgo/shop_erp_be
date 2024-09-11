package invoices_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IInvoicesRepository interface {
	CreateInvoice(ctx context.Context, invoice *Invoices) error                                   // Tạo một hóa đơn mới
	GetInvoiceByID(ctx context.Context, id primitive.ObjectID) (*Invoices, error)                 // Lấy thông tin hóa đơn theo ID
	GetInvoiceByName(ctx context.Context, name string) (*Invoices, error)                         // Lấy hóa đơn theo tên
	UpdateInvoice(ctx context.Context, invoice *Invoices) error                                   // Cập nhật thông tin hóa đơn
	DeleteInvoice(ctx context.Context, id primitive.ObjectID) error                               // Xóa hóa đơn theo ID
	ListInvoices(ctx context.Context) ([]Invoices, error)                                         // Lấy danh sách tất cả hóa đơn
	GetInvoicesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Invoices, error) // Lấy hóa đơn theo khoảng thời gian
	GetOverdueInvoices(ctx context.Context) ([]Invoices, error)                                   // Lấy danh sách hóa đơn quá hạn
	MarkInvoiceAsPaid(ctx context.Context, id primitive.ObjectID) error                           // Đánh dấu hóa đơn là đã thanh toán
}

type IInvoicesUseCase interface {
	CreateInvoice(ctx context.Context, input *Input) (InvoicesResponse, error)                         // Tạo hóa đơn từ đầu vào của người dùng
	GetInvoice(ctx context.Context, id string) (InvoicesResponse, error)                               // Lấy thông tin hóa đơn theo ID
	GetInvoiceByName(ctx context.Context, name string) (InvoicesResponse, error)                       // Lấy hóa đơn theo tên
	UpdateInvoice(ctx context.Context, id string, input *Input) (InvoicesResponse, error)              // Cập nhật hóa đơn từ đầu vào
	DeleteInvoice(ctx context.Context, id string) error                                                // Xóa hóa đơn theo ID
	ListInvoices(ctx context.Context) ([]InvoicesResponse, error)                                      // Lấy danh sách tất cả hóa đơn
	GetInvoicesByDateRange(ctx context.Context, startDate, endDate string) ([]InvoicesResponse, error) // Lấy hóa đơn theo khoảng thời gian
	GetOverdueInvoices(ctx context.Context) ([]InvoicesResponse, error)                                // Lấy danh sách hóa đơn quá hạn
	MarkInvoiceAsPaid(ctx context.Context, id string) error                                            // Đánh dấu hóa đơn là đã thanh toán
	SendInvoiceReminder(ctx context.Context, id string) error                                          // Gửi nhắc nhở thanh toán hóa đơn
	GenerateInvoiceReport(ctx context.Context, startDate, endDate string) (InvoiceReport, error)       // Tạo báo cáo hóa đơn theo thời gian
}
