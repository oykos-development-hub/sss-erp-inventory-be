package data

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/lib/pq"
	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/inventory-api/contextutil"
	newErrors "gitlab.sudovi.me/erp/inventory-api/pkg/errors"
)

// Item struct
type Item struct {
	ID                           int           `db:"id,omitempty"`
	ArticleID                    *int          `db:"article_id"`
	InvoiceArticleID             *int          `db:"invoice_article_id"`
	Type                         string        `db:"type"`
	ClassTypeID                  int           `db:"class_type_id"`
	DepreciationTypeID           int           `db:"depreciation_type_id"`
	SupplierID                   int           `db:"supplier_id"`
	InvoiceID                    *int          `db:"invoice_id"`
	DonorID                      *int          `db:"donor_id"`
	SerialNumber                 *string       `db:"serial_number"`
	InventoryNumber              *string       `db:"inventory_number"`
	Title                        string        `db:"title"`
	Abbreviation                 *string       `db:"abbreviation"`
	InternalOwnership            bool          `db:"internal_ownership"`
	OfficeID                     int           `db:"office_id"`
	ContractID                   int           `db:"contract_id"`
	Location                     *string       `db:"location"`
	TargetUserProfileID          *int          `db:"target_user_profile_id"`
	OrganizationUnitID           *int          `db:"organization_unit_id"`
	TargetOrganizationUnitID     *int          `db:"target_organization_unit_id"`
	Unit                         *string       `db:"unit"`
	Amount                       int           `db:"amount"`
	NetPrice                     *float64      `db:"net_price"`
	GrossPrice                   float64       `db:"gross_price"`
	Description                  *string       `db:"description"`
	DateOfPurchase               *time.Time    `db:"date_of_purchase"`
	Inactive                     *time.Time    `db:"inactive"`
	Source                       *string       `db:"source"`
	SourceType                   *string       `db:"source_type"`
	DonorTitle                   *string       `db:"donor_title"`
	InvoiceNumber                *string       `db:"invoice_number"`
	Active                       bool          `db:"active"`
	DeactivationDescription      *string       `db:"deactivation_description"`
	DateOfAssessment             *time.Time    `db:"date_of_assessment"`
	PriceOfAssessment            *int          `db:"price_of_assessment"`
	LifetimeOfAssessmentInMonths *int          `db:"lifetime_of_assessment_in_months"`
	DonationDescription          *string       `db:"donation_description"`
	DonationFiles                pq.Int64Array `db:"donation_files"`
	CreatedAt                    time.Time     `db:"created_at,omitempty"`
	UpdatedAt                    time.Time     `db:"updated_at"`
	InvoiceFileID                *int          `db:"invoice_file_id"`
	FileID                       *int          `db:"file_id"`
	DeactivationFileID           *int          `db:"deactivation_file_id"`
	IsExternalDonation           bool          `db:"is_external_donation"`
	Owner                        *string       `db:"owner"`
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

type ItemInOrganizationUnit struct {
	ItemID      int   `json:"item_id"`
	ReversID    int   `json:"revers_id"`
	ReturnID    int   `json:"return_id"`
	MovementsID []int `json:"movements_id"`
}

type ExcelItem struct {
	Article            Item         `json:"article"`
	FirstAmortization  Assessment   `json:"first_amortization"`
	SecondAmortization Assessment   `json:"second_amortization"`
	Dispatch           Dispatch     `json:"dispatch"`
	ReversDispatch     Dispatch     `json:"revers_dispatch"`
	DispatchItem       DispatchItem `json:"dispatch_item"`
	ReversDispatchItem DispatchItem `json:"revers_dispatch_item"`
}

type ExcelPS2Item struct {
	OrganizationUnitID int    `json:"organization_unit_id"`
	InventoryNumber    string `json:"inventory_number"`
	OfficeID           int    `json:"office_id"`
	DateOfDispatch     string `json:"date_of_dispatch"`
}

// Table returns the table name
func (t *Item) Table() string {
	return "items"
}

// GetAll gets all records from the database, using Upper
func (t *Item) GetAll(filter InventoryItemFilter) ([]*Item, *uint64, error) {
	var items []*Item
	var query string

	if filter.Status != nil && *filter.Status == "Arhiva" {
		query = buildQueryForArchive(filter)
	} else {
		query = buildQuery(filter)
	}
	rows, err := Upper.SQL().Query(query)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper exec")
	}
	defer rows.Close()

	for rows.Next() {
		var itemID int
		err = rows.Scan(&itemID)

		if err != nil {
			return nil, nil, newErrors.Wrap(err, "upper scan")
		}

		item, err := t.Get(itemID)
		if err != nil {
			return nil, nil, newErrors.Wrap(err, "item get")
		}

		items = append(items, item)
	}

	if filter.Status != nil && *filter.Status == "Arhiva" {
		query = buildQueryForArchiveTotal(filter)
	} else {
		query = buildQueryForTotal(filter)
	}

	rows, err = Upper.SQL().Query(query)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper exec")
	}
	defer rows.Close()

	var total uint64
	for rows.Next() {
		var count int
		err = rows.Scan(&count)

		if err != nil {
			return nil, nil, newErrors.Wrap(err, "upper scan")
		}

		total = uint64(count)

	}

	return items, &total, err
}

