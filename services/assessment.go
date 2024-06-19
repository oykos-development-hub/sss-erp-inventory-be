package services

import (
	"context"

	"gitlab.sudovi.me/erp/inventory-api/data"
	"gitlab.sudovi.me/erp/inventory-api/dto"
	"gitlab.sudovi.me/erp/inventory-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"

	newErrors "gitlab.sudovi.me/erp/inventory-api/pkg/errors"
)

type AssessmentServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Assessment
}

func NewAssessmentServiceImpl(app *celeritas.Celeritas, repo data.Assessment) AssessmentService {
	return &AssessmentServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *AssessmentServiceImpl) CreateAssessment(ctx context.Context, input dto.AssessmentDTO) (*dto.AssessmentResponseDTO, error) {
	data := input.ToAssessment()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo assessments insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo assessments get")
	}

	res := dto.ToAssessmentResponseDTO(*data)

	return &res, nil
}

func (h *AssessmentServiceImpl) UpdateAssessment(ctx context.Context, id int, input dto.AssessmentDTO) (*dto.AssessmentResponseDTO, error) {
	data := input.ToAssessment()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo assessments update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo assessments get")
	}

	response := dto.ToAssessmentResponseDTO(*data)

	return &response, nil
}

func (h *AssessmentServiceImpl) DeleteAssessment(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo assessments delete")
	}

	return nil
}

func (h *AssessmentServiceImpl) GetAssessment(id int) (*dto.AssessmentResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo assessments get")
	}
	response := dto.ToAssessmentResponseDTO(*data)

	return &response, nil
}

func (h *AssessmentServiceImpl) GetAssessmentList() ([]dto.AssessmentResponseDTO, error) {
	data, _, err := h.repo.GetAll(nil, nil, nil)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo assessments get all")
	}
	response := dto.ToAssessmentListResponseDTO(data)

	return response, nil
}

func (h *AssessmentServiceImpl) GetAssessmentbyItemId(id int) ([]dto.AssessmentResponseDTO, *uint64, error) {
	cond := up.Cond{
		"inventory_id": id,
	}

	data, total, err := h.repo.GetAll(nil, nil, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, newErrors.Wrap(err, "repo assessments get all")
	}

	if len(data) == 0 {
		//response = dto.ToAssessmentResponseDTO(*data[0])
		return nil, nil, newErrors.Wrap(errors.ErrNotFound, "repo assessments get all")
	}

	response := dto.ToAssessmentListResponseDTO(data)
	return response, total, nil
}
