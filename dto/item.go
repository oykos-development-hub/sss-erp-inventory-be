package dto

import (
	"time"

	"gitlab.sudovi.me/erp/inventory-api/data"
)

type ItemDTO struct {
	ArticleID                    *int       `json:"article_id"`
	Type                         string     `json:"type"`
	ClassTypeID                  int        `json:"class_type_id"`
	DepreciationTypeID           int        `json:"depreciation_type_id"`
	SupplierID                   int        `json:"supplier_id"`
	SerialNumber                 *string    `json:"serial_number"`
	InventoryNumber              string     `json:"inventory_number"`
	Title                        string     `json:"title"`
	Abbreviation                 *string    `json:"abbreviation"`
	InternalOwnership            bool       `json:"internal_ownership"`
	OfficeID                     int        `json:"office_id"`
	ContractID                   int        `json:"contract_id"`
	Location                     *string    `json:"location"`
	TargetUserProfileID          *int       `json:"target_user_profile_id"`
	OrganizationUnitID           *int       `json:"organization_unit_id"`
	TargetOrganizationUnitID     *int       `json:"target_organization_unit_id"`
	Unit                         *string    `json:"unit"`
	Amount                       int        `json:"amount"`
	NetPrice                     *float32   `json:"net_price"`
	GrossPrice                   float32    `json:"gross_price"`
	Description                  *string    `json:"description"`
	DateOfPurchase               time.Time  `json:"date_of_purchase"`
	Inactive                     *time.Time `json:"inactive"`
	Source                       *string    `json:"source"`
	SourceType                   *string    `json:"source_type"`
	DonorTitle                   *string    `json:"donor_title"`
	InvoiceNumber                *string    `json:"invoice_number"`
	Active                       bool       `json:"active"`
	DeactivationDescription      *string    `json:"deactivation_description"`
	InvoiceFileID                *int       `json:"invoice_file_id"`
	FileID                       *int       `json:"file_id"`
	DeactivationFileID           *int       `json:"deactivation_file_id"`
	DateOfAssessment             *time.Time `json:"date_of_assessment"`
	PriceOfAssessment            *int       `json:"price_of_assessment"`
	LifetimeOfAssessmentInMonths *int       `json:"lifetime_of_assessment_in_months"`
}

type ItemResponseDTO struct {
	ID                           int        `json:"id"`
	ArticleID                    *int       `json:"article_id"`
	Type                         string     `json:"type"`
	ClassTypeID                  int        `json:"class_type_id"`
	DepreciationTypeID           int        `json:"depreciation_type_id"`
	SupplierID                   int        `json:"supplier_id"`
	SerialNumber                 *string    `json:"serial_number"`
	InventoryNumber              string     `json:"inventory_number"`
	Title                        string     `json:"title"`
	Abbreviation                 *string    `json:"abbreviation"`
	InternalOwnership            bool       `json:"internal_ownership"`
	OfficeID                     int        `json:"office_id"`
	Location                     *string    `json:"location"`
	TargetUserProfileID          *int       `json:"target_user_profile_id"`
	OrganizationUnitID           *int       `json:"organization_unit_id"`
	TargetOrganizationUnitID     *int       `json:"target_organization_unit_id"`
	Unit                         *string    `json:"unit"`
	ContractID                   int        `json:"contract_id"`
	Amount                       int        `json:"amount"`
	NetPrice                     *float32   `json:"net_price"`
	GrossPrice                   float32    `json:"gross_price"`
	Description                  *string    `json:"description"`
	DateOfPurchase               time.Time  `json:"date_of_purchase"`
	Inactive                     *time.Time `json:"inactive"`
	Source                       *string    `json:"source"`
	SourceType                   *string    `json:"source_type"`
	DonorTitle                   *string    `json:"donor_title"`
	InvoiceNumber                *string    `json:"invoice_number"`
	Active                       bool       `json:"active"`
	DeactivationDescription      *string    `json:"deactivation_description"`
	DateOfAssessment             *time.Time `json:"date_of_assessment"`
	PriceOfAssessment            *int       `json:"price_of_assessment"`
	LifetimeOfAssessmentInMonths *int       `json:"lifetime_of_assessment_in_months"`
	CreatedAt                    time.Time  `json:"created_at"`
	UpdatedAt                    time.Time  `json:"updated_at"`
	InvoiceFileID                *int       `json:"invoice_file_id"`
	FileID                       *int       `json:"file_id"`
	DeactivationFileID           *int       `json:"deactivation_file_id"`
}

