package purchase_order_detail_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	purchaseorderdomain "shop_erp_mono/domain/warehouse_management/purchase_order"
	purchaseorderdetaildomain "shop_erp_mono/domain/warehouse_management/purchase_order_detail"
	"shop_erp_mono/repository"
	"shop_erp_mono/usecase/warehouse_management/purchase_order_detail/validate"
	"time"
)

type purchaseOrderDetailUseCase struct {
	contextTimeout                time.Duration
	purchaseOrderDetailRepository purchaseorderdetaildomain.IPurchaseOrderDetailRepository
	purchaseOrderRepository       purchaseorderdomain.IPurchaseOrderRepository
	productRepository             productdomain.IProductRepository
}

func NewProductOrderDetailRepository(contextTimeout time.Duration, purchaseOrderDetailRepository purchaseorderdetaildomain.IPurchaseOrderDetailRepository,
	purchaseOrderRepository purchaseorderdomain.IPurchaseOrderRepository, productRepository productdomain.IProductRepository) purchaseorderdetaildomain.IPurchaseOrderDetailUseCase {
	return &purchaseOrderDetailUseCase{contextTimeout: contextTimeout, purchaseOrderDetailRepository: purchaseOrderDetailRepository,
		purchaseOrderRepository: purchaseOrderRepository, productRepository: productRepository}
}

func (p *purchaseOrderDetailUseCase) GetByID(ctx context.Context, id string) (*purchaseorderdetaildomain.PurchaseOrderDetailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	purchaseOrderDetailID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	purchaseOrderDetailData, err := p.purchaseOrderDetailRepository.GetByID(ctx, purchaseOrderDetailID)
	if err != nil {
		return nil, err
	}

	purchaseOrderData, err := p.purchaseOrderRepository.GetByID(ctx, purchaseOrderDetailData.PurchaseOrderID)
	if err != nil {
		return nil, err
	}

	productData, err := p.productRepository.GetByID(ctx, purchaseOrderDetailData.ProductID)
	if err != nil {
		return nil, err
	}

	response := &purchaseorderdetaildomain.PurchaseOrderDetailResponse{
		PurchaseOrderDetail: *purchaseOrderDetailData,
		PurchaseOrder:       *purchaseOrderData,
		Product:             *productData,
	}

	return response, nil
}

