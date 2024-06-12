package services

import (
	"context"

	"gitlab.sudovi.me/erp/inventory-api/data"
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
	CreateItem(ctx context.Context, input dto.ItemDTO) (*dto.ItemResponseDTO, error)
	UpdateItem(ctx context.Context, id int, input dto.ItemDTO) (*dto.ItemResponseDTO, error)
	DeleteItem(ctx context.Context, id int) error
	GetItem(id int) (*dto.ItemResponseDTO, error)
	GetItemList(filter dto.InventoryItemFilter) ([]dto.ItemResponseDTO, *uint64, error)
	GetItemListInOrganizationUnit(id int) ([]data.ItemInOrganizationUnit, error)
	GetItemListForReport(input dto.ItemReportFilterDTO) ([]data.ItemReportResponse, error)
}

type AssessmentService interface {
	CreateAssessment(ctx context.Context, input dto.AssessmentDTO) (*dto.AssessmentResponseDTO, error)
	UpdateAssessment(ctx context.Context, id int, input dto.AssessmentDTO) (*dto.AssessmentResponseDTO, error)
	DeleteAssessment(ctx context.Context, id int) error
	GetAssessment(id int) (*dto.AssessmentResponseDTO, error)
	GetAssessmentList() ([]dto.AssessmentResponseDTO, error)
	GetAssessmentbyItemId(id int) ([]dto.AssessmentResponseDTO, *uint64, error)
}

type DispatchService interface {
	CreateDispatch(ctx context.Context, input dto.DispatchDTO) (*dto.DispatchResponseDTO, error)
	UpdateDispatch(ctx context.Context, id int, input dto.DispatchDTO) (*dto.DispatchResponseDTO, error)
	DeleteDispatch(ctx context.Context, id int) error
	GetDispatch(id int) (*dto.DispatchResponseDTO, error)
	GetDispatchList(input *dto.GetDispatchListInput) ([]dto.DispatchResponseDTO, *uint64, error)
}

type DispatchItemService interface {
	CreateDispatchItem(input dto.DispatchItemDTO) (*dto.DispatchItemResponseDTO, error)
	UpdateDispatchItem(id int, input dto.DispatchItemDTO) (*dto.DispatchItemResponseDTO, error)
	DeleteDispatchItem(id int) error
	GetDispatchItemList(id int) ([]dto.DispatchItemResponseDTO, error)
	GetDispatchItemListbyStatus(Type *string, DispatchID *int) ([]dto.DispatchItemResponseDTO, error)
	GetItemListOfDispatch(dispatchID int) ([]dto.ItemResponseDTO, error)
}

type LogService interface {
	CreateLog(input dto.LogDTO) (*dto.LogResponseDTO, error)
	UpdateLog(id int, input dto.LogDTO) (*dto.LogResponseDTO, error)
	DeleteLog(id int) error
	GetLog(id int) (*dto.LogResponseDTO, error)
	GetLogList(filter dto.LogFilterDTO) ([]dto.LogResponseDTO, *uint64, error)
}
