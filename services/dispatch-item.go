package services

import (
	"gitlab.sudovi.me/erp/inventory-api/data"
	"gitlab.sudovi.me/erp/inventory-api/dto"
	"gitlab.sudovi.me/erp/inventory-api/errors"

	"github.com/oykos-development-hub/celeritas"
)

type DispatchItemServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.DispatchItem
}

func NewDispatchItemServiceImpl(app *celeritas.Celeritas, repo data.DispatchItem) DispatchItemService {
	return &DispatchItemServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *DispatchItemServiceImpl) CreateDispatchItem(input dto.DispatchItemDTO) (*dto.DispatchItemResponseDTO, error) {
	data := input.ToDispatchItem()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToDispatchItemResponseDTO(*data)

	return &res, nil
}

func (h *DispatchItemServiceImpl) UpdateDispatchItem(id int, input dto.DispatchItemDTO) (*dto.DispatchItemResponseDTO, error) {
	data := input.ToDispatchItem()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToDispatchItemResponseDTO(*data)

	return &response, nil
}

func (h *DispatchItemServiceImpl) DeleteDispatchItem(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *DispatchItemServiceImpl) GetDispatchItemList(id int) ([]dto.DispatchItemResponseDTO, error) {
	data, err := h.repo.GetAll(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToDispatchItemListResponseDTO(data)

	return response, nil
}

func (h *DispatchItemServiceImpl) GetDispatchItemListbyStatus(Type string) ([]dto.DispatchItemResponseDTO, error) {
	data, err := h.repo.GetAllInv(Type)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToDispatchItemListResponseDTO(data)

	return response, nil
}