func buildQuery(filter InventoryItemFilter) string {
	selectPart := `SELECT DISTINCT(i.id)
	FROM items i
	LEFT JOIN (
		SELECT MAX(d.id) AS max_dispatch_id, di.inventory_id AS inventory_id
		FROM dispatches d
		JOIN dispatch_items di ON d.id = di.dispatch_id
		GROUP BY di.inventory_id
	) AS max_dispatches ON i.id = max_dispatches.inventory_id
	LEFT JOIN (
		SELECT MAX(a.id) AS max_assessment_id, a.inventory_id AS inventory_id
		FROM assessments a
		GROUP BY a.inventory_id
	) AS max_assessments ON i.id = max_assessments.inventory_id
	LEFT JOIN dispatch_items di ON di.inventory_id = i.id
	LEFT JOIN dispatches d ON d.id = di.dispatch_id
	LEFT JOIN assessments a ON a.inventory_id = i.id
	LEFT JOIN real_estates r ON r.item_id = i.id
	WHERE (d.id IS NULL OR d.id = max_dispatches.max_dispatch_id)
		AND (a.id IS NULL OR a.id = max_assessments.max_assessment_id)
	 `
	var conditions string

	if filter.ArticleID != nil {
		articleIDString := strconv.Itoa(*filter.ArticleID)
		conditions = conditions + " and i.article_id = " + articleIDString
	}

	if filter.InvoiceArticleID != nil {
		articleIDString := strconv.Itoa(*filter.InvoiceArticleID)
		conditions = conditions + " and i.invoice_article_id = " + articleIDString
	}

	if filter.ClassTypeID != nil {
		classTypeIDString := strconv.Itoa(*filter.ClassTypeID)
		conditions = conditions + " and i.class_type_id = " + classTypeIDString
	}

	if filter.ContractID != nil {
		contractIDString := strconv.Itoa(*filter.ContractID)
		conditions = conditions + " and i.contract_id = " + contractIDString
	}

	if filter.DeprecationTypeID != nil {
		depreciationTypeIDString := strconv.Itoa(*filter.DeprecationTypeID)
		conditions = conditions + " and i.depreciation_type_id = " + depreciationTypeIDString
	}

	if filter.InventoryNumber != nil {
		conditions = conditions + " and i.inventory_number = '" + *filter.InventoryNumber + "'"
	}

	if filter.IsExternalDonation != nil {
		var externalDonation string
		if *filter.IsExternalDonation {
			externalDonation = "true"
		} else {
			externalDonation = "false"
		}
		conditions = conditions + " and i.is_external_donation = " + externalDonation
	}

	if filter.OfficeID != nil {
		officeIDString := strconv.Itoa(*filter.OfficeID)
		conditions = conditions + " and i.office_id = " + officeIDString
	}

	if filter.OrganizationUnitID != nil {
		organizationUnitIDString := strconv.Itoa(*filter.OrganizationUnitID)
		conditions = conditions + " and (i.organization_unit_id = " + organizationUnitIDString + " or i.target_organization_unit_id = " + organizationUnitIDString + " )"
	}

	if filter.SerialNumber != nil {
		conditions = conditions + " and i.serial_number = '" + *filter.SerialNumber + "'"
	}

	if filter.SourceOrganizationUnitID != nil {
		sourceOrganizationUnitIDString := strconv.Itoa(*filter.SourceOrganizationUnitID)
		conditions = conditions + " and i.organization_unit_id = " + sourceOrganizationUnitIDString
	}

	if filter.Search != nil {
		conditions = conditions + " and (i.inventory_number like '" + *filter.Search + "%' or i.title like '" + *filter.Search + "%' )"
	}

	if filter.Type != nil {
		conditions = conditions + " and i.type = '" + *filter.Type + "'"
	}

	if filter.TypeOfImmovableProperty != nil {
		conditions = conditions + " and r.type_id = '" + *filter.TypeOfImmovableProperty + "'"
	}

	if filter.Status != nil {
		currentOrganizationUnitIDString := strconv.Itoa(filter.CurrentOrganizationUnitID)
		switch *filter.Status {
		case "Otpisano":
			conditions = conditions + " and i.active = false "
		case "Prihvaćeno":
			conditions = conditions + " and (d.type = 'revers' and d.is_accepted = true and i.organization_unit_id =" + currentOrganizationUnitIDString + " ) "
		case "Poslato":
			conditions = conditions + " and (d.type = 'revers' and d.is_accepted = false and i.organization_unit_id =" + currentOrganizationUnitIDString + " ) "
		case "Zaduženo":
			conditions = conditions + " and d.type = 'allocation' "
		case "Povraćaj":
			conditions = conditions + " and d.is_accepted = false and  d.type = 'return-revers' and d.source_organization_unit_id = " + currentOrganizationUnitIDString
		case "Nezaduženo":
			conditions = conditions + ` and NOT (
				(d.type = 'revers' AND NOT d.is_accepted) 
				OR ((i.target_organization_unit_id != 0 AND i.target_organization_unit_id = ` + currentOrganizationUnitIDString + `)
					 OR (d.type = 'revers' AND d.is_accepted AND i.organization_unit_id = ` + currentOrganizationUnitIDString + `)) 
				OR (d.type = 'allocation')
				OR (d.type = 'return-revers' AND d.source_organization_unit_id = ` + currentOrganizationUnitIDString + `) 
			)
			AND i.active = TRUE `
		}

	}

	if filter.SourceType != nil {
		currentOrganizationUnitIDString := strconv.Itoa(filter.CurrentOrganizationUnitID)
		switch *filter.SourceType {
		case "NS1":
			conditions = conditions + " and (i.type = 'immovable' and (i.organization_unit_id = " + currentOrganizationUnitIDString + " and i.is_external_donation = false))"
		case "NS2":
			conditions = conditions + " and (i.type = 'immovable' and (i.target_organization_unit_id = " + currentOrganizationUnitIDString + " or i.is_external_donation = true)) "
		case "PS1":
			conditions = conditions + " and (i.type = 'movable' and (i.organization_unit_id = " + currentOrganizationUnitIDString + " ))"
		case "PS2":
			conditions = conditions + " and (i.type = 'movable' and i.target_organization_unit_id = " + currentOrganizationUnitIDString + ") "
		}
	}

	if filter.Expire != nil {
		conditions = conditions + " and NOW() > a.date_of_assessment + interval '1 year' * a.estimated_duration "
	}

	if filter.Page != nil && filter.Size != nil {
		pageString := strconv.Itoa(*filter.Page)
		sizeString := strconv.Itoa(*filter.Size)
		conditions = conditions + "order by i.id limit " + sizeString + " offset (" + pageString + " - 1) * " + sizeString
	}

	return selectPart + conditions
}

