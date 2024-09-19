package invoices_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	invoicesdomain "shop_erp_mono/domain/accounting_management/invoices"
	"time"
)

type invoiceRepository struct {
	database          *mongo.Database
	invoiceCollection string
}

func NewInvoiceRepository(database *mongo.Database, invoiceCollection string) invoicesdomain.InvoicesRepository {
	return &invoiceRepository{database: database, invoiceCollection: invoiceCollection}
}

func (i *invoiceRepository) CreateInvoice(ctx context.Context, invoice *invoicesdomain.Invoices) error {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	_, err := invoiceCollection.InsertOne(ctx, invoice)
	if err != nil {
		return err
	}
	return nil
}

func (i *invoiceRepository) GetInvoiceByID(ctx context.Context, id primitive.ObjectID) (*invoicesdomain.Invoices, error) {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{"_id": id}
	var invoice *invoicesdomain.Invoices
	if err := invoiceCollection.FindOne(ctx, filter).Decode(&invoice); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return invoice, nil
}

func (i *invoiceRepository) GetInvoiceByName(ctx context.Context, name string) (*invoicesdomain.Invoices, error) {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{"name": name}
	var invoice *invoicesdomain.Invoices
	if err := invoiceCollection.FindOne(ctx, filter).Decode(&invoice); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return invoice, nil
}

func (i *invoiceRepository) UpdateInvoice(ctx context.Context, invoice *invoicesdomain.Invoices) error {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{"_id": invoice.ID}
	update := bson.M{"$set": bson.M{
		"invoice_number": invoice.InvoiceNumber,
		"invoice_date":   invoice.InvoiceDate,
		"customer_id":    invoice.CustomerID,
		"total_amount":   invoice.TotalAmount,
		"status":         invoice.Status,
		"due_date":       invoice.DueDate,
		"updated_at":     time.Now(),
	}}

	_, err := invoiceCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (i *invoiceRepository) DeleteInvoice(ctx context.Context, id primitive.ObjectID) error {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{"_id": id}

	_, err := invoiceCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (i *invoiceRepository) ListInvoices(ctx context.Context) ([]invoicesdomain.Invoices, error) {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{}
	cursor, err := invoiceCollection.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	// Lấy dữ liệu từ cursor và chuyển thành slice of Invoice
	var invoices []invoicesdomain.Invoices
	invoices = make([]invoicesdomain.Invoices, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var invoice invoicesdomain.Invoices
		if err = cursor.Decode(&invoice); err != nil {
			return nil, err
		}

		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

func (i *invoiceRepository) GetInvoicesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]invoicesdomain.Invoices, error) {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{
		"invoice_date": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"invoice_date", 1}})

	cursor, err := invoiceCollection.Find(ctx, filter, findOptions)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	// Lấy dữ liệu từ cursor và chuyển thành slice of Invoice
	var invoices []invoicesdomain.Invoices
	invoices = make([]invoicesdomain.Invoices, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var invoice invoicesdomain.Invoices
		if err = cursor.Decode(&invoice); err != nil {
			return nil, err
		}

		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

func (i *invoiceRepository) GetOverdueInvoices(ctx context.Context) ([]invoicesdomain.Invoices, error) {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{
		"due_date": bson.M{"$lt": time.Now()},
		//"status":   "overdue",
	}

	cursor, err := invoiceCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var invoices []invoicesdomain.Invoices
	invoices = make([]invoicesdomain.Invoices, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var invoice invoicesdomain.Invoices
		if err = cursor.Decode(&invoice); err != nil {
			return nil, err
		}

		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

func (i *invoiceRepository) MarkInvoiceAsPaid(ctx context.Context, id primitive.ObjectID) error {
	invoiceCollection := i.database.Collection(i.invoiceCollection)

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"status": "paid",
		},
	}
	_, err := invoiceCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
