package services

import (
	"gitlab.sudovi.me/erp/inventory-api/data"
	"gitlab.sudovi.me/erp/inventory-api/dto"
	"gitlab.sudovi.me/erp/inventory-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type DispatchServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Dispatch
}

func NewDispatchServiceImpl(app *celeritas.Celeritas, repo data.Dispatch) DispatchService {
	return &DispatchServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *DispatchServiceImpl) CreateDispatch(input dto.DispatchDTO) (*dto.DispatchResponseDTO, error) {
	data := input.ToDispatch()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToDispatchResponseDTO(*data)

	return &res, nil
}

func (h *DispatchServiceImpl) UpdateDispatch(id int, input dto.DispatchDTO) (*dto.DispatchResponseDTO, error) {
	data := input.ToDispatch()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToDispatchResponseDTO(*data)

	return &response, nil
}

func (h *DispatchServiceImpl) DeleteDispatch(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *DispatchServiceImpl) GetDispatch(id int) (*dto.DispatchResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToDispatchResponseDTO(*data)

	return &response, nil
}

func (h *DispatchServiceImpl) GetDispatchList(input *dto.GetDispatchListInput) ([]dto.DispatchResponseDTO, *uint64, error) {

	conditionAndExp := &up.AndExpr{}

	if input.Id != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"id": *input.Id})
	}

	if input.Type != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"type ILIKE": *input.Type})
	}

	if input.OrganizationUnitID != nil {
		orgUnit := up.Or(
			up.Cond{"source_organization_unit_id": *input.OrganizationUnitID},
			up.Cond{"target_organization_unit_id": *input.OrganizationUnitID},
		)
		conditionAndExp = up.And(conditionAndExp, orgUnit)
	}

	if input.Accepted != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"is_accepted": *input.Accepted})
	}

	if input.InventoryType != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"inventory_type": *input.InventoryType})
	}

	data, total, err := h.repo.GetAll(input.Page, input.Size, conditionAndExp)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}

	response := dto.ToDispatchListResponseDTO(data)
	return response, total, nil
}
