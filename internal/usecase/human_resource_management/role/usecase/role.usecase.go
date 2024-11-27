package role_usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	employees_domain "shop_erp_mono/internal/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	"shop_erp_mono/internal/usecase/human_resource_management/role/validate"
	"shop_erp_mono/pkg/shared/constant"
	"strconv"
	"time"
)

type roleUseCase struct {
	contextTimeout     time.Duration
	roleRepository     roledomain.IRoleRepository
	userRepository     userdomain.IUserRepository
	employeeRepository employees_domain.IEmployeeRepository
	cache              *bigcache.BigCache
}

func NewRoleUseCase(contextTimeout time.Duration, roleRepository roledomain.IRoleRepository, userRepository userdomain.IUserRepository,
	employeeRepository employees_domain.IEmployeeRepository, cacheTTL time.Duration) roledomain.IRoleUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &roleUseCase{contextTimeout: contextTimeout, cache: cache, roleRepository: roleRepository,
		employeeRepository: employeeRepository, userRepository: userRepository}
}

func (r *roleUseCase) checkRoleLevelFromUserID(ctx context.Context, idUser string) (int, error) {
	userID, err := primitive.ObjectIDFromHex(idUser)
	if err != nil {
		log.Printf("error of convert id to hex %s", err)
	}

	userData, err := r.userRepository.GetByID(ctx, userID)
	if err != nil {
		return 0, err
	}

	employeeData, err := r.employeeRepository.GetByID(ctx, userData.EmployeeID)
	if err != nil {
		return 0, err
	}

	roleLevel, err := r.roleRepository.GetByID(ctx, employeeData.RoleID)
	if err != nil {
		return 0, err
	}

	return roleLevel.Level, nil
}

func (r *roleUseCase) CreateOne(ctx context.Context, input *roledomain.Input, idUser string) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	if err := validate.Role(input); err != nil {
		return err
	}

	role := &roledomain.Role{
		ID:          primitive.NewObjectID(),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := r.roleRepository.CreateOne(ctx, role)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleUseCase) GetByName(ctx context.Context, name string) (roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.cache.Get(name)
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}
	if data != nil {
		var role roledomain.Output
		err = json.Unmarshal(data, &role)
		if err != nil {
			log.Printf("failed to unmarshal roles cache: %v", err)
		}
	}

	roleData, err := r.roleRepository.GetByName(ctx, name)
	if err != nil {
		return roledomain.Output{}, err
	}

	countEmployee, err := r.employeeRepository.CountEmployeeByRoleID(ctx, roleData.ID)
	if err != nil {
		return roledomain.Output{}, err
	}

	output := roledomain.Output{
		Role:           roleData,
		NumberOfPeople: countEmployee,
	}

	data, err = json.Marshal(output)
	err = r.cache.Set(name, data)
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}

	return output, nil
}

func (r *roleUseCase) GetByID(ctx context.Context, id string) (roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.cache.Get(id)
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}
	if data != nil {
		var role roledomain.Output
		err = json.Unmarshal(data, &role)
		if err != nil {
			log.Printf("failed to unmarshal roles cache: %v", err)
		}
	}

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

	data, err = json.Marshal(output)
	err = r.cache.Set(id, data)
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}

	return output, nil
}

func (r *roleUseCase) GetByEnable(ctx context.Context, enable int) ([]roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.cache.Get(strconv.Itoa(enable))
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}
	if data != nil {
		var role roledomain.Output
		err = json.Unmarshal(data, &role)
		if err != nil {
			log.Printf("failed to unmarshal roles cache: %v", err)
		}
	}

	roleData, err := r.roleRepository.GetByEnable(ctx, enable)
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

	data, err = json.Marshal(outputs)
	err = r.cache.Set(strconv.Itoa(enable), data)
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}

	return outputs, nil
}

func (r *roleUseCase) GetByLevel(ctx context.Context, level int) ([]roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.cache.Get(strconv.Itoa(level))
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}
	if data != nil {
		var role roledomain.Output
		err = json.Unmarshal(data, &role)
		if err != nil {
			log.Printf("failed to unmarshal roles cache: %v", err)
		}
	}

	roleData, err := r.roleRepository.GetByLevel(ctx, level)
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

	data, err = json.Marshal(outputs)
	err = r.cache.Set(strconv.Itoa(level), data)
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}

	return outputs, nil
}