func (dto ItemDTO) ToItem() *data.Item {
	return &data.Item{
		ArticleID:                    dto.ArticleID,
		Type:                         dto.Type,
		ClassTypeID:                  dto.ClassTypeID,
		DepreciationTypeID:           dto.DepreciationTypeID,
		SupplierID:                   dto.SupplierID,
		SerialNumber:                 dto.SerialNumber,
		InventoryNumber:              dto.InventoryNumber,
		Title:                        dto.Title,
		ContractID:                   dto.ContractID,
		Inactive:                     dto.Inactive,
		Abbreviation:                 dto.Abbreviation,
		InternalOwnership:            dto.InternalOwnership,
		OfficeID:                     dto.OfficeID,
		Location:                     dto.Location,
		TargetUserProfileID:          dto.TargetUserProfileID,
		Unit:                         dto.Unit,
		Amount:                       dto.Amount,
		NetPrice:                     dto.NetPrice,
		GrossPrice:                   dto.GrossPrice,
		Description:                  dto.Description,
		DateOfPurchase:               dto.DateOfPurchase,
		Source:                       dto.Source,
		SourceType:                   dto.SourceType,
		DonorTitle:                   dto.DonorTitle,
		InvoiceNumber:                dto.InvoiceNumber,
		Active:                       dto.Active,
		DeactivationDescription:      dto.DeactivationDescription,
		DateOfAssessment:             dto.DateOfAssessment,
		PriceOfAssessment:            dto.PriceOfAssessment,
		LifetimeOfAssessmentInMonths: dto.LifetimeOfAssessmentInMonths,
		OrganizationUnitID:           dto.OrganizationUnitID,
		TargetOrganizationUnitID:     dto.TargetOrganizationUnitID,
		InvoiceFileID:                dto.InvoiceFileID,
		FileID:                       dto.FileID,
		DeactivationFileID:           dto.DeactivationFileID,
	}
}

func ToItemResponseDTO(data data.Item) ItemResponseDTO {
	return ItemResponseDTO{
		ID:                           data.ID,
		ArticleID:                    data.ArticleID,
		Type:                         data.Type,
		ClassTypeID:                  data.ClassTypeID,
		DepreciationTypeID:           data.DepreciationTypeID,
		SupplierID:                   data.SupplierID,
		Inactive:                     data.Inactive,
		SerialNumber:                 data.SerialNumber,
		InventoryNumber:              data.InventoryNumber,
		Title:                        data.Title,
		ContractID:                   data.ContractID,
		Abbreviation:                 data.Abbreviation,
		InternalOwnership:            data.InternalOwnership,
		OfficeID:                     data.OfficeID,
		Location:                     data.Location,
		TargetUserProfileID:          data.TargetUserProfileID,
		Unit:                         data.Unit,
		Amount:                       data.Amount,
		NetPrice:                     data.NetPrice,
		GrossPrice:                   data.GrossPrice,
		Description:                  data.Description,
		DateOfPurchase:               data.DateOfPurchase,
		Source:                       data.Source,
		SourceType:                   data.SourceType,
		DonorTitle:                   data.DonorTitle,
		InvoiceNumber:                data.InvoiceNumber,
		Active:                       data.Active,
		DeactivationDescription:      data.DeactivationDescription,
		DateOfAssessment:             data.DateOfAssessment,
		PriceOfAssessment:            data.PriceOfAssessment,
		LifetimeOfAssessmentInMonths: data.LifetimeOfAssessmentInMonths,
		OrganizationUnitID:           data.OrganizationUnitID,
		TargetOrganizationUnitID:     data.TargetOrganizationUnitID,
		CreatedAt:                    data.CreatedAt,
		UpdatedAt:                    data.UpdatedAt,
		InvoiceFileID:                data.InvoiceFileID,
		FileID:                       data.FileID,
		DeactivationFileID:           data.DeactivationFileID,
	}
}

type InventoryItemFilter struct {
	ID                 *int    `json:"id"`
	Type               *string `json:"type"`
	ClassTypeID        *int    `json:"class_type_id"`
	OfficeID           *int    `json:"office_id"`
	Search             *string `json:"search"`
	ContractID         *int    `json:"contract_id"`
	SourceType         *string `json:"source_type"`
	DeprecationTypeID  *int    `json:"depreciation_type_id"`
	ArticleID          *int    `json:"article_id"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	SerialNumber       *string `json:"serial_number"`
	InventoryNumber    *string `json:"inventory_number"`
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
}

func ToItemListResponseDTO(items []*data.Item) []ItemResponseDTO {
	dtoList := make([]ItemResponseDTO, len(items))
	for i, x := range items {
		dtoList[i] = ToItemResponseDTO(*x)
	}
	return dtoList
}
