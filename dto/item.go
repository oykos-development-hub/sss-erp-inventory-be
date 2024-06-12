package dto

import (
	"time"

	"github.com/lib/pq"
	"gitlab.sudovi.me/erp/inventory-api/data"
)

type ItemDTO struct {
	ArticleID                    *int          `json:"article_id"`
	InvoiceArticleID             *int          `json:"invoice_article_id"`
	Type                         string        `json:"type"`
	ClassTypeID                  int           `json:"class_type_id"`
	DepreciationTypeID           int           `json:"depreciation_type_id"`
	SupplierID                   int           `json:"supplier_id"`
	InvoiceID                    *int          `json:"invoice_id"`
	DonorID                      *int          `json:"donor_id"`
	SerialNumber                 *string       `json:"serial_number"`
	InventoryNumber              *string       `json:"inventory_number"`
	Title                        string        `json:"title"`
	Abbreviation                 *string       `json:"abbreviation"`
	InternalOwnership            bool          `json:"internal_ownership"`
	OfficeID                     int           `json:"office_id"`
	ContractID                   int           `json:"contract_id"`
	Location                     *string       `json:"location"`
	TargetUserProfileID          *int          `json:"target_user_profile_id"`
	OrganizationUnitID           *int          `json:"organization_unit_id"`
	TargetOrganizationUnitID     *int          `json:"target_organization_unit_id"`
	Unit                         *string       `json:"unit"`
	Amount                       int           `json:"amount"`
	NetPrice                     *float32      `json:"net_price"`
	GrossPrice                   float32       `json:"gross_price"`
	Description                  *string       `json:"description"`
	DateOfPurchase               *time.Time    `json:"date_of_purchase"`
	Inactive                     *time.Time    `json:"inactive"`
	Source                       *string       `json:"source"`
	SourceType                   *string       `json:"source_type"`
	DonorTitle                   *string       `json:"donor_title"`
	IsExternalDonation           bool          `json:"is_external_donation"`
	DonationDescription          *string       `json:"donation_description"`
	DonationFiles                pq.Int64Array `json:"donation_files"`
	InvoiceNumber                *string       `json:"invoice_number"`
	Active                       bool          `json:"active"`
	DeactivationDescription      *string       `json:"deactivation_description"`
	InvoiceFileID                *int          `json:"invoice_file_id"`
	FileID                       *int          `json:"file_id"`
	DeactivationFileID           *int          `json:"deactivation_file_id"`
	DateOfAssessment             *time.Time    `json:"date_of_assessment"`
	PriceOfAssessment            *int          `json:"price_of_assessment"`
	LifetimeOfAssessmentInMonths *int          `json:"lifetime_of_assessment_in_months"`
	Owner                        *string       `json:"owner"`
	AssessmentPrice              *float32      `json:"assessment_price"`
}

type ItemResponseDTO struct {
	ID                           int        `json:"id"`
	ArticleID                    *int       `json:"article_id"`
	InvoiceArticleID             *int       `json:"invoice_article_id"`
	Type                         string     `json:"type"`
	ClassTypeID                  int        `json:"class_type_id"`
	DepreciationTypeID           int        `json:"depreciation_type_id"`
	SupplierID                   int        `json:"supplier_id"`
	DonorID                      *int       `json:"donor_id"`
	InvoiceID                    *int       `json:"invoice_id"`
	SerialNumber                 *string    `json:"serial_number"`
	InventoryNumber              *string    `json:"inventory_number"`
	Title                        string     `json:"title"`
	Abbreviation                 *string    `json:"abbreviation"`
	InternalOwnership            bool       `json:"internal_ownership"`
	OfficeID                     int        `json:"office_id"`
	Location                     *string    `json:"location"`
	TargetUserProfileID          *int       `json:"target_user_profile_id"`
	OrganizationUnitID           *int       `json:"organization_unit_id"`
	TargetOrganizationUnitID     *int       `json:"target_organization_unit_id"`
	Unit                         *string    `json:"unit"`
	IsExternalDonation           bool       `json:"is_external_donation"`
	ContractID                   int        `json:"contract_id"`
	Amount                       int        `json:"amount"`
	NetPrice                     *float32   `json:"net_price"`
	GrossPrice                   float32    `json:"gross_price"`
	Description                  *string    `json:"description"`
	DateOfPurchase               *time.Time `json:"date_of_purchase"`
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
	DonationDescription          *string    `json:"donation_description"`
	DonationFiles                []int      `json:"donation_files"`
	CreatedAt                    time.Time  `json:"created_at"`
	UpdatedAt                    time.Time  `json:"updated_at"`
	InvoiceFileID                *int       `json:"invoice_file_id"`
	FileID                       *int       `json:"file_id"`
	DeactivationFileID           *int       `json:"deactivation_file_id"`
	Owner                        *string    `json:"owner"`
	AssessmentPrice              *float32   `json:"assessment_price"`
}

