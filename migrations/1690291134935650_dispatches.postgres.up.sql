CREATE TABLE IF NOT EXISTS dispatches (
    id serial PRIMARY KEY,
    type TEXT NOT NULL,
    source_user_profile_id INTEGER NOT NULL,
    source_organization_unit_id INTEGER NOT NULL,
    target_user_profile_id INTEGER,
    target_organization_unit_id INTEGER NOT NULL,
    is_accepted BOOLEAN NOT NULL DEFAULT FALSE,
    serial_number TEXT,
    office_id INTEGER,
    dispatch_description TEXT,
    file_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
