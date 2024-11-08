package validate

import (
	"errors"
	sale_reports_domain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_reports"
)

func SaleReport(input *sale_reports_domain.Input) error {
	if input.ReportDate.IsZero() {
		return errors.New("the sale report's information is invalid")
	}

	if input.TotalSales < 0 {
		return errors.New("the payment's information is invalid")
	}

	if input.ProductName == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.Product == "" {
		return errors.New("the payment's information is invalid")
	}

	if input.QuantitySold < 0 {
		return errors.New("the payment's information is invalid")
	}

	return nil
}
