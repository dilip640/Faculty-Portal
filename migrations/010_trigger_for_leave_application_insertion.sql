-- +migrate Up

-- +migrate StatementBegin

CREATE FUNCTION check_leave_req()
RETURNS TRIGGER AS $$
DECLARE 
    leave_rec RECORD;
BEGIN
    SELECT * INTO leave_rec FROM leave_application WHERE emp_id=NEW.emp_id 
        AND (status='INITIALIZED' OR status='PENDING');
    IF FOUND THEN
        RAISE EXCEPTION 'You already have one leave in Progress!';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_leave_req
    BEFORE INSERT 
    ON leave_application
    FOR each ROW 
    EXECUTE PROCEDURE check_leave_req();
-- +migrate StatementEnd

-- +migrate Down
DROP TRIGGER trigger_leave_req ON leave_application;
DROP FUNCTION check_leave_req;