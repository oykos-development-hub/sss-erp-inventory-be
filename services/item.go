package services

import (
	"context"

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

func (h *ItemServiceImpl) CreateItem(ctx context.Context, input dto.ItemDTO) (*dto.ItemResponseDTO, error) {
	data := input.ToItem()

	id, err := h.repo.Insert(ctx, *data)
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

func (h *ItemServiceImpl) UpdateItem(ctx context.Context, id int, input dto.ItemDTO) (*dto.ItemResponseDTO, error) {
	data := input.ToItem()
	data.ID = id

	err := h.repo.Update(ctx, *data)
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

func (h *ItemServiceImpl) DeleteItem(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
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

	filterOnData := data.InventoryItemFilter{
		ID:                        filter.ID,
		Type:                      filter.Type,
		ClassTypeID:               filter.ClassTypeID,
		OfficeID:                  filter.OfficeID,
		Search:                    filter.Search,
		ContractID:                filter.ContractID,
		DeprecationTypeID:         filter.DeprecationTypeID,
		InvoiceArticleID:          filter.InvoiceArticleID,
		ArticleID:                 filter.ArticleID,
		SourceOrganizationUnitID:  filter.SourceOrganizationUnitID,
		OrganizationUnitID:        filter.OrganizationUnitID,
		SerialNumber:              filter.SerialNumber,
		InventoryNumber:           filter.InventoryNumber,
		Location:                  filter.Location,
		Page:                      filter.Page,
		Size:                      filter.Size,
		CurrentOrganizationUnitID: filter.CurrentOrganizationUnitID,
		SourceType:                filter.SourceType,
		IsExternalDonation:        filter.IsExternalDonation,
		Expire:                    filter.Expire,
		Status:                    filter.Status,
		TypeOfImmovableProperty:   filter.TypeOfImmovableProperty,
	}

	data, total, err := h.repo.GetAll(filterOnData)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}

	response := dto.ToItemListResponseDTO(data)
	return response, total, nil
}

func (h *ItemServiceImpl) GetItemListInOrganizationUnit(id int) ([]data.ItemInOrganizationUnit, error) {
	data, err := h.repo.GetAllInOrgUnit(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	return data, nil
}

func (h *ItemServiceImpl) GetItemListForReport(input dto.ItemReportFilterDTO) ([]data.ItemReportResponse, error) {
	data, err := h.repo.GetAllForReport(input.Type, input.SourceType, input.OrganizationUnitID, input.OfficeID, input.Date)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	return data, nil
}
