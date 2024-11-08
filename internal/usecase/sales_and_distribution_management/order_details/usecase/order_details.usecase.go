package order_detail_usecase

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	orderdetailsdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/order_details"
	saleordersdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_orders"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
	"shop_erp_mono/internal/usecase/sales_and_distribution_management/order_details/validate"
	"time"
)

type orderDetailUseCase struct {
	contextTimeout        time.Duration
	orderDetailRepository orderdetailsdomain.IOrderDetailRepository
	saleOrderRepository   saleordersdomain.ISalesOrderRepository
	productRepository     productdomain.IProductRepository
	cache                 *bigcache.BigCache
}

func NewOrderDetailUseCase(contextTimeout time.Duration, orderDetailRepository orderdetailsdomain.IOrderDetailRepository,
	saleOrderRepository saleordersdomain.ISalesOrderRepository, productRepository productdomain.IProductRepository,
	cacheTTL time.Duration) orderdetailsdomain.IOrderDetailUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &orderDetailUseCase{contextTimeout: contextTimeout, cache: cache, orderDetailRepository: orderDetailRepository, saleOrderRepository: saleOrderRepository, productRepository: productRepository}
}

func (o *orderDetailUseCase) CreateOne(ctx context.Context, input *orderdetailsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, o.contextTimeout)
	defer cancel()

	if err := validate.OrderDetail(input); err != nil {
		return err
	}

	idOrder, err := primitive.ObjectIDFromHex(input.OrderID)
	if err != nil {
		return err
	}

	saleOrderData, err := o.saleOrderRepository.GetByID(ctx, idOrder)
	if err != nil {
		return err
	}

	order := orderdetailsdomain.OrderDetail{
		ID:      primitive.NewObjectID(),
		OrderID: saleOrderData.ID,
	}

	return o.orderDetailRepository.CreateOne(ctx, order)
}

func (o *orderDetailUseCase) GetByID(ctx context.Context, id string) (*orderdetailsdomain.OrderDetailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, o.contextTimeout)
	defer cancel()

	idOrder, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	orderDetailData, err := o.orderDetailRepository.GetByID(ctx, idOrder)
	if err != nil {
		return nil, err
	}

	orderData, err := o.saleOrderRepository.GetByID(ctx, orderDetailData.OrderID)
	if err != nil {
		return nil, err
	}

	productData, err := o.productRepository.GetByID(ctx, orderDetailData.ProductID)
	if err != nil {
		return nil, err
	}

	response := &orderdetailsdomain.OrderDetailResponse{
		OrderDetail: *orderDetailData,
		Order:       *orderData,
		Product:     *productData,
	}

	return response, nil
}

func (o *orderDetailUseCase) GetByOrderID(ctx context.Context, orderID string) ([]orderdetailsdomain.OrderDetailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, o.contextTimeout)
	defer cancel()

	idOrder, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, err
	}

	orderDetailData, err := o.orderDetailRepository.GetByOrderID(ctx, idOrder)
	if err != nil {
		return nil, err
	}

	var responses []orderdetailsdomain.OrderDetailResponse
	responses = make([]orderdetailsdomain.OrderDetailResponse, 0, len(orderDetailData))
	for _, order := range orderDetailData {
		orderData, err := o.saleOrderRepository.GetByID(ctx, order.OrderID)
		if err != nil {
			return nil, err
		}

		productData, err := o.productRepository.GetByID(ctx, order.ProductID)
		if err != nil {
			return nil, err
		}

		response := orderdetailsdomain.OrderDetailResponse{
			OrderDetail: order,
			Order:       *orderData,
			Product:     *productData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (o *orderDetailUseCase) GetByProductID(ctx context.Context, productID string) ([]orderdetailsdomain.OrderDetailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, o.contextTimeout)
	defer cancel()

	idProduct, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}

	orderDetailData, err := o.orderDetailRepository.GetByProductID(ctx, idProduct)
	if err != nil {
		return nil, err
	}

	var responses []orderdetailsdomain.OrderDetailResponse
	responses = make([]orderdetailsdomain.OrderDetailResponse, 0, len(orderDetailData))
	for _, order := range orderDetailData {
		orderData, err := o.saleOrderRepository.GetByID(ctx, order.OrderID)
		if err != nil {
			return nil, err
		}

		productData, err := o.productRepository.GetByID(ctx, order.ProductID)
		if err != nil {
			return nil, err
		}

		response := orderdetailsdomain.OrderDetailResponse{
			OrderDetail: order,
			Order:       *orderData,
			Product:     *productData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (o *orderDetailUseCase) UpdateOne(ctx context.Context, id string, input *orderdetailsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, o.contextTimeout)
	defer cancel()

	if err := validate.OrderDetail(input); err != nil {
		return err
	}

	idOrderDetail, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	idOrder, err := primitive.ObjectIDFromHex(input.OrderID)
	if err != nil {
		return err
	}

	idProduct, err := primitive.ObjectIDFromHex(input.ProductID)
	if err != nil {
		return err
	}

	totalPrice := float64(input.Quantity) * input.UnitPrice

	order := orderdetailsdomain.OrderDetail{
		ID:         idOrderDetail,
		OrderID:    idOrder,
		ProductID:  idProduct,
		Quantity:   input.Quantity,
		UnitPrice:  input.UnitPrice,
		TotalPrice: totalPrice,
		UpdatedAt:  time.Now(),
	}

	return o.orderDetailRepository.UpdateOne(ctx, order)
}

func (o *orderDetailUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, o.contextTimeout)
	defer cancel()

	idOrderDetail, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return o.orderDetailRepository.DeleteOne(ctx, idOrderDetail)
}

func (o *orderDetailUseCase) GetAll(ctx context.Context) ([]orderdetailsdomain.OrderDetailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, o.contextTimeout)
	defer cancel()

	orderDetailData, err := o.orderDetailRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []orderdetailsdomain.OrderDetailResponse
	responses = make([]orderdetailsdomain.OrderDetailResponse, 0, len(orderDetailData))
	for _, order := range orderDetailData {
		orderData, err := o.saleOrderRepository.GetByID(ctx, order.OrderID)
		if err != nil {
			return nil, err
		}

		productData, err := o.productRepository.GetByID(ctx, order.ProductID)
		if err != nil {
			return nil, err
		}

		response := orderdetailsdomain.OrderDetailResponse{
			OrderDetail: order,
			Order:       *orderData,
			Product:     *productData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
