package attendance_usecase

import (
	"context"
	attendancedomain "shop_erp_mono/domain/human_resource_management/attendance"
	"time"
)

type attendanceUseCase struct {
	contextTimeout       time.Duration
	attendanceRepository attendancedomain.IAttendanceRepository
}

func NewAttendanceUseCase(contextTimeout time.Duration, attendanceRepository attendancedomain.IAttendanceRepository) attendancedomain.IAttendanceUseCase {
	return &attendanceUseCase{contextTimeout: contextTimeout, attendanceRepository: attendanceRepository}
}

func (a attendanceUseCase) CreateOne(ctx context.Context, input *attendancedomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err := a.attendanceRepository.CreateOne(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (a attendanceUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err := a.attendanceRepository.DeleteOne(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (a attendanceUseCase) UpdateOne(ctx context.Context, id string, input *attendancedomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err := a.attendanceRepository.UpdateOne(ctx, id, input)
	if err != nil {
		return err
	}

	return nil
}

func (a attendanceUseCase) GetOneByID(ctx context.Context, id string) (attendancedomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	data, err := a.attendanceRepository.GetOneByID(ctx, id)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	return data, nil
}

func (a attendanceUseCase) GetOneByEmail(ctx context.Context, email string) (attendancedomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	data, err := a.attendanceRepository.GetOneByEmail(ctx, email)
	if err != nil {
		return attendancedomain.Output{}, err
	}

	return data, nil
}

func (a attendanceUseCase) GetAll(ctx context.Context) ([]attendancedomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	data, err := a.attendanceRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
