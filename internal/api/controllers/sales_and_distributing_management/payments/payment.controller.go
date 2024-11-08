package payment_controller

import (
	"shop_erp_mono/internal/config"
	payments_domain "shop_erp_mono/internal/domain/sales_and_distribution_management/payments"
)

type PaymentController struct {
	Database       *config.Database
	PaymentUseCase payments_domain.IPaymentUseCase
}
