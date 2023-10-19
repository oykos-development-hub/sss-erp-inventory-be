package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Dispatch struct
type Dispatch struct {
	ID                       int        `db:"id,omitempty"`
	Type                     string     `db:"type"`
	InventoryType            string     `db:"inventory_type"`
	SourceUserProfileID      int        `db:"source_user_profile_id"`
	TargetUserProfileID      *int       `db:"target_user_profile_id"`
	SourceOrganizationUnitID int        `db:"source_organization_unit_id"`
	TargetOrganizationUnitID int        `db:"target_organization_unit_id"`
	IsAccepted               bool       `db:"is_accepted"`
	SerialNumber             *string    `db:"serial_number"`
	OfficeID                 *int       `db:"office_id"`
	Date                     *time.Time `db:"date"`
	DispatchDescription      *string    `db:"dispatch_description"`
	FileID                   *int       `db:"file_id"`
	CreatedAt                time.Time  `db:"created_at,omitempty"`
	UpdatedAt                time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *Dispatch) Table() string {
	return "dispatches"
}

// GetAll gets all records from the database, using upper
func (t *Dispatch) GetAll(page *int, size *int, condition *up.AndExpr) ([]*Dispatch, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*Dispatch
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
func (t *Dispatch) Get(id int) (*Dispatch, error) {
	var one Dispatch
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Dispatch) Update(m Dispatch) error {
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
func (t *Dispatch) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Dispatch) Insert(m Dispatch) (int, error) {
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