func (dto ItemDTO) ToItem() *data.Item {
	return &data.Item{
		ArticleID:                    dto.ArticleID,
		InvoiceArticleID:             dto.InvoiceArticleID,
		Type:                         dto.Type,
		ClassTypeID:                  dto.ClassTypeID,
		DepreciationTypeID:           dto.DepreciationTypeID,
		SupplierID:                   dto.SupplierID,
		DonorID:                      dto.DonorID,
		InvoiceID:                    dto.InvoiceID,
		SerialNumber:                 dto.SerialNumber,
		InventoryNumber:              dto.InventoryNumber,
		Title:                        dto.Title,
		ContractID:                   dto.ContractID,
		Inactive:                     dto.Inactive,
		IsExternalDonation:           dto.IsExternalDonation,
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
		DonationDescription:          dto.DonationDescription,
		DonationFiles:                dto.DonationFiles,
		Owner:                        dto.Owner,
		AssessmentPrice:              dto.AssessmentPrice,
	}
}

func ToItemResponseDTO(data data.Item) ItemResponseDTO {

	var donationFiles []int

	for _, fileID := range data.DonationFiles {
		donationFiles = append(donationFiles, int(fileID))
	}

	return ItemResponseDTO{
		ID:                           data.ID,
		InvoiceArticleID:             data.InvoiceArticleID,
		ArticleID:                    data.ArticleID,
		Type:                         data.Type,
		ClassTypeID:                  data.ClassTypeID,
		DepreciationTypeID:           data.DepreciationTypeID,
		SupplierID:                   data.SupplierID,
		DonorID:                      data.DonorID,
		Inactive:                     data.Inactive,
		SerialNumber:                 data.SerialNumber,
		InventoryNumber:              data.InventoryNumber,
		Title:                        data.Title,
		IsExternalDonation:           data.IsExternalDonation,
		ContractID:                   data.ContractID,
		Abbreviation:                 data.Abbreviation,
		InternalOwnership:            data.InternalOwnership,
		OfficeID:                     data.OfficeID,
		InvoiceID:                    data.InvoiceID,
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
		DonationDescription:          data.DonationDescription,
		DonationFiles:                donationFiles,
		Owner:                        data.Owner,
		AssessmentPrice:              data.AssessmentPrice,
	}
}

type InventoryItemFilter struct {
	ID                        *int    `json:"id"`
	Type                      *string `json:"type"`
	ClassTypeID               *int    `json:"class_type_id"`
	OfficeID                  *int    `json:"office_id"`
	Search                    *string `json:"search"`
	ContractID                *int    `json:"contract_id"`
	DeprecationTypeID         *int    `json:"depreciation_type_id"`
	ArticleID                 *int    `json:"article_id"`
	InvoiceArticleID          *int    `json:"invoice_article_id"`
	SourceOrganizationUnitID  *int    `json:"source_organization_unit_id"`
	OrganizationUnitID        *int    `json:"organization_unit_id"`
	SerialNumber              *string `json:"serial_number"`
	InventoryNumber           *string `json:"inventory_number"`
	Location                  *string `json:"location"`
	Page                      *int    `json:"page"`
	Size                      *int    `json:"size"`
	CurrentOrganizationUnitID int     `json:"current_organization_unit_id"`
	SourceType                *string `json:"source_type"`
	IsExternalDonation        *bool   `json:"is_external_donation"`
	Expire                    *bool   `json:"expire"`
	Status                    *string `json:"status"`
	TypeOfImmovableProperty   *string `json:"type_of_immovable_property"`
}

type ItemReportFilterDTO struct {
	Type               *string `json:"type"`
	SourceType         *string `json:"source_type"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	OfficeID           *int    `json:"office_id"`
	Date               *string `json:"date"`
}

func ToItemListResponseDTO(items []*data.Item) []ItemResponseDTO {
	dtoList := make([]ItemResponseDTO, len(items))
	for i, x := range items {
		dtoList[i] = ToItemResponseDTO(*x)
	}
	return dtoList
}
