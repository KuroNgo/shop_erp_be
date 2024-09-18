package payment_controller

import (
	"shop_erp_mono/bootstrap"
	payments_domain "shop_erp_mono/domain/sales_and_distribution_management/payments"
)

type PaymentController struct {
	Database       *bootstrap.Database
	PaymentUseCase payments_domain.IPaymentUseCase
}
