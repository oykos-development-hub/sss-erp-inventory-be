package data

import (
	up "github.com/upper/db/v4"
)

// DispatchItem struct
type DispatchItem struct {
	ID          int `db:"id,omitempty"`
	InventoryId int `db:"inventory_id"`
	DispatchId  int `db:"dispatch_id"`
}

// Table returns the table name
func (t *DispatchItem) Table() string {
	return "dispatch_items"
}

// GetAll gets all records from the database, using upper
func (t *DispatchItem) GetAll(id int) ([]*DispatchItem, error) {
	collection := upper.Collection(t.Table())
	var all []*DispatchItem

	res := collection.Find(up.Cond{"inventory_id": id})

	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *DispatchItem) Get(id int) (*DispatchItem, error) {
	var one DispatchItem
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *DispatchItem) Update(m DispatchItem) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *DispatchItem) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *DispatchItem) Insert(m DispatchItem) (int, error) {
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
