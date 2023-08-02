CREATE TABLE assessments (
    id SERIAL PRIMARY KEY,
    inventory_id INTEGER NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    active BOOLEAN NOT NULL,
    depreciation_type_id INTEGER NOT NULL,
    user_profile_id INTEGER,
    gross_price_new INTEGER NOT NULL,
    gross_price_difference INTEGER NOT NULL,
    date_of_assessment DATE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    file_id INTEGER
);
