package services

import (
	"fmt"

	"github.com/upper/db/v4"
	up "github.com/upper/db/v4"

	"gitlab.sudovi.me/erp/inventory-api/data"
	"gitlab.sudovi.me/erp/inventory-api/dto"
	"gitlab.sudovi.me/erp/inventory-api/errors"

	"github.com/oykos-development-hub/celeritas"
)

type ItemServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Item
}

func NewItemServiceImpl(app *celeritas.Celeritas, repo data.Item) ItemService {
	return &ItemServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *ItemServiceImpl) CreateItem(input dto.ItemDTO) (*dto.ItemResponseDTO, error) {
	data := input.ToItem()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToItemResponseDTO(*data)

	return &res, nil
}

func (h *ItemServiceImpl) UpdateItem(id int, input dto.ItemDTO) (*dto.ItemResponseDTO, error) {
	data := input.ToItem()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToItemResponseDTO(*data)

	return &response, nil
}

func (h *ItemServiceImpl) DeleteItem(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *ItemServiceImpl) GetItem(id int) (*dto.ItemResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToItemResponseDTO(*data)

	return &response, nil
}

func (h *ItemServiceImpl) GetItemList(filter dto.InventoryItemFilter) ([]dto.ItemResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}

	if filter.ID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"id": *filter.ID})
	}

	if filter.Type != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"type ILIKE": *filter.Type})
	}

	if filter.ClassTypeID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"class_type_id": *filter.ClassTypeID})
	}

	if filter.OfficeID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"office_id": *filter.OfficeID})
	}

	if filter.SourceType != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"source_type ILIKE": *filter.SourceType})
	}

	if filter.DeprecationTypeID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"depreciation_type_id": *filter.DeprecationTypeID})
	}

	if filter.OrganizationUnitID != nil {
		orgUnit := up.Or(
			db.Cond{"organization_unit_id ILIKE": *filter.OrganizationUnitID},
			db.Cond{"target_organization_unit_id ILIKE": *filter.OrganizationUnitID},
		)
		conditionAndExp = up.And(conditionAndExp, orgUnit)
	}

	//bozo reko pretrazujemo samo naslov
	if filter.Search != nil {
		likeCondition := fmt.Sprintf("%%%s%%", *filter.Search)
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"title ILIKE": likeCondition})
	}

	data, total, err := h.repo.GetAll(filter.Page, filter.Size, conditionAndExp)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}

	response := dto.ToItemListResponseDTO(data)
	return response, total, nil
}
