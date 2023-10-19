package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// RealEstate struct
type RealEstate struct {
	ID                       int       `db:"id,omitempty"`
	ItemID                   int       `db:"item_id"`
	TypeID                   string    `db:"type_id"`
	Title                    string    `db:"title"`
	SquareArea               int       `db:"square_area"`
	LandSerialNumber         string    `db:"land_serial_number"`
	EstateSerialNumber       string    `db:"estate_serial_number"`
	OwnershipType            string    `db:"ownership_type"`
	OwnershipScope           string    `db:"ownership_scope"`
	OwnershipInvestmentScope string    `db:"ownership_investment_scope"`
	LimitationsDescription   string    `db:"limitations_description"`
	LimitationID             string    `db:"limitation_id"`
	PropertyDocument         string    `db:"property_document"`
	Document                 string    `db:"document"`
	FileID                   int       `db:"file_id"`
	CreatedAt                time.Time `db:"created_at,omitempty"`
	UpdatedAt                time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *RealEstate) Table() string {
	return "real_estates"
}

// GetAll gets all records from the database, using upper
func (t *RealEstate) GetAll(page *int, size *int, condition *up.Cond) ([]*RealEstate, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*RealEstate
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

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, nil
}

// Get gets one record from the database, by id, using upper
func (t *RealEstate) Get(id int) (*RealEstate, error) {
	var one RealEstate
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *RealEstate) Update(m RealEstate) error {
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
func (t *RealEstate) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *RealEstate) Insert(m RealEstate) (int, error) {
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
