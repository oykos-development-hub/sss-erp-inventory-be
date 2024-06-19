package services

import (
	"gitlab.sudovi.me/erp/inventory-api/data"
	"gitlab.sudovi.me/erp/inventory-api/dto"
	newErrors "gitlab.sudovi.me/erp/inventory-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
)

type DispatchItemServiceImpl struct {
	App      *celeritas.Celeritas
	repo     data.DispatchItem
	itemRepo data.Item
}

func NewDispatchItemServiceImpl(app *celeritas.Celeritas, repo data.DispatchItem, itemRepo data.Item) DispatchItemService {
	return &DispatchItemServiceImpl{
		App:      app,
		repo:     repo,
		itemRepo: itemRepo,
	}
}

func (h *DispatchItemServiceImpl) CreateDispatchItem(input dto.DispatchItemDTO) (*dto.DispatchItemResponseDTO, error) {
	data := input.ToDispatchItem()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo dispatch item insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo dispatch item get")
	}

	res := dto.ToDispatchItemResponseDTO(*data)

	return &res, nil
}

func (h *DispatchItemServiceImpl) UpdateDispatchItem(id int, input dto.DispatchItemDTO) (*dto.DispatchItemResponseDTO, error) {
	data := input.ToDispatchItem()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo dispatch item update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo dispatch item get")
	}

	response := dto.ToDispatchItemResponseDTO(*data)

	return &response, nil
}

func (h *DispatchItemServiceImpl) DeleteDispatchItem(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo dispatch item delete")
	}

	return nil
}

func (h *DispatchItemServiceImpl) GetDispatchItemList(id int) ([]dto.DispatchItemResponseDTO, error) {
	data, err := h.repo.GetAll(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo dispatch item get all")
	}
	response := dto.ToDispatchItemListResponseDTO(data)

	return response, nil
}

func (h *DispatchItemServiceImpl) GetItemListOfDispatch(dispatchID int) ([]dto.ItemResponseDTO, error) {
	dispatchItems, err := h.repo.GetItemListOfDispatch(dispatchID)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo dispatch item get all")
	}

	var items []*data.Item
	for _, dispatchItem := range dispatchItems {
		item, err := h.itemRepo.Get(dispatchItem.InventoryId)
		if err != nil {
			return nil, newErrors.Wrap(err, "repo dispatches get")
		}
		items = append(items, item)
	}
	response := dto.ToItemListResponseDTO(items)

	return response, nil
}

func (h *DispatchItemServiceImpl) GetDispatchItemListbyStatus(Type *string, DispatchID *int) ([]dto.DispatchItemResponseDTO, error) {
	data, err := h.repo.GetAllInv(Type, DispatchID)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo dispatch item get all")
	}
	response := dto.ToDispatchItemListResponseDTO(data)

	return response, nil
}
