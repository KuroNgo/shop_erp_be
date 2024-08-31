package leave_request_usecase

import (
	"context"
	leaverequestdomain "shop_erp_mono/domain/human_resource_management/leave_request"
	"time"
)

type leaveRequestUseCase struct {
	contextTimeout         time.Duration
	leaveRequestRepository leaverequestdomain.ILeaveRequestRepository
}

func (l leaveRequestUseCase) CreateOne(ctx context.Context, leaveRequest *leaverequestdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	err := l.leaveRequestRepository.CreateOne(ctx, leaveRequest)
	if err != nil {
		return err
	}

	return nil
}

func (l leaveRequestUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	err := l.leaveRequestRepository.DeleteOne(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (l leaveRequestUseCase) UpdateOne(ctx context.Context, leaveRequest *leaverequestdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	err := l.leaveRequestRepository.UpdateOne(ctx, leaveRequest)
	if err != nil {
		return err
	}

	return nil
}

func (l leaveRequestUseCase) GetOneByID(ctx context.Context, id string) (leaverequestdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	data, err := l.leaveRequestRepository.GetOneByID(ctx, id)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	return data, nil
}

func (l leaveRequestUseCase) GetOneByEmailEmployee(ctx context.Context, name string) (leaverequestdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	data, err := l.leaveRequestRepository.GetOneByEmailEmployee(ctx, name)
	if err != nil {
		return leaverequestdomain.Output{}, err
	}

	return data, nil
}

func (l leaveRequestUseCase) GetAll(ctx context.Context) ([]leaverequestdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	data, err := l.leaveRequestRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewLeaveRequestUseCase(contextTimeout time.Duration, leaveRequestRepository leaverequestdomain.ILeaveRequestRepository) leaverequestdomain.ILeaveRequestUseCase {
	return &leaveRequestUseCase{contextTimeout: contextTimeout, leaveRequestRepository: leaveRequestRepository}
}
