package data

import (
	"strconv"
	"time"

	"github.com/lib/pq"
	up "github.com/upper/db/v4"
)

// Item struct
type Item struct {
	ID                           int           `db:"id,omitempty"`
	ArticleID                    *int          `db:"article_id"`
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
	NetPrice                     *float32      `db:"net_price"`
	GrossPrice                   float32       `db:"gross_price"`
	Description                  *string       `db:"description"`
	DateOfPurchase               time.Time     `db:"date_of_purchase"`
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
	SourceOrganizationUnitID  *int    `json:"source_organization_unit_id"`
	OrganizationUnitID        *int    `json:"organization_unit_id"`
	SerialNumber              *string `json:"serial_number"`
	InventoryNumber           *string `json:"inventory_number"`
	Location                  *string `json:"location"`
	Page                      *int    `json:"page"`
	Size                      *int    `json:"size"` //dovde
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

// Table returns the table name
func (t *Item) Table() string {
	return "items"
}

// GetAll gets all records from the database, using upper
func (t *Item) GetAll(filter InventoryItemFilter) ([]*Item, *uint64, error) {
	var items []*Item
	query := buildQuery(filter)

	rows, err := upper.SQL().Query(query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var itemID int
		err = rows.Scan(&itemID)

		item, err := t.Get(itemID)
		if err != nil {
			return nil, nil, err
		}

		items = append(items, item)
	}

	total := uint64(2000)

	return items, &total, err
}

func buildQuery(filter InventoryItemFilter) string {
	selectPart := `SELECT i.id
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
		conditions = conditions + " and (i.inventory_number = '" + *filter.Search + "' or i.title = '" + *filter.Search + "' )"
	}

	if filter.Type != nil {
		conditions = conditions + " and i.type = '" + *filter.Type + "'"
	}

	if filter.TypeOfImmovableProperty != nil {
		conditions = conditions + " and r.type = '" + *filter.TypeOfImmovableProperty + "'"
	}

	if filter.Status != nil {
		currentOrganizationUnitIDString := strconv.Itoa(filter.CurrentOrganizationUnitID)
		switch *filter.Status {
		case "Otpisano":
			conditions = conditions + " and i.active = false "
		case "Prihvaćeno":
			conditions = conditions + " and ((i.target_organization_unit_id = " + currentOrganizationUnitIDString + " ) or (d.type = 'revers' and d.is_accepted = true and i.organization_unit_id " + currentOrganizationUnitIDString + " )) "
		case "Zaduženo":
			conditions = conditions + " and d.type = 'allocation' "
		case "Povraćaj":
			conditions = conditions + " and d.type = 'return-revers' and d.source_organization_unit_id = " + currentOrganizationUnitIDString
		case "Nezaduženo":
			conditions = conditions + " and not ((i.active = false) or ((i.target_organization_unit_id = " + currentOrganizationUnitIDString + " ) or (d.type = 'revers' and d.is_accepted = true and i.organization_unit_id " + currentOrganizationUnitIDString + " )) or (d.type = 'allocation') or (d.type = 'return-revers' and d.source_organization_unit_id = " + currentOrganizationUnitIDString + "))"
		}
	}

	if filter.SourceType != nil {
		currentOrganizationUnitIDString := strconv.Itoa(filter.CurrentOrganizationUnitID)
		switch *filter.SourceType {
		case "NS1":
			conditions = conditions + " and (i.type = 'immovable' and (i.organization_unit_id = i.target_organization_unit_id or i.organization_unit_id = " + currentOrganizationUnitIDString + " ))"
		case "NS2":
			conditions = conditions + " and (i.type = 'immovable' and (i.target_organization_unit_id = " + currentOrganizationUnitIDString + " or i.is_external_donation = true)) "
		case "PS1":
			conditions = conditions + " and (i.type = 'movable' and (i.organization_unit_id = i.target_organization_unit_id or i.organization_unit_id = " + currentOrganizationUnitIDString + " ))"
		case "PS2":
			conditions = conditions + " and (i.type = 'movable' and (i.target_organization_unit_id = " + currentOrganizationUnitIDString + " or i.is_external_donation = true)) "
		}
	}

	if filter.Expire != nil {
		conditions = conditions + " and NOW() > a.date_of_assessment + interval '1 year' * a.estimated_duration "
	}

	if filter.Page != nil && filter.Size != nil {
		pageString := strconv.Itoa(*filter.Page)
		sizeString := strconv.Itoa(*filter.Size)
		conditions = conditions + "order by id limit " + sizeString + " offset (" + pageString + " - 1) * " + sizeString
	}

	return selectPart + conditions
}

// Get gets one record from the database, by id, using upper
func (t *Item) Get(id int) (*Item, error) {
	var one Item
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Item) Update(m Item) error {
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *Item) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Item) Insert(m Item) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}

// Get items by org unit
func (t *Item) GetAllInOrgUnit(id int) ([]ItemInOrganizationUnit, error) {
	var items []ItemInOrganizationUnit
	query1 := `select i.id, d.id from items i, dispatches d, dispatch_items di 
			   where i.id = di.inventory_id and d.id = di.dispatch_id and d.target_organization_unit_id = $1;`

	rows, err := upper.SQL().Query(query1, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item ItemInOrganizationUnit
		err = rows.Scan(&item.ItemID, &item.ReversID)
		if err != nil {
			return nil, err
		}

		query2 := `select d.id from dispatches d, dispatch_items i 
				 where d.id > $1 and d.type = 'return-revers' and i.inventory_id = $2 
				 and d.id = i.dispatch_id order by d.id;`
		rowDispatch, err := upper.SQL().Query(query2, item.ReversID, item.ItemID)
		if err != nil {
			return nil, err
		}
		defer rowDispatch.Close()

		for rowDispatch.Next() {
			err = rowDispatch.Scan(&item.ReturnID)
			if err != nil {
				return nil, err
			}

			query3 := `select d.id from items i, dispatches d, dispatch_items di 
					 where i.id = di.inventory_id and d.id = di.dispatch_id and i.id = $1 
					 and d.id >= $2 and d.id <= $3;
			`
			rowMedium, err := upper.SQL().Query(query3, item.ItemID, item.ReversID, item.ReturnID)
			if err != nil {
				return nil, err
			}
			defer rowMedium.Close()

			for rowMedium.Next() {
				var id int
				err = rowMedium.Scan(&id)
				if err != nil {
					return nil, err
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
	ProcurementPrice         float32 `json:"procurement_price"`
	LostValue                float32 `json:"lost_value"`
	Price                    float32 `json:"price"`
	Date                     string  `json:"date"`
	DateOfPurchase           string  `json:"date_of_purchase"`
}

func (t *Item) GetAllForReport(itemType *string, sourceType *string, organizationUnitID *int, officeID *int, date *string) ([]ItemReportResponse, error) {
	var items []ItemReportResponse
	//NS1 && PS1 items in moment 'date'
	query1 := `SELECT i.id, i.type, i.is_external_donation
	  FROM items i
	  WHERE i.organization_unit_id = $1 and i.date_of_purchase < $2`

	rows1, err := upper.SQL().Query(query1, *organizationUnitID, *date)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()

	for rows1.Next() {
		var item ItemReportResponse
		var sourceTypeQuery string
		var isDonation bool
		err = rows1.Scan(&item.ID, &sourceTypeQuery, &isDonation)
		if err != nil {
			return nil, err
		}
		//checks was item donation
		query3 := ` SELECT i.id
		FROM items i, dispatches d, dispatch_items di
		WHERE i.id = $1 and date > $2 and d.id = di.dispatch_id 
		and d.type = 'convert' and di.inventory_id = i.id;`

		rows3, err := upper.SQL().Query(query3, item.ID, *date)
		if err != nil {
			return nil, err
		}
		defer rows3.Close()

		var donationID int

		for rows3.Next() {
			err = rows3.Scan(&donationID)
			if err != nil {
				return nil, err
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
		if (sourceType == nil || (sourceType != nil && *sourceType == item.SourceType)) &&
			(itemType == nil || (itemType != nil && *itemType == sourceTypeQuery)) {
			items = append(items, item)
		}
	}

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

	rows2, err := upper.SQL().Query(query2, *organizationUnitID, *date)
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

	//checks office of item in moment date
	query4 := `WITH RankedDispatches AS (
			SELECT i.id, d.type, d.office_id, d.date
			FROM items i
			JOIN dispatch_items di ON i.id = di.inventory_id
			JOIN dispatches d ON di.dispatch_id = d.id
			WHERE (d.type = 'allocation' OR d.type = 'return' OR d.type='created')
			AND date < $1 AND i.id = $2
			ORDER BY date DESC)
			  SELECT office_id
			  FROM RankedDispatches
			  LIMIT 1;`
	var currentResponse []ItemReportResponse

	for _, item := range items {
		rows4, err := upper.SQL().Query(query4, *date, item.ID)
		if err != nil {
			return nil, err
		}
		defer rows4.Close()

		var officeIDQuery int
		for rows4.Next() {
			err = rows4.Scan(&officeIDQuery)
			if err != nil {
				return nil, err
			}
		}
		if officeID == nil || (officeID != nil && *officeID == officeIDQuery) {
			item.OfficeID = officeIDQuery
			currentResponse = append(currentResponse, item)
		}
	}

	items = currentResponse

	query5 := `SELECT i.id, i.title, i.inventory_number, a.gross_price_difference,
		 a.estimated_duration, a.date_of_assessment, i.date_of_purchase
		FROM items i
		JOIN assessments a ON i.id = a.inventory_id
		WHERE (i.id, a.id) IN (
		  SELECT i.id,  MAX(a.id) AS max_date
		  FROM items i
		  JOIN assessments a ON i.id = a.inventory_id
		  WHERE a.date_of_assessment < $1 and i.id = $2
		  GROUP BY i.id)
		  LIMIT 1;`

	for i := 0; i < len(items); i++ {
		rows5, err := upper.SQL().Query(query5, *date, items[i].ID)
		if err != nil {
			return nil, err
		}
		defer rows5.Close()

		for rows5.Next() {
			var estimatedDuration int
			var dateOfAssessment string
			err = rows5.Scan(&items[i].ID, &items[i].Title, &items[i].InventoryNumber, &items[i].ProcurementPrice,
				&estimatedDuration, &dateOfAssessment, &items[i].DateOfPurchase)
			if err != nil {
				return nil, err
			}
			depreciationRate := 100 / estimatedDuration
			monthlyDepreciationRate := float32(depreciationRate) / 12

			dateOfAssessmentTime, err := time.Parse(time.RFC3339, dateOfAssessment)
			if err != nil {
				return nil, err
			}
			dateTime, err := time.Parse(time.RFC3339, *date)
			if err != nil {
				return nil, err
			}
			sub := dateTime.Sub(dateOfAssessmentTime)
			months := sub.Hours() / 24 / 30
			monthsInt := int(months)

			items[i].Price = items[i].ProcurementPrice - float32(monthsInt)*(items[i].ProcurementPrice*monthlyDepreciationRate/100)
			if items[i].Price < 0 {
				items[i].Price = 0
			}
			items[i].LostValue = items[i].ProcurementPrice - items[i].Price
		}
	}

	return items, err
}
