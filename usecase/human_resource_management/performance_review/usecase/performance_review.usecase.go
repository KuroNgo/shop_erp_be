package performance_review_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	employees_domain "shop_erp_mono/domain/human_resource_management/employees"
	performancereviewdomain "shop_erp_mono/domain/human_resource_management/performance_review"
	"shop_erp_mono/usecase/human_resource_management/performance_review/validate"
	"time"
)

type performanceReviewUseCase struct {
	contextTimeout              time.Duration
	performanceReviewRepository performancereviewdomain.IPerformanceReviewRepository
	employeeRepository          employees_domain.IEmployeeRepository
}

func NewPerformanceReviewUseCase(contextTimeout time.Duration, performanceReviewRepository performancereviewdomain.IPerformanceReviewRepository,
	employeeRepository employees_domain.IEmployeeRepository) performancereviewdomain.IPerformanceReviewUseCase {
	return &performanceReviewUseCase{contextTimeout: contextTimeout, performanceReviewRepository: performanceReviewRepository, employeeRepository: employeeRepository}
}

func (p *performanceReviewUseCase) CreateOneWithEmailEmployee(ctx context.Context, input *performancereviewdomain.Input1) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	if err := validate.ValidatePerformanceReviewV1(input); err != nil {
		return err
	}

	employeeData, err := p.employeeRepository.GetByEmail(ctx, input.EmployeeEmail)
	if err != nil {
		return err
	}

	reviewerData, err := p.employeeRepository.GetByEmail(ctx, input.ReviewerEmail)
	if err != nil {
		return err
	}

	performanceReview := &performancereviewdomain.PerformanceReview{
		ID:               primitive.NewObjectID(),
		EmployeeID:       employeeData.ID,
		ReviewerID:       reviewerData.ID,
		ReviewDate:       input.ReviewDate,
		PerformanceScore: input.PerformanceScore,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	return p.performanceReviewRepository.CreateOne(ctx, performanceReview)
}

func (p *performanceReviewUseCase) CreateOneWithIDEmployee(ctx context.Context, input *performancereviewdomain.Input2) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	if err := validate.ValidatePerformanceReviewV2(input); err != nil {
		return err
	}

	employeeID, err := primitive.ObjectIDFromHex(input.EmployeeID)
	if err != nil {
		return err
	}

	reviewerID, err := primitive.ObjectIDFromHex(input.ReviewerID)
	if err != nil {
		return err
	}

	employeeData, err := p.employeeRepository.GetByID(ctx, employeeID)
	if err != nil {
		return err
	}

	reviewerData, err := p.employeeRepository.GetByID(ctx, reviewerID)
	if err != nil {
		return err
	}

	performanceReview := &performancereviewdomain.PerformanceReview{
		ID:               primitive.NewObjectID(),
		EmployeeID:       employeeData.ID,
		ReviewerID:       reviewerData.ID,
		ReviewDate:       input.ReviewDate,
		PerformanceScore: input.PerformanceScore,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	return p.performanceReviewRepository.CreateOne(ctx, performanceReview)
}

func (p *performanceReviewUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	performanceReviewID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return p.performanceReviewRepository.DeleteOne(ctx, performanceReviewID)
}

func (p *performanceReviewUseCase) UpdateOneWithEmailEmployee(ctx context.Context, id string, input *performancereviewdomain.Input1) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	performanceReviewID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	employeeData, err := p.employeeRepository.GetByEmail(ctx, input.EmployeeEmail)
	if err != nil {
		return err
	}

	reviewerData, err := p.employeeRepository.GetByEmail(ctx, input.ReviewerEmail)
	if err != nil {
		return err
	}

	performanceReview := &performancereviewdomain.PerformanceReview{
		EmployeeID:       employeeData.ID,
		ReviewerID:       reviewerData.ID,
		ReviewDate:       input.ReviewDate,
		PerformanceScore: input.PerformanceScore,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	return p.performanceReviewRepository.UpdateOne(ctx, performanceReviewID, performanceReview)
}

func (p *performanceReviewUseCase) UpdateOneWithIDEmployee(ctx context.Context, id string, input *performancereviewdomain.Input2) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	performanceReviewID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	employeeID, err := primitive.ObjectIDFromHex(input.EmployeeID)
	if err != nil {
		return err
	}

	reviewerID, err := primitive.ObjectIDFromHex(input.ReviewerID)
	if err != nil {
		return err
	}

	employeeData, err := p.employeeRepository.GetByID(ctx, employeeID)
	if err != nil {
		return err
	}

	reviewerData, err := p.employeeRepository.GetByID(ctx, reviewerID)
	if err != nil {
		return err
	}

	performanceReview := &performancereviewdomain.PerformanceReview{
		EmployeeID:       employeeData.ID,
		ReviewerID:       reviewerData.ID,
		ReviewDate:       input.ReviewDate,
		PerformanceScore: input.PerformanceScore,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	return p.performanceReviewRepository.UpdateOne(ctx, performanceReviewID, performanceReview)
}

func (p *performanceReviewUseCase) GetByID(ctx context.Context, id string) (performancereviewdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	performanceReviewID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return performancereviewdomain.Output{}, err
	}

	performanceReviewData, err := p.performanceReviewRepository.GetByID(ctx, performanceReviewID)
	if err != nil {
		return performancereviewdomain.Output{}, err
	}

	employeeData, err := p.employeeRepository.GetByID(ctx, performanceReviewData.EmployeeID)
	if err != nil {
		return performancereviewdomain.Output{}, err
	}

	reviewerData, err := p.employeeRepository.GetByID(ctx, performanceReviewData.ReviewerID)
	if err != nil {
		return performancereviewdomain.Output{}, err
	}

	output := performancereviewdomain.Output{
		PerformanceReview: performanceReviewData,
		Employee:          employeeData,
		Reviewer:          reviewerData,
	}

	return output, nil
}

func (p *performanceReviewUseCase) GetByEmailEmployee(ctx context.Context, email string) (performancereviewdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	employeeData, err := p.employeeRepository.GetByEmail(ctx, email)
	if err != nil {
		return performancereviewdomain.Output{}, err
	}

	reviewerData, err := p.employeeRepository.GetByEmail(ctx, email)
	if err != nil {
		return performancereviewdomain.Output{}, err
	}

	performanceReviewData, err := p.performanceReviewRepository.GetByEmployeeID(ctx, employeeData.ID)
	if err != nil {
		return performancereviewdomain.Output{}, err
	}

	output := performancereviewdomain.Output{
		PerformanceReview: performanceReviewData,
		Employee:          employeeData,
		Reviewer:          reviewerData,
	}

	return output, nil
}

func (p *performanceReviewUseCase) GetAll(ctx context.Context) ([]performancereviewdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	performanceReviewData, err := p.performanceReviewRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []performancereviewdomain.Output
	outputs = make([]performancereviewdomain.Output, 0, len(performanceReviewData))
	for _, performanceReview := range performanceReviewData {
		employeeData, err := p.employeeRepository.GetByID(ctx, performanceReview.EmployeeID)
		if err != nil {
			return nil, err
		}

		reviewerData, err := p.employeeRepository.GetByID(ctx, performanceReview.ReviewerID)
		if err != nil {
			return nil, err
		}

		output := performancereviewdomain.Output{
			PerformanceReview: performanceReview,
			Employee:          employeeData,
			Reviewer:          reviewerData,
		}

		outputs = append(outputs, output)
	}
	return outputs, nil
}
