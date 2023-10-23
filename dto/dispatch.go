package dto

import (
	"time"

	"gitlab.sudovi.me/erp/inventory-api/data"
)

type GetDispatchListInput struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	Id                 *int    `json:"id"`
	Type               *string `json:"type"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Accepted           *bool   `json:"accepted"`
	InventoryType      *string `json:"inventory_type"`
}

type DispatchDTO struct {
	InventoryType            string     `json:"inventory_type"`
	Type                     string     `json:"type"`
	SourceUserProfileID      int        `json:"source_user_profile_id"`
	TargetUserProfileID      *int       `json:"target_user_profile_id"`
	SourceOrganizationUnitID int        `json:"source_organization_unit_id"`
	TargetOrganizationUnitID int        `json:"target_organization_unit_id"`
	IsAccepted               bool       `json:"is_accepted"`
	Date                     *time.Time `json:"date"`
	SerialNumber             *string    `json:"serial_number"`
	OfficeID                 *int       `json:"office_id"`
	DispatchDescription      *string    `json:"dispatch_description"`
	FileID                   *int       `json:"file_id"`
}

type DispatchResponseDTO struct {
	ID                       int        `json:"id"`
	InventoryType            string     `json:"inventory_type"`
	Type                     string     `json:"type"`
	SourceUserProfileID      int        `json:"source_user_profile_id"`
	TargetUserProfileID      *int       `json:"target_user_profile_id"`
	SourceOrganizationUnitID int        `json:"source_organization_unit_id"`
	TargetOrganizationUnitID int        `json:"target_organization_unit_id"`
	IsAccepted               bool       `json:"is_accepted"`
	Date                     *time.Time `json:"date"`
	SerialNumber             *string    `json:"serial_number"`
	DispatchDescription      *string    `json:"dispatch_description"`
	OfficeID                 *int       `json:"office_id"`
	FileID                   *int       `json:"file_id"`
	CreatedAt                time.Time  `json:"created_at"`
	UpdatedAt                time.Time  `json:"updated_at"`
}

func (dto DispatchDTO) ToDispatch() *data.Dispatch {
	return &data.Dispatch{
		Type:                     dto.Type,
		InventoryType:            dto.InventoryType,
		SourceUserProfileID:      dto.SourceUserProfileID,
		TargetUserProfileID:      dto.TargetUserProfileID,
		SourceOrganizationUnitID: dto.SourceOrganizationUnitID,
		TargetOrganizationUnitID: dto.TargetOrganizationUnitID,
		OfficeID:                 dto.OfficeID,
		Date:                     dto.Date,
		IsAccepted:               dto.IsAccepted,
		SerialNumber:             dto.SerialNumber,
		DispatchDescription:      dto.DispatchDescription,
		FileID:                   dto.FileID,
	}
}

func ToDispatchResponseDTO(data data.Dispatch) DispatchResponseDTO {
	return DispatchResponseDTO{
		ID:                       data.ID,
		Type:                     data.Type,
		InventoryType:            data.InventoryType,
		SourceUserProfileID:      data.SourceUserProfileID,
		TargetUserProfileID:      data.TargetUserProfileID,
		SourceOrganizationUnitID: data.SourceOrganizationUnitID,
		TargetOrganizationUnitID: data.TargetOrganizationUnitID,
		IsAccepted:               data.IsAccepted,
		OfficeID:                 data.OfficeID,
		Date:                     data.Date,
		SerialNumber:             data.SerialNumber,
		DispatchDescription:      data.DispatchDescription,
		FileID:                   data.FileID,
		CreatedAt:                data.CreatedAt,
		UpdatedAt:                data.UpdatedAt,
	}
}

func ToDispatchListResponseDTO(dispatches []*data.Dispatch) []DispatchResponseDTO {
	dtoList := make([]DispatchResponseDTO, len(dispatches))
	for i, x := range dispatches {
		dtoList[i] = ToDispatchResponseDTO(*x)
	}
	return dtoList
}
