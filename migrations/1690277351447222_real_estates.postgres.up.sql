CREATE TABLE IF NOT EXISTS real_estates (
    id SERIAL PRIMARY KEY,
    item_id INTEGER REFERENCES items(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    type_id TEXT NOT NULL,
    square_area INTEGER NOT NULL,
    land_serial_number TEXT NOT NULL,
    estate_serial_number TEXT NOT NULL,
    ownership_type TEXT NOT NULL,
    ownership_scope TEXT NOT NULL,
    ownership_investment_scope TEXT NOT NULL,
    limitations_description TEXT NOT NULL,
    limitation_id TEXT,
    property_document TEXT,
    document TEXT,
    file_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
