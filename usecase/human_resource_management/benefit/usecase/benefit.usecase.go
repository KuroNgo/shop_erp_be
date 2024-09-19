package benefit_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	benefitsdomain "shop_erp_mono/domain/human_resource_management/benefits"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/usecase/human_resource_management/benefit/validate"
	"time"
)

type benefitUseCase struct {
	contextTimeout     time.Duration
	benefitRepository  benefitsdomain.IBenefitRepository
	employeeRepository employeesdomain.IEmployeeRepository
}

func NewBenefitUseCase(contextTimeout time.Duration, benefitRepository benefitsdomain.IBenefitRepository, employeeRepository employeesdomain.IEmployeeRepository) benefitsdomain.IBenefitUseCase {
	return &benefitUseCase{contextTimeout: contextTimeout, benefitRepository: benefitRepository, employeeRepository: employeeRepository}
}

func (b *benefitUseCase) CreateOne(ctx context.Context, input *benefitsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	if err := validate.ValidateBenefit(input); err != nil {
		return err
	}

	employeeData, err := b.employeeRepository.GetByEmail(ctx, input.EmployeeEmail)
	if err != nil {
		return err
	}

	benefit := benefitsdomain.Benefit{
		ID:          primitive.NewObjectID(),
		EmployeeID:  employeeData.ID,
		BenefitType: input.BenefitType,
		Amount:      input.Amount,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return b.benefitRepository.CreateOne(ctx, &benefit)
}

func (b *benefitUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	benefitID, _ := primitive.ObjectIDFromHex(id)

	err := b.benefitRepository.DeleteOne(ctx, benefitID)
	if err != nil {
		return err
	}

	return nil
}

func (b *benefitUseCase) UpdateOne(ctx context.Context, id string, input *benefitsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	if err := validate.ValidateBenefit(input); err != nil {
		return err
	}

	benefitID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	employeeData, err := b.employeeRepository.GetByEmail(ctx, input.EmployeeEmail)
	if err != nil {
		return err
	}

	benefit := benefitsdomain.Benefit{
		EmployeeID:  employeeData.ID,
		BenefitType: input.BenefitType,
		Amount:      input.Amount,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
	}

	return b.benefitRepository.UpdateOne(ctx, benefitID, &benefit)
}

func (b *benefitUseCase) GetByID(ctx context.Context, id string) (benefitsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	benefitID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return benefitsdomain.Output{}, err
	}

	benefitData, err := b.benefitRepository.GetByID(ctx, benefitID)
	if err != nil {
		return benefitsdomain.Output{}, err
	}

	employeeData, err := b.employeeRepository.GetByID(ctx, benefitData.EmployeeID)
	if err != nil {
		return benefitsdomain.Output{}, err
	}

	output := benefitsdomain.Output{
		Benefit:  benefitData,
		Employee: employeeData,
	}

	return output, nil
}

func (b *benefitUseCase) GetByEmail(ctx context.Context, email string) (benefitsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	employeeData, err := b.employeeRepository.GetByEmail(ctx, email)
	if err != nil {
		return benefitsdomain.Output{}, err
	}

	benefitData, err := b.benefitRepository.GetByEmployeeID(ctx, employeeData.ID)
	if err != nil {
		return benefitsdomain.Output{}, err
	}

	output := benefitsdomain.Output{
		Benefit:  benefitData,
		Employee: employeeData,
	}

	return output, nil
}

func (b *benefitUseCase) GetAll(ctx context.Context) ([]benefitsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	benefitData, err := b.benefitRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []benefitsdomain.Output
	outputs = make([]benefitsdomain.Output, 0, len(benefitData))
	for _, benefit := range benefitData {
		employeeData, err := b.employeeRepository.GetByID(ctx, benefit.EmployeeID)
		if err != nil {
			return nil, err
		}

		output := benefitsdomain.Output{
			Benefit:  benefit,
			Employee: employeeData,
		}

		outputs = append(outputs, output)
	}
	return outputs, nil
}
