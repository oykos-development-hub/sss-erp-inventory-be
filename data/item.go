package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Item struct
type Item struct {
	ID                           int        `db:"id,omitempty"`
	ArticleID                    *int       `db:"article_id"`
	Type                         string     `db:"type"`
	ClassTypeID                  int        `db:"class_type_id"`
	DepreciationTypeID           int        `db:"depreciation_type_id"`
	SupplierID                   int        `db:"supplier_id"`
	InvoiceID                    int        `db:"invoice_id"`
	DonorID                      *int       `db:"donor_id"`
	SerialNumber                 *string    `db:"serial_number"`
	InventoryNumber              string     `db:"inventory_number"`
	Title                        string     `db:"title"`
	Abbreviation                 *string    `db:"abbreviation"`
	InternalOwnership            bool       `db:"internal_ownership"`
	OfficeID                     int        `db:"office_id"`
	ContractID                   int        `db:"contract_id"`
	Location                     *string    `db:"location"`
	TargetUserProfileID          *int       `db:"target_user_profile_id"`
	OrganizationUnitID           *int       `db:"organization_unit_id"`
	TargetOrganizationUnitID     *int       `db:"target_organization_unit_id"`
	Unit                         *string    `db:"unit"`
	Amount                       int        `db:"amount"`
	NetPrice                     *float32   `db:"net_price"`
	GrossPrice                   float32    `db:"gross_price"`
	Description                  *string    `db:"description"`
	DateOfPurchase               time.Time  `db:"date_of_purchase"`
	Inactive                     *time.Time `db:"inactive"`
	Source                       *string    `db:"source"`
	SourceType                   *string    `db:"source_type"`
	DonorTitle                   *string    `db:"donor_title"`
	InvoiceNumber                *string    `db:"invoice_number"`
	Active                       bool       `db:"active"`
	DeactivationDescription      *string    `db:"deactivation_description"`
	DateOfAssessment             *time.Time `db:"date_of_assessment"`
	PriceOfAssessment            *int       `db:"price_of_assessment"`
	LifetimeOfAssessmentInMonths *int       `db:"lifetime_of_assessment_in_months"`
	CreatedAt                    time.Time  `db:"created_at,omitempty"`
	UpdatedAt                    time.Time  `db:"updated_at"`
	InvoiceFileID                *int       `db:"invoice_file_id"`
	FileID                       *int       `db:"file_id"`
	DeactivationFileID           *int       `db:"deactivation_file_id"`
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
