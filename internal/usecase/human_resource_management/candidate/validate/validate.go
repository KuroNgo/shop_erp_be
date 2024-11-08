package validate

import (
	"errors"
	candidatedomain "shop_erp_mono/internal/domain/human_resource_management/candidate"
)

func Candidate(candidate *candidatedomain.Candidate) error {
	if candidate.Education == nil {
		return errors.New("the candidate's information do not nil")
	}

	if candidate.Email == "" {
		return errors.New("the candidate's information do not nil")
	}

	if candidate.Phone == "" {
		return errors.New("the candidate's information do not nil")
	}

	if candidate.Resume == "" {
		return errors.New("the candidate's information do not nil")
	}

	if candidate.Skills == nil {
		return errors.New("the candidate's information do not nil")
	}

	if candidate.Experience == nil {
		return errors.New("the candidate's information do not nil")
	}

	if candidate.Education == nil {
		return errors.New("the candidate's information do not nil")
	}

	if candidate.Status == "" {
		return errors.New("the candidate's information do not nil")
	}

	return nil
}