func buildQueryForArchive(filter InventoryItemFilter) string {
	currentOrganizationUnitIDString := strconv.Itoa(filter.CurrentOrganizationUnitID)
	query := `SELECT di.inventory_id
	FROM dispatch_items di
	JOIN dispatches d1 ON di.dispatch_id = d1.id AND d1.type = 'revers'
	WHERE EXISTS (
		SELECT 1
		FROM dispatches d2
		JOIN dispatch_items di2 ON d2.id = di2.dispatch_id
		WHERE d2.type = 'return-revers' AND d2.is_accepted = true AND d1.target_organization_unit_id = d2.source_organization_unit_id 
		AND di2.inventory_id = di.inventory_id and d1.target_organization_unit_id = ` + currentOrganizationUnitIDString + ` 
	) `

	return query
}

func buildQueryForArchiveTotal(filter InventoryItemFilter) string {
	currentOrganizationUnitIDString := strconv.Itoa(filter.CurrentOrganizationUnitID)
	query := `SELECT count(*)
	FROM dispatch_items di
	JOIN dispatches d1 ON di.dispatch_id = d1.id AND d1.type = 'revers'
	WHERE EXISTS (
		SELECT 1
		FROM dispatches d2
		JOIN dispatch_items di2 ON d2.id = di2.dispatch_id
		WHERE d2.type = 'return-revers' AND d2.is_accepted = true AND d1.target_organization_unit_id = d2.source_organization_unit_id 
		AND di2.inventory_id = di.inventory_id and d1.target_organization_unit_id = ` + currentOrganizationUnitIDString + ` 
	) `

	return query
}

