package dto

import (
	"time"

	"gitlab.sudovi.me/erp/inventory-api/data"
)

type GetRealEstateListInput struct {
	Page *int `json:"page"`
	Size *int `json:"size"`
}

type RealEstateDTO struct {
	Title                    string  `json:"title"`
	ItemID                   int     `json:"item_id" validate:"required"`
	TypeID                   string  `json:"type_id" validate:"required"`
	SquareArea               float64 `json:"square_area"`
	LandSerialNumber         string  `json:"land_serial_number" validate:"required"`
	EstateSerialNumber       string  `json:"estate_serial_number"`
	OwnershipType            string  `json:"ownership_type" `
	OwnershipScope           string  `json:"ownership_scope"`
	OwnershipInvestmentScope string  `json:"ownership_investment_scope"`
	LimitationsDescription   string  `json:"limitations_description"`
	LimitationID             bool    `json:"limitation_id"`
	PropertyDocument         string  `json:"property_document"`
	Document                 string  `json:"document"`
	FileID                   int     `json:"file_id"`
}

type RealEstateResponseDTO struct {
	ID                       int       `json:"id"`
	ItemID                   int       `json:"item_id"`
	TypeID                   string    `json:"type_id"`
	Title                    string    `json:"title"`
	SquareArea               float64   `json:"square_Area"`
	LandSerialNumber         string    `json:"land_serial_number"`
	EstateSerialNumber       string    `json:"estate_serial_number"`
	OwnershipType            string    `json:"ownership_type"`
	OwnershipScope           string    `json:"ownership_scope"`
	OwnershipInvestmentScope string    `json:"ownership_investment_scope"`
	LimitationsDescription   string    `json:"limitations_description"`
	LimitationID             bool      `json:"limitation_id"`
	PropertyDocument         string    `json:"property_document"`
	Document                 string    `json:"document"`
	FileID                   int       `json:"file_id"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

func (dto RealEstateDTO) ToRealEstate() *data.RealEstate {
	return &data.RealEstate{
		Title:                    dto.Title,
		ItemID:                   dto.ItemID,
		TypeID:                   dto.TypeID,
		SquareArea:               dto.SquareArea,
		LandSerialNumber:         dto.LandSerialNumber,
		EstateSerialNumber:       dto.EstateSerialNumber,
		OwnershipType:            dto.OwnershipType,
		OwnershipScope:           dto.OwnershipScope,
		OwnershipInvestmentScope: dto.OwnershipInvestmentScope,
		LimitationsDescription:   dto.LimitationsDescription,
		LimitationID:             dto.LimitationID,
		PropertyDocument:         dto.PropertyDocument,
		Document:                 dto.Document,
		FileID:                   dto.FileID,
	}
}

func ToRealEstateResponseDTO(data data.RealEstate) RealEstateResponseDTO {
	return RealEstateResponseDTO{
		ID:                       data.ID,
		ItemID:                   data.ItemID,
		Title:                    data.Title,
		TypeID:                   data.TypeID,
		SquareArea:               data.SquareArea,
		LandSerialNumber:         data.LandSerialNumber,
		EstateSerialNumber:       data.EstateSerialNumber,
		OwnershipType:            data.OwnershipType,
		OwnershipScope:           data.OwnershipScope,
		OwnershipInvestmentScope: data.OwnershipInvestmentScope,
		LimitationsDescription:   data.LimitationsDescription,
		LimitationID:             data.LimitationID,
		PropertyDocument:         data.PropertyDocument,
		Document:                 data.Document,
		FileID:                   data.FileID,
		CreatedAt:                data.CreatedAt,
		UpdatedAt:                data.UpdatedAt,
	}
}

func ToRealEstateListResponseDTO(realestates []*data.RealEstate) []RealEstateResponseDTO {
	dtoList := make([]RealEstateResponseDTO, len(realestates))
	for i, x := range realestates {
		dtoList[i] = ToRealEstateResponseDTO(*x)
	}
	return dtoList
}
