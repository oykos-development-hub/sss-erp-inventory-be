CREATE TABLE IF NOT EXISTS dispatches (
    id serial PRIMARY KEY,
    dispatch_id INTEGER,
    type TEXT NOT NULL,
    inventory_type TEXT NOT NULL,
    source_user_profile_id INTEGER NOT NULL,
    source_organization_unit_id INTEGER NOT NULL,
    target_user_profile_id INTEGER,
    target_organization_unit_id INTEGER NOT NULL,
    is_accepted BOOLEAN NOT NULL DEFAULT FALSE,
    serial_number TEXT,
    office_id INTEGER,
    date TIMESTAMP,
    dispatch_description TEXT,
    file_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE OR REPLACE FUNCTION increment_dispatch_id_for_allocation() RETURNS TRIGGER AS $$
DECLARE
    max_dispatch_id INT;
BEGIN
    SELECT dispatch_id
    INTO max_dispatch_id
    FROM dispatches
    WHERE type = 'allocation' AND source_organization_unit_id = NEW.source_organization_unit_id
    ORDER BY dispatch_id DESC
    LIMIT 1;
    
    IF max_dispatch_id IS NOT NULL THEN
        NEW.dispatch_id := max_dispatch_id + 1;
    ELSE
        NEW.dispatch_id := 1;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION increment_dispatch_id_for_revers() RETURNS TRIGGER AS $$
DECLARE
    max_dispatch_id INT;
BEGIN
    SELECT dispatch_id
    INTO max_dispatch_id
    FROM dispatches
    WHERE type = 'revers' AND source_organization_unit_id = NEW.source_organization_unit_id
    ORDER BY dispatch_id DESC
    LIMIT 1;
    
    IF max_dispatch_id IS NOT NULL THEN
        NEW.dispatch_id := max_dispatch_id + 1;
    ELSE
        NEW.dispatch_id := 1;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_increment_dispatch_id_for_allocation
BEFORE INSERT ON dispatches
FOR EACH ROW
WHEN (NEW.type = 'allocation')
EXECUTE FUNCTION increment_dispatch_id_for_allocation();

CREATE TRIGGER trigger_increment_dispatch_id_for_revers
BEFORE INSERT ON dispatches
FOR EACH ROW
WHEN (NEW.type = 'revers')
EXECUTE FUNCTION increment_dispatch_id_for_revers();

