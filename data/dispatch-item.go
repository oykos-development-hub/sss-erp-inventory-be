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

// GetAll gets all records from the database, using Upper
func (t *DispatchItem) GetAll(id int) ([]*DispatchItem, error) {
	collection := Upper.Collection(t.Table())
	var all []*DispatchItem

	res := collection.Find(up.Cond{"inventory_id": id})

	err := res.OrderBy("id desc").All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

func (t *DispatchItem) GetItemListOfDispatch(dispatchID int) ([]*DispatchItem, error) {
	collection := Upper.Collection(t.Table())
	var all []*DispatchItem

	res := collection.Find(up.Cond{"dispatch_id": dispatchID})
	err := res.OrderBy("id desc").All(&all)
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (t *DispatchItem) GetAllInv(status *string, dispatch *int) ([]*DispatchItem, error) {
	var all []*DispatchItem

	if status != nil {
		query := `SELECT d.id, d.inventory_id, d.dispatch_id
			FROM items i, dispatch_items d 
			WHERE i.id = d.inventory_id and i.type = $1`

		rows, err := Upper.SQL().Query(query, status)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var item DispatchItem
			err = rows.Scan(&item.ID, &item.InventoryId, &item.DispatchId)
			if err != nil {
				return nil, err
			}
			all = append(all, &item)
		}
	} else if dispatch != nil {
		var res up.Result
		collection := Upper.Collection(t.Table())
		res = collection.Find(up.Cond{"dispatch_id": &dispatch})

		err := res.OrderBy("id desc").All(&all)
		if err != nil {
			return nil, err
		}

		return all, nil
	}

	return all, nil
}

// Get gets one record from the database, by id, using Upper
func (t *DispatchItem) Get(id int) (*DispatchItem, error) {
	var one DispatchItem
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *DispatchItem) Update(m DispatchItem) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *DispatchItem) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *DispatchItem) Insert(m DispatchItem) (int, error) {
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
