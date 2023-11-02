CREATE TABLE IF NOT EXISTS items (
   id SERIAL PRIMARY KEY NOT NULL,
    article_id INTEGER,
    type TEXT  NOT NULL,
    class_type_id INTEGER  NOT NULL,
    depreciation_type_id INTEGER  NOT NULL,
    supplier_id INTEGER  NOT NULL,
    serial_number TEXT,
    inventory_number TEXT  NOT NULL,
    title TEXT  NOT NULL,
    abbreviation TEXT,
    internal_ownership BOOLEAN  NOT NULL,
    office_id INTEGER,
    location TEXT,
    organization_unit_id INTEGER,
    target_organization_unit_id INTEGER,
    target_user_profile_id INTEGER,
    unit TEXT,
    amount INTEGER  NOT NULL,
    net_price INTEGER,
    gross_price INTEGER  NOT NULL,
    description TEXT,
    date_of_purchase DATE,
    inactive DATE,
    source TEXT,
    source_type TEXT,
    donor_title TEXT,
    invoice_number TEXT,
    active BOOLEAN,
    deactivation_description TEXT,
    date_of_assessment DATE,
    price_of_assessment INTEGER,
    lifetime_of_assessment_in_months INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    invoice_file_id INTEGER,
    file_id INTEGER
)