func (p *purchaseOrderDetailUseCase) GetByPurchaseOrderID(ctx context.Context, purchaseOrderID string) ([]purchaseorderdetaildomain.PurchaseOrderDetailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	purchaseOrderDetailID, err := primitive.ObjectIDFromHex(purchaseOrderID)
	if err != nil {
		return nil, err
	}

	purchaseOrderDetailData, err := p.purchaseOrderDetailRepository.GetByPurchaseOrderID(ctx, purchaseOrderDetailID)
	if err != nil {
		return nil, err
	}

	var responses []purchaseorderdetaildomain.PurchaseOrderDetailResponse
	responses = make([]purchaseorderdetaildomain.PurchaseOrderDetailResponse, 0, len(purchaseOrderDetailData))
	for _, purchaseOrderDetail := range purchaseOrderDetailData {
		purchaseOrderData, err := p.purchaseOrderRepository.GetByID(ctx, purchaseOrderDetail.PurchaseOrderID)
		if err != nil {
			return nil, err
		}

		productData, err := p.productRepository.GetByID(ctx, purchaseOrderDetail.ProductID)
		if err != nil {
			return nil, err
		}

		response := purchaseorderdetaildomain.PurchaseOrderDetailResponse{
			PurchaseOrderDetail: purchaseOrderDetail,
			PurchaseOrder:       *purchaseOrderData,
			Product:             *productData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (p *purchaseOrderDetailUseCase) CreateOne(ctx context.Context, input *purchaseorderdetaildomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	if err := validate.ValidatePurchaseOrderDetail(input); err != nil {
		return err
	}

	productData, err := p.productRepository.GetByName(ctx, input.Product)
	if err != nil {
		return err
	}

	totalPrice := input.UnitPrice * float64(input.Quantity)

	purchaseOrderDetail := purchaseorderdetaildomain.PurchaseOrderDetail{
		ID:              primitive.NewObjectID(),
		PurchaseOrderID: input.PurchaseOrderID,
		ProductID:       productData.ID,
		Quantity:        input.Quantity,
		UnitPrice:       input.UnitPrice,
		TotalPrice:      totalPrice,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	return p.purchaseOrderDetailRepository.CreateOne(ctx, &purchaseOrderDetail)
}

func (p *purchaseOrderDetailUseCase) UpdateOne(ctx context.Context, id string, input *purchaseorderdetaildomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	if err := validate.ValidatePurchaseOrderDetail(input); err != nil {
		return err
	}

	purchaseOrderDetailID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	productData, err := p.productRepository.GetByName(ctx, input.Product)
	if err != nil {
		return err
	}

	totalPrice := input.UnitPrice * float64(input.Quantity)

	purchaseOrderDetail := &purchaseorderdetaildomain.PurchaseOrderDetail{
		ID:              purchaseOrderDetailID,
		PurchaseOrderID: input.PurchaseOrderID,
		ProductID:       productData.ID,
		Quantity:        input.Quantity,
		UnitPrice:       input.UnitPrice,
		TotalPrice:      totalPrice,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	return p.purchaseOrderDetailRepository.UpdateOne(ctx, purchaseOrderDetail)
}

func (p *purchaseOrderDetailUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	purchaseOrderDetailID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return p.purchaseOrderDetailRepository.DeleteOne(ctx, purchaseOrderDetailID)
}

func (p *purchaseOrderDetailUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]purchaseorderdetaildomain.PurchaseOrderDetailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	purchaseOrderDetailData, err := p.purchaseOrderDetailRepository.GetAllWithPagination(ctx, pagination)
	if err != nil {
		return nil, err
	}

	var responses []purchaseorderdetaildomain.PurchaseOrderDetailResponse
	responses = make([]purchaseorderdetaildomain.PurchaseOrderDetailResponse, 0, len(purchaseOrderDetailData))
	for _, purchaseOrderDetail := range purchaseOrderDetailData {
		purchaseOrderData, err := p.purchaseOrderRepository.GetByID(ctx, purchaseOrderDetail.PurchaseOrderID)
		if err != nil {
			return nil, err
		}

		productData, err := p.productRepository.GetByID(ctx, purchaseOrderDetail.ProductID)
		if err != nil {
			return nil, err
		}

		response := purchaseorderdetaildomain.PurchaseOrderDetailResponse{
			PurchaseOrderDetail: purchaseOrderDetail,
			PurchaseOrder:       *purchaseOrderData,
			Product:             *productData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (p *purchaseOrderDetailUseCase) GetAll(ctx context.Context) ([]purchaseorderdetaildomain.PurchaseOrderDetailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	purchaseOrderDetailData, err := p.purchaseOrderDetailRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []purchaseorderdetaildomain.PurchaseOrderDetailResponse
	responses = make([]purchaseorderdetaildomain.PurchaseOrderDetailResponse, 0, len(purchaseOrderDetailData))
	for _, purchaseOrderDetail := range purchaseOrderDetailData {
		purchaseOrderData, err := p.purchaseOrderRepository.GetByID(ctx, purchaseOrderDetail.PurchaseOrderID)
		if err != nil {
			return nil, err
		}

		productData, err := p.productRepository.GetByID(ctx, purchaseOrderDetail.ProductID)
		if err != nil {
			return nil, err
		}

		response := purchaseorderdetaildomain.PurchaseOrderDetailResponse{
			PurchaseOrderDetail: purchaseOrderDetail,
			PurchaseOrder:       *purchaseOrderData,
			Product:             *productData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