func buildQueryForTotal(filter InventoryItemFilter) string {
	selectPart := `SELECT count(*)
	FROM items i
	LEFT JOIN (
		SELECT MAX(d.id) AS max_dispatch_id, di.inventory_id AS inventory_id
		FROM dispatches d
		JOIN dispatch_items di ON d.id = di.dispatch_id
		GROUP BY di.inventory_id
	) AS max_dispatches ON i.id = max_dispatches.inventory_id
	LEFT JOIN (
		SELECT MAX(a.id) AS max_assessment_id, a.inventory_id AS inventory_id
		FROM assessments a
		GROUP BY a.inventory_id
	) AS max_assessments ON i.id = max_assessments.inventory_id
	LEFT JOIN dispatch_items di ON di.inventory_id = i.id
	LEFT JOIN dispatches d ON d.id = di.dispatch_id
	LEFT JOIN assessments a ON a.inventory_id = i.id
	LEFT JOIN real_estates r ON r.item_id = i.id
	WHERE (d.id IS NULL OR d.id = max_dispatches.max_dispatch_id)
		AND (a.id IS NULL OR a.id = max_assessments.max_assessment_id)
	 `
	var conditions string

	if filter.ArticleID != nil {
		articleIDString := strconv.Itoa(*filter.ArticleID)
		conditions = conditions + " and i.article_id = " + articleIDString
	}

	if filter.InvoiceArticleID != nil {
		articleIDString := strconv.Itoa(*filter.InvoiceArticleID)
		conditions = conditions + " and i.invoice_article_id = " + articleIDString
	}

	if filter.ClassTypeID != nil {
		classTypeIDString := strconv.Itoa(*filter.ClassTypeID)
		conditions = conditions + " and i.class_type_id = " + classTypeIDString
	}

	if filter.ContractID != nil {
		contractIDString := strconv.Itoa(*filter.ContractID)
		conditions = conditions + " and i.contract_id = " + contractIDString
	}

	if filter.DeprecationTypeID != nil {
		depreciationTypeIDString := strconv.Itoa(*filter.DeprecationTypeID)
		conditions = conditions + " and i.depreciation_type_id = " + depreciationTypeIDString
	}

	if filter.InventoryNumber != nil {
		conditions = conditions + " and i.inventory_number = '" + *filter.InventoryNumber + "'"
	}

	if filter.IsExternalDonation != nil {
		var externalDonation string
		if *filter.IsExternalDonation {
			externalDonation = "true"
		} else {
			externalDonation = "false"
		}
		conditions = conditions + " and i.is_external_donation = " + externalDonation
	}

	if filter.OfficeID != nil {
		officeIDString := strconv.Itoa(*filter.OfficeID)
		conditions = conditions + " and i.office_id = " + officeIDString
	}

	if filter.OrganizationUnitID != nil {
		organizationUnitIDString := strconv.Itoa(*filter.OrganizationUnitID)
		conditions = conditions + " and (i.organization_unit_id = " + organizationUnitIDString + " or i.target_organization_unit_id = " + organizationUnitIDString + " )"
	}

	if filter.SerialNumber != nil {
		conditions = conditions + " and i.serial_number = '" + *filter.SerialNumber + "'"
	}

	if filter.SourceOrganizationUnitID != nil {
		sourceOrganizationUnitIDString := strconv.Itoa(*filter.SourceOrganizationUnitID)
		conditions = conditions + " and i.organization_unit_id = " + sourceOrganizationUnitIDString
	}

	if filter.Search != nil {
		conditions = conditions + " and (i.inventory_number like '%" + *filter.Search + "%' or i.title like '" + *filter.Search + "%' )"
	}

	if filter.Type != nil {
		conditions = conditions + " and i.type = '" + *filter.Type + "'"
	}

	if filter.TypeOfImmovableProperty != nil {
		conditions = conditions + " and r.type_id = '" + *filter.TypeOfImmovableProperty + "'"
	}

	if filter.Status != nil {
		currentOrganizationUnitIDString := strconv.Itoa(filter.CurrentOrganizationUnitID)
		switch *filter.Status {
		case "Otpisano":
			conditions = conditions + " and i.active = false "
		case "Prihvaćeno":
			conditions = conditions + " and (d.type = 'revers' and d.is_accepted = true and i.organization_unit_id =" + currentOrganizationUnitIDString + " ) "
		case "Poslato":
			conditions = conditions + " and (d.type = 'revers' and d.is_accepted = false and i.organization_unit_id =" + currentOrganizationUnitIDString + " ) "
		case "Zaduženo":
			conditions = conditions + " and d.type = 'allocation' "
		case "Povraćaj":
			conditions = conditions + " and d.is_accepted = false and  d.type = 'return-revers' and d.source_organization_unit_id = " + currentOrganizationUnitIDString
		case "Nezaduženo":
			conditions = conditions + ` and NOT (
				(d.type = 'revers' AND NOT d.is_accepted) 
				OR ((i.target_organization_unit_id != 0 AND i.target_organization_unit_id = ` + currentOrganizationUnitIDString + `)
					 OR (d.type = 'revers' AND d.is_accepted AND i.organization_unit_id = ` + currentOrganizationUnitIDString + `)) 
				OR (d.type = 'allocation')
				OR (d.type = 'return-revers' AND d.source_organization_unit_id = ` + currentOrganizationUnitIDString + `) 
			)
			AND i.active = TRUE `
		}

	}

	if filter.SourceType != nil {
		currentOrganizationUnitIDString := strconv.Itoa(filter.CurrentOrganizationUnitID)
		switch *filter.SourceType {
		case "NS1":
			conditions = conditions + " and (i.type = 'immovable' and (i.organization_unit_id = " + currentOrganizationUnitIDString + " and i.is_external_donation = false))"
		case "NS2":
			conditions = conditions + " and (i.type = 'immovable' and (i.target_organization_unit_id = " + currentOrganizationUnitIDString + " or i.is_external_donation = true)) "
		case "PS1":
			conditions = conditions + " and (i.type = 'movable' and (i.organization_unit_id = " + currentOrganizationUnitIDString + " ))"
		case "PS2":
			conditions = conditions + " and (i.type = 'movable' and i.target_organization_unit_id = " + currentOrganizationUnitIDString + ") "
		}
	}

	if filter.Expire != nil {
		conditions = conditions + " and NOW() > a.date_of_assessment + interval '1 year' * a.estimated_duration "
	}

	return selectPart + conditions
}

