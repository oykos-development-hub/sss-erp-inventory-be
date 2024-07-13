CREATE TABLE IF NOT EXISTS logs (
    id serial PRIMARY KEY,
    operation VARCHAR(10),
    entity TEXT,
    old_state JSONB,
    new_state JSONB,
    user_id INTEGER,
    item_id INTEGER,
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION log_changes() RETURNS TRIGGER AS $$
DECLARE
    user_id INTEGER;
BEGIN
    BEGIN
        SELECT current_setting('myapp.user_id')::INTEGER INTO user_id;
    EXCEPTION
        WHEN others THEN
            user_id := 0;  
    END;

    IF TG_OP = 'INSERT' THEN
        INSERT INTO logs (operation, new_state, user_id, item_id, entity)
        VALUES ('INSERT', row_to_json(NEW)::jsonb, user_id, NEW.id, TG_TABLE_NAME);
        RETURN NEW;
    ELSIF TG_OP = 'UPDATE' THEN
        INSERT INTO logs (operation, old_state, new_state, user_id, item_id, entity)
        VALUES ('UPDATE', row_to_json(OLD)::jsonb, row_to_json(NEW)::jsonb, user_id, NEW.id, TG_TABLE_NAME);
        RETURN NEW;
    ELSIF TG_OP = 'DELETE' THEN
        INSERT INTO logs (operation, old_state, user_id, item_id, entity)
        VALUES ('DELETE', row_to_json(OLD)::jsonb, user_id, OLD.id, TG_TABLE_NAME);
        RETURN OLD;
    END IF;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER dispatches_insert
AFTER INSERT ON dispatches
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER dispatches_update
AFTER UPDATE ON dispatches
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER dispatches_delete
AFTER DELETE ON dispatches
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER items_insert
AFTER INSERT ON items
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER items_update
AFTER UPDATE ON items
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER items_delete
AFTER DELETE ON items
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER assessments_insert
AFTER INSERT ON assessments
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER assessments_update
AFTER UPDATE ON assessments
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER assessments_delete
AFTER DELETE ON assessments
FOR EACH ROW
EXECUTE FUNCTION log_changes();