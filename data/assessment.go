package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Assessment struct
type Assessment struct {
	ID                   int        `db:"id,omitempty"`
	InventoryID          int        `db:"inventory_id"`
	Active               bool       `db:"active"`
	DepreciationTypeID   int        `db:"depreciation_type_id"`
	UserProfileID        *int       `db:"user_profile_id"`
	GrossPriceNew        int        `db:"gross_price_new"`
	GrossPriceDifference int        `db:"gross_price_difference"`
	DateOfAssessment     *time.Time `db:"date_of_assessment"`
	CreatedAt            time.Time  `db:"created_at,omitempty"`
	UpdatedAt            time.Time  `db:"updated_at"`
	Type                 string     `db:"type_value"`
	FileID               *int       `db:"file_id"`
}

// Table returns the table name
func (t *Assessment) Table() string {
	return "assessments"
}

// GetAll gets all records from the database, using upper
func (t *Assessment) GetAll(page *int, size *int, condition *up.Cond) ([]*Assessment, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*Assessment
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
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

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, nil
}

// Get gets one record from the database, by id, using upper
func (t *Assessment) Get(id int) (*Assessment, error) {
	var one Assessment
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Assessment) Update(m Assessment) error {
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
func (t *Assessment) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Assessment) Insert(m Assessment) (int, error) {
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
