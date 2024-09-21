package validate

import (
	"errors"
	invoice_domain "shop_erp_mono/domain/sales_and_distribution_management/invoices"
)

func Invoices(input *invoice_domain.Input) error {
	if input.OrderID == "" {
		return errors.New("the invoice's information is invalid")
	}

	if input.AmountPaid == 0 {
		return errors.New("the invoice's information is invalid")
	}

	if input.AmountDue == 0 {
		return errors.New("the invoice's information is invalid")
	}

	if input.Status == "" {
		return errors.New("the invoice's information is invalid")
	}

	if input.DueDate.IsZero() {
		return errors.New("the invoice's information is invalid")
	}

	if input.InvoiceDate.IsZero() {
		return errors.New("the invoice's information is invalid")
	}

	return nil
}