// Get gets one record from the database, by id, using Upper
func (t *Item) Get(id int) (*Item, error) {
	var one Item
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Item) Update(ctx context.Context, m Item) error {
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return newErrors.Wrap(err, "context get user id")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec")
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(m.ID)
		if err := res.Update(&m); err != nil {
			return newErrors.Wrap(err, "upper update")
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *Item) Delete(ctx context.Context, id int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return newErrors.Wrap(err, "context get user id")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec")
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return newErrors.Wrap(err, "upper delete")
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using
func (t *Item) Insert(ctx context.Context, m Item) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return 0, newErrors.Wrap(err, "context get user id")
	}

	var id int

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec")
		}

		collection := sess.Collection(t.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(m); err != nil {
			return newErrors.Wrap(err, "upper insert")
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get items by org unit
func (t *Item) GetAllInOrgUnit(id int) ([]ItemInOrganizationUnit, error) {
	var items []ItemInOrganizationUnit
	query1 := `select i.id, d.id from items i, dispatches d, dispatch_items di 
			   where i.id = di.inventory_id and d.id = di.dispatch_id and d.target_organization_unit_id = $1;`

	rows, err := Upper.SQL().Query(query1, id)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper exec")
	}
	defer rows.Close()

	for rows.Next() {
		var item ItemInOrganizationUnit
		err = rows.Scan(&item.ItemID, &item.ReversID)
		if err != nil {
			return nil, newErrors.Wrap(err, "upper scan")
		}

		query2 := `select d.id from dispatches d, dispatch_items i 
				 where d.id > $1 and d.type = 'return-revers' and i.inventory_id = $2 
				 and d.id = i.dispatch_id order by d.id;`
		rowDispatch, err := Upper.SQL().Query(query2, item.ReversID, item.ItemID)
		if err != nil {
			return nil, newErrors.Wrap(err, "upper exec")
		}
		defer rowDispatch.Close()

		for rowDispatch.Next() {
			err = rowDispatch.Scan(&item.ReturnID)
			if err != nil {
				return nil, newErrors.Wrap(err, "upper scan")
			}

			query3 := `select d.id from items i, dispatches d, dispatch_items di 
					 where i.id = di.inventory_id and d.id = di.dispatch_id and i.id = $1 
					 and d.id >= $2 and d.id <= $3;
			`
			rowMedium, err := Upper.SQL().Query(query3, item.ItemID, item.ReversID, item.ReturnID)
			if err != nil {
				return nil, newErrors.Wrap(err, "upper exec")
			}
			defer rowMedium.Close()

			for rowMedium.Next() {
				var id int
				err = rowMedium.Scan(&id)
				if err != nil {
					return nil, newErrors.Wrap(err, "upper scan")
				}
				item.MovementsID = append(item.MovementsID, id)
			}

			items = append(items, item)
			break
		}
	}

	return items, err
}

type ItemReportResponse struct {
	ID                       int     `json:"id"`
	Title                    string  `json:"title"`
	SourceType               string  `json:"source_type"`
	InventoryNumber          string  `json:"inventory_number"`
	OfficeID                 int     `json:"office_id"`
	SourceOrganizationUnitID int     `json:"source_organization_unit_id"`
	TargetOrganizationUnitID int     `json:"target_organization_unit_id"`
	ProcurementPrice         float64 `json:"procurement_price"`
	LostValue                float64 `json:"lost_value"`
	Price                    float64 `json:"price"`
	Date                     string  `json:"date"`
	DateOfPurchase           string  `json:"date_of_purchase"`
}

func (t *Item) GetAllForReport(itemType *string, sourceType *string, organizationUnitID *int, officeID *int, date *string) ([]ItemReportResponse, error) {
	var items []ItemReportResponse
	//NS1 && PS1 items in moment 'date'
	query1 := `SELECT i.id, i.type, i.is_external_donation
	  FROM items i
	  WHERE i.organization_unit_id = $1 and i.date_of_purchase <= $2 and i.is_external_donation = false`

	rows1, err := Upper.SQL().Query(query1, *organizationUnitID, *date)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper exec")
	}
	defer rows1.Close()

	for rows1.Next() {
		var item ItemReportResponse
		var sourceTypeQuery string
		var isDonation bool
		err = rows1.Scan(&item.ID, &sourceTypeQuery, &isDonation)
		if err != nil {
			return nil, newErrors.Wrap(err, "upper scan")
		}
		//checks was item donation
		query3 := ` SELECT i.id
		FROM items i, dispatches d, dispatch_items di
		WHERE i.id = $1 and date >= $2 and d.id = di.dispatch_id 
		and d.type = 'convert' and di.inventory_id = i.id;`

		rows3, err := Upper.SQL().Query(query3, item.ID, *date)
		if err != nil {
			return nil, newErrors.Wrap(err, "upper exec")
		}
		defer rows3.Close()

		var donationID int

		for rows3.Next() {
			err = rows3.Scan(&donationID)
			if err != nil {
				return nil, newErrors.Wrap(err, "upper scan")
			}
		}

		if sourceTypeQuery == "movable" {
			if isDonation || donationID != 0 {
				item.SourceType = "PS2"
			} else {
				item.SourceType = "PS1"
			}
		} else {
			if isDonation || donationID != 0 {
				item.SourceType = "NS2"
			} else {
				item.SourceType = "NS1"
			}
		}
		if (sourceType == nil || (*sourceType == item.SourceType)) &&
			(itemType == nil || (*itemType == sourceTypeQuery)) {
			items = append(items, item)
		}
	}

	/*
		//PS2 items in moment 'date'
		query2 := `WITH RankedDispatches AS (
			SELECT i.id, d.source_organization_unit_id, d.type, d.date,
			ROW_NUMBER() OVER (PARTITION BY i.id ORDER BY d.date DESC) AS rn
			FROM items i
			JOIN dispatch_items di ON i.id = di.inventory_id
			JOIN dispatches d ON di.dispatch_id = d.id
			WHERE ((d.type = 'revers' AND d.target_organization_unit_id = $1)
			OR (d.type = 'return-revers' AND d.source_organization_unit_id = $1))
			  AND d.date < $2
		  )
		  SELECT id, source_organization_unit_id
		  FROM RankedDispatches
		  WHERE rn <= 1 and type = 'revers';`

		rows2, err := Upper.SQL().Query(query2, *organizationUnitID, *date)
		if err != nil {
			return nil, err
		}
		defer rows2.Close()

		for rows2.Next() {
			var item ItemReportResponse
			err = rows2.Scan(&item.ID, &item.SourceOrganizationUnitID)
			if err != nil {
				return nil, err
			}
			item.SourceType = "PS2"
			if (sourceType == nil || (sourceType != nil && *sourceType == item.SourceType)) &&
				(itemType == nil || (itemType != nil && *itemType == item.SourceType)) {
				items = append(items, item)
			}
		}
	*/
	//checks office of item in moment date
	query4 := `WITH RankedDispatches AS (
			SELECT i.id, d.type, d.office_id, d.date
			FROM items i
			JOIN dispatch_items di ON i.id = di.inventory_id
			JOIN dispatches d ON di.dispatch_id = d.id
			WHERE (d.type = 'allocation' OR d.type = 'return' OR d.type='created')
			AND date <= $1 AND i.id = $2
			ORDER BY date DESC)
			  SELECT office_id
			  FROM RankedDispatches
			  LIMIT 1;`
	var currentResponse []ItemReportResponse

	for _, item := range items {
		rows4, err := Upper.SQL().Query(query4, *date, item.ID)
		if err != nil {
			return nil, newErrors.Wrap(err, "upper exec")
		}
		defer rows4.Close()

		var officeIDQuery int
		for rows4.Next() {
			err = rows4.Scan(&officeIDQuery)
			if err != nil {
				return nil, newErrors.Wrap(err, "upper scan")
			}
		}
		if officeID == nil || (*officeID == officeIDQuery) {
			item.OfficeID = officeIDQuery
			currentResponse = append(currentResponse, item)
		}
	}

	items = currentResponse

	//makni obavezno office_id - zakucavanje da rade ove popisne liste njihove
	query5 := `SELECT i.id, i.title, i.inventory_number, a.gross_price_difference,
		 a.estimated_duration, a.date_of_assessment, i.date_of_purchase
		FROM items i
		JOIN assessments a ON i.id = a.inventory_id
		WHERE (i.id, a.id) IN (
		  SELECT i.id,  MAX(a.id) AS max_date
		  FROM items i
		  JOIN assessments a ON i.id = a.inventory_id
		  WHERE a.date_of_assessment <= $1 and i.id = $2
		  GROUP BY i.id)
		  LIMIT 1;`

	for i := 0; i < len(items); i++ {
		rows5, err := Upper.SQL().Query(query5, *date, items[i].ID)
		if err != nil {
			return nil, newErrors.Wrap(err, "upper exec")
		}
		defer rows5.Close()

		for rows5.Next() {
			var estimatedDuration int
			var dateOfAssessment string
			err = rows5.Scan(&items[i].ID, &items[i].Title, &items[i].InventoryNumber, &items[i].ProcurementPrice,
				&estimatedDuration, &dateOfAssessment, &items[i].DateOfPurchase)
			if err != nil {
				return nil, newErrors.Wrap(err, "upper scan")
			}

			dateOfAssessmentTime, err := time.Parse(time.RFC3339, dateOfAssessment)
			if err != nil {
				return nil, err
			}
			dateTime, err := time.Parse(time.RFC3339, *date)
			if err != nil {
				return nil, err
			}

			years := dateTime.Year() - dateOfAssessmentTime.Year()
			months := int(dateTime.Month()) - int(dateOfAssessmentTime.Month())
			if months < 0 {
				years--
				months += 12
			}

			months = years*12 + months

			totalConsumption := float64(0)
			var percentage float64
			if estimatedDuration != 0 {
				percentage = float64(100) / float64(estimatedDuration)

				monthlyConsumption := items[i].ProcurementPrice * percentage / 100 / 12
				for i := 0; i < months; i++ {
					totalConsumption += monthlyConsumption
				}
			}

			items[i].LostValue = totalConsumption
			items[i].Price = items[i].ProcurementPrice - items[i].LostValue

			if items[i].Price < 0 {
				items[i].Price = 0
				items[i].LostValue = items[i].ProcurementPrice
			}
		}
	}

	return items, err
}

