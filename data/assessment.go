package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/inventory-api/contextutil"
)

// Assessment struct
type Assessment struct {
	ID                   int        `db:"id,omitempty"`
	InventoryID          int        `db:"inventory_id"`
	Active               bool       `db:"active"`
	DepreciationTypeID   int        `db:"depreciation_type_id"`
	UserProfileID        *int       `db:"user_profile_id"`
	EstimatedDuration    int        `db:"estimated_duration"`
	GrossPriceNew        float32    `db:"gross_price_new"`
	GrossPriceDifference float32    `db:"gross_price_difference"`
	ResidualPrice        *float32   `db:"residual_price"`
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

// GetAll gets all records from the database, using Upper
func (t *Assessment) GetAll(page *int, size *int, condition *up.Cond) ([]*Assessment, *uint64, error) {
	collection := Upper.Collection(t.Table())
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

	err = res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, nil
}

// Get gets one record from the database, by id, using Upper
func (t *Assessment) Get(id int) (*Assessment, error) {
	var one Assessment
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Assessment) Update(ctx context.Context, m Assessment) error {
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(m.ID)
		if err := res.Update(&m); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *Assessment) Delete(ctx context.Context, id int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *Assessment) Insert(ctx context.Context, m Assessment) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}

	var id int

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(m); err != nil {
			return err
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
