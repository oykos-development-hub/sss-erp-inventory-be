package services

import (
	"gitlab.sudovi.me/erp/inventory-api/dto"
)

type BaseService interface {
	RandomString(n int) string
	Encrypt(text string) (string, error)
	Decrypt(crypto string) (string, error)
}

type RealEstateService interface {
	CreateRealEstate(input dto.RealEstateDTO) (*dto.RealEstateResponseDTO, error)
	UpdateRealEstate(id int, input dto.RealEstateDTO) (*dto.RealEstateResponseDTO, error)
	DeleteRealEstate(id int) error
	GetRealEstate(id int) (*dto.RealEstateResponseDTO, error)
	GetRealEstatebyItemId(id int) (*dto.RealEstateResponseDTO, error)
	GetRealEstateList(input dto.GetRealEstateListInput) ([]dto.RealEstateResponseDTO, *uint64, error)
}

type ItemService interface {
	CreateItem(input dto.ItemDTO) (*dto.ItemResponseDTO, error)
	UpdateItem(id int, input dto.ItemDTO) (*dto.ItemResponseDTO, error)
	DeleteItem(id int) error
	GetItem(id int) (*dto.ItemResponseDTO, error)
	GetItemList(filter dto.InventoryItemFilter) ([]dto.ItemResponseDTO, *uint64, error)
}

type AssessmentService interface {
	CreateAssessment(input dto.AssessmentDTO) (*dto.AssessmentResponseDTO, error)
	UpdateAssessment(id int, input dto.AssessmentDTO) (*dto.AssessmentResponseDTO, error)
	DeleteAssessment(id int) error
	GetAssessment(id int) (*dto.AssessmentResponseDTO, error)
	GetAssessmentList() ([]dto.AssessmentResponseDTO, error)
	GetAssessmentbyItemId(id int) (*dto.AssessmentResponseDTO, error)
}
