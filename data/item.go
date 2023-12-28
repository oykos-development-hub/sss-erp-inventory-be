package data

import (
	"encoding/json"
	"fmt"
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
func (t *Item) GetAll(page *int, size *int, condition *up.AndExpr) ([]*Item, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*Item
	var res up.Result

	if condition != nil {
		res = collection.Find(condition)
	} else {
		res = collection.Find()
	}
	total, err := res.Count()
	if err != nil {
		return nil, nil, err
	}

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
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
		WHERE i.id = $1 and d.created_at > $2 and d.id = di.dispatch_id 
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

	jsonData, _ := json.Marshal(items)
	fmt.Println(string(jsonData))

	//PS2 items in moment 'date'
	query2 := `WITH RankedDispatches AS (
		SELECT i.id, d.source_organization_unit_id, d.type,
		ROW_NUMBER() OVER (PARTITION BY i.id ORDER BY d.created_at DESC) AS rn
		FROM items i
		JOIN dispatch_items di ON i.id = di.inventory_id
		JOIN dispatches d ON di.dispatch_id = d.id
		WHERE ((d.type = 'revers' AND d.target_organization_unit_id = $1)
		OR (d.type = 'return-revers' AND d.source_organization_unit_id = $1))
		  AND d.created_at < $2
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
		if (sourceType == nil || *sourceType == item.SourceType) && (itemType != nil || *itemType == "movable") {
			items = append(items, item)
		}
	}
	fmt.Println("---------------------------------------------------------")
	jsonData, _ = json.Marshal(items)
	fmt.Println(string(jsonData))

	if officeID != nil {
		var currentItems []ItemReportResponse

		//checks office of item in moment date
		query4 := `WITH RankedDispatches AS (
			SELECT i.id,
			ROW_NUMBER() OVER (PARTITION BY i.id ORDER BY d.created_at DESC) AS rn
			FROM items i
			JOIN dispatch_items di ON i.id = di.inventory_id
			JOIN dispatches d ON di.dispatch_id = d.id
			WHERE ((d.type = 'allocation' AND d.office_id = $1) OR d.type = 'return')
			AND d.created_at < $2 AND i.id = $3)
			  SELECT id, type, office_id, created_at
			  FROM RankedDispatches
			  WHERE rn <= 1 and type = 'allocation';`
		for _, item := range items {
			rows4, err := upper.SQL().Query(query4, *officeID, *date, item.ID)
			if err != nil {
				return nil, err
			}
			defer rows4.Close()

			var itemID int
			for rows4.Next() {
				err = rows4.Scan(&itemID)
				if err != nil {
					return nil, err
				}
			}
			if itemID != 0 {
				currentItems = append(currentItems, item)
			}
		}
		items = currentItems
	}

	query5 := `SELECT i.id, i.title, i.inventory_number, a.gross_price_difference,
		 a.estimated_duration, a.date_of_assessment, i.date_of_purchase
		FROM items i
		JOIN assessments a ON i.id = a.inventory_id
		WHERE (i.id, a.date_of_assessment, a.id) IN (
		  SELECT i.id,  MAX(a.date_of_assessment) AS max_date, MAX(a.id) AS max_id
		  FROM items i
		  JOIN assessments a ON i.id = a.inventory_id
		  WHERE a.date_of_assessment < $1 and i.id = $2
		  GROUP BY i.id)
		  LIMIT 1;`

	for _, item := range items {
		rows5, err := upper.SQL().Query(query5, *date, item.ID)
		if err != nil {
			return nil, err
		}
		defer rows5.Close()

		for rows5.Next() {
			var estimatedDuration int
			var dateOfAssessment string
			err = rows5.Scan(&item.ID, &item.Title, &item.InventoryNumber, &item.ProcurementPrice,
				&estimatedDuration, &dateOfAssessment, &item.DateOfPurchase)
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
			months := float32(sub.Hours() / 24 / 30)

			item.LostValue = item.ProcurementPrice - months*(item.ProcurementPrice*monthlyDepreciationRate/100)
		}
	}

	return items, err
}
