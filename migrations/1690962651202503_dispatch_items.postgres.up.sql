CREATE TABLE IF NOT EXISTS dispatch_items (
    id serial PRIMARY KEY,
    inventory_id INTEGER REFERENCES items(id) ON DELETE CASCADE,
    dispatch_id INTEGER REFERENCES dispatches(id) ON DELETE CASCADE
);
