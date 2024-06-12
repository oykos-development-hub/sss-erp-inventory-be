package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/inventory-api/contextutil"
)

// Dispatch struct
type Dispatch struct {
	ID                       int        `db:"id,omitempty"`
	DispatchID               int        `db:"dispatch_id"`
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

// GetAll gets all records from the database, using Upper
func (t *Dispatch) GetAll(page *int, size *int, condition *up.AndExpr) ([]*Dispatch, *uint64, error) {
	collection := Upper.Collection(t.Table())
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

// Get gets one record from the database, by id, using Upper
func (t *Dispatch) Get(id int) (*Dispatch, error) {
	var one Dispatch
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Dispatch) Update(ctx context.Context, m Dispatch) error {
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
func (t *Dispatch) Delete(ctx context.Context, id int) error {
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
func (t *Dispatch) Insert(ctx context.Context, m Dispatch) (int, error) {
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
