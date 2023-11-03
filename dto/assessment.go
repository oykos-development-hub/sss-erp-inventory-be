package dto

import (
	"time"

	"gitlab.sudovi.me/erp/inventory-api/data"
)

type AssessmentDTO struct {
	InventoryID          int        `json:"inventory_id"`
	Active               bool       `json:"active"`
	DepreciationTypeID   int        `json:"depreciation_type_id"`
	UserProfileID        *int       `json:"user_profile_id"`
	GrossPriceNew        float32    `json:"gross_price_new"`
	GrossPriceDifference float32    `json:"gross_price_difference"`
	DateOfAssessment     *time.Time `json:"date_of_assessment"`
	FileID               *int       `json:"file_id,omitempty"`
	Type                 string     `json:"type"`
}

type AssessmentResponseDTO struct {
	ID                   int        `json:"id"`
	InventoryID          int        `json:"inventory_id"`
	Active               bool       `json:"active"`
	DepreciationTypeID   int        `json:"depreciation_type_id"`
	UserProfileID        *int       `json:"user_profile_id"`
	GrossPriceNew        float32    `json:"gross_price_new"`
	GrossPriceDifference float32    `json:"gross_price_difference"`
	DateOfAssessment     *time.Time `json:"date_of_assessment"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at,omitempty"`
	FileID               *int       `json:"file_id,omitempty"`
	Type                 string     `json:"type"`
}

func (dto AssessmentDTO) ToAssessment() *data.Assessment {
	return &data.Assessment{
		InventoryID:          dto.InventoryID,
		Active:               dto.Active,
		DepreciationTypeID:   dto.DepreciationTypeID,
		UserProfileID:        dto.UserProfileID,
		GrossPriceNew:        dto.GrossPriceNew,
		GrossPriceDifference: dto.GrossPriceDifference,
		DateOfAssessment:     dto.DateOfAssessment,
		FileID:               dto.FileID,
		Type:                 dto.Type,
	}
}

func ToAssessmentResponseDTO(data data.Assessment) AssessmentResponseDTO {
	return AssessmentResponseDTO{
		ID:                   data.ID,
		InventoryID:          data.InventoryID,
		Active:               data.Active,
		DepreciationTypeID:   data.DepreciationTypeID,
		UserProfileID:        data.UserProfileID,
		GrossPriceNew:        data.GrossPriceNew,
		GrossPriceDifference: data.GrossPriceDifference,
		DateOfAssessment:     data.DateOfAssessment,
		CreatedAt:            data.CreatedAt,
		UpdatedAt:            data.UpdatedAt,
		FileID:               data.FileID,
		Type:                 data.Type,
	}
}

func ToAssessmentListResponseDTO(assessments []*data.Assessment) []AssessmentResponseDTO {
	dtoList := make([]AssessmentResponseDTO, len(assessments))
	for i, x := range assessments {
		dtoList[i] = ToAssessmentResponseDTO(*x)
	}
	return dtoList
}