func (t *Item) CreateExcelItem(ctx context.Context, items []ExcelItem) error {

	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return newErrors.Wrap(err, "context get user id")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec")
		}

		collectionItems := sess.Collection(t.Table())
		collectionAssessments := sess.Collection("assessments")
		collectionDispatches := sess.Collection("dispatches")
		collectionDispatchItems := sess.Collection("dispatch_items")

		for _, item := range items {
			var res up.InsertResult
			var err error

			item.Article.CreatedAt = time.Now()

			if res, err = collectionItems.Insert(item.Article); err != nil {
				return newErrors.Wrap(err, "upper insert - insert article")
			}

			id := getInsertId(res.ID())

			item.FirstAmortization.InventoryID = id
			item.FirstAmortization.CreatedAt = time.Now()

			if _, err = collectionAssessments.Insert(item.FirstAmortization); err != nil {
				return newErrors.Wrap(err, "upper insert - insert first amortization")
			}

			if item.SecondAmortization.GrossPriceDifference != 0 {
				item.SecondAmortization.CreatedAt = time.Now()
				item.SecondAmortization.InventoryID = id

				if _, err = collectionAssessments.Insert(item.SecondAmortization); err != nil {
					return newErrors.Wrap(err, "upper insert - insert second amortization")
				}
			}

			if item.ReversDispatch.TargetOrganizationUnitID != 0 {
				var resDispatch up.InsertResult
				item.ReversDispatch.CreatedAt = time.Now()
				if resDispatch, err = collectionDispatches.Insert(item.ReversDispatch); err != nil {
					return newErrors.Wrap(err, "upper insert - insert revers dispatch")
				}

				resDispatchID := getInsertId(resDispatch.ID())

				item.ReversDispatchItem.DispatchId = resDispatchID
				item.ReversDispatchItem.InventoryId = id

				if _, err = collectionDispatchItems.Insert(item.ReversDispatchItem); err != nil {
					return newErrors.Wrap(err, "upper insert - insert revers dispatch item")
				}
			} else {
				var resDispatch up.InsertResult
				item.Dispatch.CreatedAt = time.Now()
				if resDispatch, err = collectionDispatches.Insert(item.Dispatch); err != nil {
					return newErrors.Wrap(err, "upper insert - insert dispatch")
				}

				resDispatchID := getInsertId(resDispatch.ID())

				item.DispatchItem.DispatchId = resDispatchID
				item.DispatchItem.InventoryId = id

				if _, err = collectionDispatchItems.Insert(item.DispatchItem); err != nil {
					return newErrors.Wrap(err, "upper insert - insert dispatch item")
				}
			}

		}

		return nil
	})

	return err
}