func (r *roleUseCase) GetByStatus(ctx context.Context, status string) ([]roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.cache.Get(status)
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}
	if data != nil {
		var role roledomain.Output
		err = json.Unmarshal(data, &role)
		if err != nil {
			log.Printf("failed to unmarshal roles cache: %v", err)
		}
	}

	roleData, err := r.roleRepository.GetByStatus(ctx, status)
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

	data, err = json.Marshal(outputs)
	err = r.cache.Set(status, data)
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}

	return outputs, nil
}

func (r *roleUseCase) GetAll(ctx context.Context) ([]roledomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data, err := r.cache.Get("roles")
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}
	if data != nil {
		var role roledomain.Output
		err = json.Unmarshal(data, &role)
		if err != nil {
			log.Printf("failed to unmarshal roles cache: %v", err)
		}
	}

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

	data, err = json.Marshal(outputs)
	err = r.cache.Set("roles", data)
	if err != nil {
		log.Printf("failed to get roles cache: %v", err)
	}

	return outputs, nil
}

func (r *roleUseCase) UpdateOne(ctx context.Context, id string, input *roledomain.Input, idUser string) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	if err := validate.Role(input); err != nil {
		return err
	}

	roleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	role := &roledomain.Role{
		ID:          roleID,
		Name:        input.Name,
		Description: input.Description,
		Level:       input.Level,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = r.roleRepository.UpdateOne(ctx, roleID, role)
	if err != nil {
		return err
	}

	err = r.cache.Delete("roles")
	if err != nil {
		log.Printf("failed to delete roles cache: %v", err)
	}

	err = r.cache.Delete(strconv.Itoa(input.Level))
	if err != nil {
		log.Printf("failed to delete roles cache: %v", err)
	}

	err = r.cache.Delete(id)
	if err != nil {
		log.Printf("failed to delete roles cache: %v", err)
	}

	return nil
}

func (r *roleUseCase) UpdateStatus(ctx context.Context, id string, status string, idUser string) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	roleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = r.roleRepository.UpdateStatus(ctx, roleID, status)
	if err != nil {
		return err
	}

	err = r.cache.Delete("roles")
	if err != nil {
		log.Printf("failed to delete roles cache: %v", err)
	}

	err = r.cache.Delete(status)
	if err != nil {
		log.Printf("failed to delete roles cache: %v", err)
	}

	err = r.cache.Delete(id)
	if err != nil {
		log.Printf("failed to delete roles cache: %v", err)
	}

	return nil
}

func (r *roleUseCase) DeleteOne(ctx context.Context, id string, idUser string) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	roleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	count, err := r.employeeRepository.CountEmployeeByRoleID(ctx, roleID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(constant.MsgDataDeletionFailure)
	}

	err = r.roleRepository.DeleteOne(ctx, roleID)
	if err != nil {
		return err
	}

	err = r.cache.Delete("roles")
	if err != nil {
		log.Printf("failed to delete roles cache: %v", err)
	}

	err = r.cache.Delete(id)
	if err != nil {
		log.Printf("failed to delete roles cache: %v", err)
	}

	return nil
}

func (r *roleUseCase) DeleteSoft(ctx context.Context, id string, idUser string) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	roleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	count, err := r.employeeRepository.CountEmployeeByRoleID(ctx, roleID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(constant.MsgDataDeletionFailure)
	}

	err = r.roleRepository.DeleteSoft(ctx, roleID)
	if err != nil {
		return err
	}

	err = r.cache.Delete("roles")
	if err != nil {
		log.Printf("failed to delete roles cache: %v", err)
	}

	err = r.cache.Delete(id)
	if err != nil {
		log.Printf("failed to delete roles cache: %v", err)
	}

	return nil
}

func (r *roleUseCase) CountRole(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	return r.roleRepository.CountRole(ctx)
}

func (r *roleUseCase) Lifecycle(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
