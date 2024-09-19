package role_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"shop_erp_mono/usecase/human_resource_management/role/validate"
	"time"
)

type roleUseCase struct {
	contextTimeout time.Duration
	roleRepository roledomain.IRoleRepository
}

func NewRoleUseCase(contextTimeout time.Duration, roleRepository roledomain.IRoleRepository) roledomain.IRoleUseCase {
	return &roleUseCase{contextTimeout: contextTimeout, roleRepository: roleRepository}
}

func (r *roleUseCase) CreateOne(ctx context.Context, input *roledomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	if err := validate.ValidateRole(input); err != nil {
		return err
	}

	role := &roledomain.Role{
		ID:          primitive.NewObjectID(),
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return r.roleRepository.CreateOne(ctx, role)
}

func (r *roleUseCase) GetByTitle(ctx context.Context, title string) (roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	roleData, err := r.roleRepository.GetByTitle(ctx, title)
	if err != nil {
		return roledomain.Output{}, err
	}

	output := roledomain.Output{
		Role: roleData,
	}

	return output, nil
}

func (r *roleUseCase) GetByID(ctx context.Context, id string) (roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	roleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return roledomain.Output{}, err
	}

	roleData, err := r.roleRepository.GetByID(ctx, roleID)
	if err != nil {
		return roledomain.Output{}, err
	}

	output := roledomain.Output{
		Role: roleData,
	}
	return output, nil
}

func (r *roleUseCase) GetAll(ctx context.Context) ([]roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	roleData, err := r.roleRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []roledomain.Output
	outputs = make([]roledomain.Output, 0, len(roleData))
	for _, role := range roleData {
		output := roledomain.Output{
			Role: role,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}

func (r *roleUseCase) UpdateOne(ctx context.Context, id string, input *roledomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	if err := validate.ValidateRole(input); err != nil {
		return err
	}

	roleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	role := &roledomain.Role{
		ID:          roleID,
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return r.roleRepository.UpdateOne(ctx, roleID, role)
}

func (r *roleUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	roleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return r.roleRepository.DeleteOne(ctx, roleID)
}
