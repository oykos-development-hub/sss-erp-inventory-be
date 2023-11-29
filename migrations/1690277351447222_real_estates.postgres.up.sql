CREATE TABLE IF NOT EXISTS real_estates (
    id SERIAL PRIMARY KEY,
    item_id INTEGER REFERENCES items(id) ON DELETE CASCADE,
    title TEXT,
    type_id TEXT NOT NULL,
    square_area FLOAT NOT NULL,
    land_serial_number TEXT NOT NULL,
    estate_serial_number TEXT,
    ownership_type TEXT ,
    ownership_scope TEXT,
    ownership_investment_scope TEXT NOT NULL,
    limitations_description TEXT,
    limitation_id BOOLEAN,
    property_document TEXT,
    document TEXT,
    file_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