func (t *Item) CreatePS2ExcelItem(ctx context.Context, items []ExcelPS2Item) error {

	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return newErrors.Wrap(err, "context get user id")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec")
		}

		collectionDispatches := sess.Collection("dispatches")
		collectionDispatchItems := sess.Collection("dispatch_items")

		for _, item := range items {
			var articleItem Item

			queryForItem := `select i.id from items i where i.target_organization_unit_id = $1 and i.inventory_number = $2 LIMIT 1;`
			queryForUpdate := `update items set office_id = $1 where id = $2`

			rows1, err := Upper.SQL().Query(queryForItem, item.OrganizationUnitID, item.InventoryNumber)
			if err != nil {
				return newErrors.Wrap(err, "upper exec")
			}
			defer rows1.Close()

			for rows1.Next() {
				err = rows1.Scan(&articleItem.ID)
				if err != nil {
					return newErrors.Wrap(err, "upper scan")
				}
			}

			if articleItem.ID != 0 {
				_, err := Upper.SQL().Exec(queryForUpdate, item.OfficeID, articleItem.ID)
				if err != nil {
					return newErrors.Wrap(err, "upper exec")
				}

				dateOfDispatch, err := time.Parse(time.RFC3339, item.DateOfDispatch)

				if err != nil {
					return newErrors.Wrap(err, "time parse")
				}

				dispatch := Dispatch{
					Type:       "allocation",
					IsAccepted: true,
					OfficeID:   &item.OfficeID,
					Date:       &dateOfDispatch,
					CreatedAt:  time.Now(),
				}

				var resDispatch up.InsertResult

				if resDispatch, err = collectionDispatches.Insert(dispatch); err != nil {
					return newErrors.Wrap(err, "upper insert - insert dispatch")
				}

				resDispatchID := getInsertId(resDispatch.ID())

				dispatchItem := DispatchItem{
					InventoryId: articleItem.ID,
					DispatchId:  resDispatchID,
				}

				if _, err = collectionDispatchItems.Insert(dispatchItem); err != nil {
					return newErrors.Wrap(err, "upper insert - insert dispatch item")
				}

				queryForDispatch := `select d.id from dispatches d left join dispatch_items di on di.dispatch_id = d.id where di.inventory_id = $1 and d.type = 'revers'`
				queryForUpdateDispatch := `update dispatches set date = $1 where id = $2`

				rows3, err := Upper.SQL().Query(queryForDispatch, articleItem.ID)
				if err != nil {
					return newErrors.Wrap(err, "upper exec")
				}
				defer rows3.Close()

				var currentDispatch Dispatch

				for rows3.Next() {
					err = rows3.Scan(&currentDispatch.ID)
					if err != nil {
						return newErrors.Wrap(err, "upper scan")
					}
				}

				if currentDispatch.ID != 0 {
					_, err := Upper.SQL().Exec(queryForUpdateDispatch, item.DateOfDispatch, currentDispatch.ID)
					if err != nil {
						return newErrors.Wrap(err, "upper exec")
					}
				}

			}
		}

		return nil
	})

	return err
}
