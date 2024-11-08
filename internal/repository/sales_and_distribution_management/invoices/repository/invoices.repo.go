package invoice_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	invoicesdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/invoices"
	"time"
)

type invoicesRepository struct {
	database          *mongo.Database
	invoiceCollection string
}

func NewInvoiceRepository(database *mongo.Database, invoiceCollection string) invoicesdomain.InvoiceRepository {
	return &invoicesRepository{database: database, invoiceCollection: invoiceCollection}
}

func (i *invoicesRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*invoicesdomain.Invoice, error) {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{"_id": id}
	var invoice invoicesdomain.Invoice
	err := invoiceCollection.FindOne(ctx, filter).Decode(&invoice)
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (i *invoicesRepository) CreateOne(ctx context.Context, invoice invoicesdomain.Invoice) error {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	_, err := invoiceCollection.InsertOne(ctx, invoice)
	if err != nil {
		return err
	}

	return nil
}

func (i *invoicesRepository) GetByOrderID(ctx context.Context, orderID primitive.ObjectID) ([]invoicesdomain.Invoice, error) {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{"order_id": orderID}
	cursor, err := invoiceCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []invoicesdomain.Invoice
	orders = make([]invoicesdomain.Invoice, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var order invoicesdomain.Invoice
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (i *invoicesRepository) GetByStatus(ctx context.Context, status string) ([]invoicesdomain.Invoice, error) {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{"status": status}
	cursor, err := invoiceCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []invoicesdomain.Invoice
	orders = make([]invoicesdomain.Invoice, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var order invoicesdomain.Invoice
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (i *invoicesRepository) UpdateOne(ctx context.Context, invoice invoicesdomain.Invoice) error {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{"_id": invoice.ID}
	update := bson.M{"$set": bson.M{
		"order_id":     invoice.OrderID,
		"invoice_date": invoice.InvoiceDate,
		"due_date":     invoice.DueDate,
		"amount_due":   invoice.AmountDue,
		"amount_paid":  invoice.AmountPaid,
		"status":       invoice.Status,
		"updated_at":   time.Now(),
	}}

	_, err := invoiceCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (i *invoicesRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{"_id": id}
	_, err := invoiceCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (i *invoicesRepository) GetAll(ctx context.Context) ([]invoicesdomain.Invoice, error) {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{}
	cursor, err := invoiceCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []invoicesdomain.Invoice
	orders = make([]invoicesdomain.Invoice, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var order invoicesdomain.Invoice
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
