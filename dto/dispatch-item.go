package dto

import (
	"gitlab.sudovi.me/erp/inventory-api/data"
)

type DispatchItemDTO struct {
	InventoryId int `json:"inventory_id"`
	DispatchId  int `json:"dispatch_id"`
}

type DispatchItemResponseDTO struct {
	ID          int `json:"id"`
	InventoryId int `json:"inventory_id"`
	DispatchId  int `json:"dispatch_id"`
}

func (dto DispatchItemDTO) ToDispatchItem() *data.DispatchItem {
	return &data.DispatchItem{
		InventoryId: dto.InventoryId,
		DispatchId:  dto.DispatchId,
	}
}

func ToDispatchItemResponseDTO(data data.DispatchItem) DispatchItemResponseDTO {
	return DispatchItemResponseDTO{
		ID:          data.ID,
		InventoryId: data.InventoryId,
		DispatchId:  data.DispatchId,
	}
}

func ToDispatchItemListResponseDTO(dispatch_items []*data.DispatchItem) []DispatchItemResponseDTO {
	dtoList := make([]DispatchItemResponseDTO, len(dispatch_items))
	for i, x := range dispatch_items {
		dtoList[i] = ToDispatchItemResponseDTO(*x)
	}
	return dtoList
}
