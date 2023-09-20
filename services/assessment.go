package services

import (
	"gitlab.sudovi.me/erp/inventory-api/data"
	"gitlab.sudovi.me/erp/inventory-api/dto"
	"gitlab.sudovi.me/erp/inventory-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
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

func (h *AssessmentServiceImpl) CreateAssessment(input dto.AssessmentDTO) (*dto.AssessmentResponseDTO, error) {
	data := input.ToAssessment()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToAssessmentResponseDTO(*data)

	return &res, nil
}

func (h *AssessmentServiceImpl) UpdateAssessment(id int, input dto.AssessmentDTO) (*dto.AssessmentResponseDTO, error) {
	data := input.ToAssessment()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToAssessmentResponseDTO(*data)

	return &response, nil
}

func (h *AssessmentServiceImpl) DeleteAssessment(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *AssessmentServiceImpl) GetAssessment(id int) (*dto.AssessmentResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToAssessmentResponseDTO(*data)

	return &response, nil
}

func (h *AssessmentServiceImpl) GetAssessmentList() ([]dto.AssessmentResponseDTO, error) {
	data, _, err := h.repo.GetAll(nil, nil, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToAssessmentListResponseDTO(data)

	return response, nil
}

func (h *AssessmentServiceImpl) GetAssessmentbyItemId(id int) (*dto.AssessmentResponseDTO, error) {
	cond := up.Cond{
		"inventory_id": id,
	}

	data, _, err := h.repo.GetAll(nil, nil, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.AssessmentResponseDTO{}

	if len(data) != 0 {
		response = dto.ToAssessmentResponseDTO(*data[0])
		return &response, nil
	}

	return nil, nil
}
