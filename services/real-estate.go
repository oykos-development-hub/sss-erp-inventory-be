package services

import (
	"gitlab.sudovi.me/erp/inventory-api/data"
	"gitlab.sudovi.me/erp/inventory-api/dto"
	"gitlab.sudovi.me/erp/inventory-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type RealEstateServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.RealEstate
}

func NewRealEstateServiceImpl(app *celeritas.Celeritas, repo data.RealEstate) RealEstateService {
	return &RealEstateServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *RealEstateServiceImpl) CreateRealEstate(input dto.RealEstateDTO) (*dto.RealEstateResponseDTO, error) {
	data := input.ToRealEstate()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToRealEstateResponseDTO(*data)

	return &res, nil
}

func (h *RealEstateServiceImpl) UpdateRealEstate(id int, input dto.RealEstateDTO) (*dto.RealEstateResponseDTO, error) {
	data := input.ToRealEstate()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToRealEstateResponseDTO(*data)

	return &response, nil
}

func (h *RealEstateServiceImpl) DeleteRealEstate(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *RealEstateServiceImpl) GetRealEstate(id int) (*dto.RealEstateResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToRealEstateResponseDTO(*data)

	return &response, nil
}

func (h *RealEstateServiceImpl) GetRealEstatebyItemId(id int) (*dto.RealEstateResponseDTO, error) {
	cond := up.Cond{
		"item_id": id,
	}
	curr := 1
	data, _, err := h.repo.GetAll(&curr, &curr, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.RealEstateResponseDTO{}

	if len(data) != 0 {
		response = dto.ToRealEstateResponseDTO(*data[0])
		return &response, nil
	}

	return nil, nil
}

func (h *RealEstateServiceImpl) GetRealEstateList(input dto.GetRealEstateListInput) ([]dto.RealEstateResponseDTO, *uint64, error) {
	data, total, err := h.repo.GetAll(input.Page, input.Size, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToRealEstateListResponseDTO(data)

	return response, total, nil
}
