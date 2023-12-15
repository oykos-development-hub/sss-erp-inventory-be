CREATE TABLE assessments (
    id SERIAL PRIMARY KEY,
    inventory_id INTEGER NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    active BOOLEAN NOT NULL,
    estimated_duration INTEGER,
    depreciation_type_id INTEGER NOT NULL,
    user_profile_id INTEGER,
    gross_price_new FLOAT NOT NULL,
    gross_price_difference FLOAT NOT NULL,
    residual_price FLOAT,
    date_of_assessment DATE,
    type_value TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    file_id INTEGER
);
