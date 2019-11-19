-- +migrate Up

-- +migrate StatementBegin
---- Store deleted faculty
CREATE FUNCTION log_faculty_history()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO faculty_history VALUES (OLD.emp_id, OLD.dept_id, OLD.start_date, CURRENT_DATE);
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_faculty_delete
    BEFORE DELETE ON faculty
    FOR each ROW EXECUTE PROCEDURE log_faculty_history();

---- Store deleted cc_faculty
CREATE FUNCTION log_cc_faculty_history()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO cc_faculty_history VALUES (OLD.emp_id, OLD.post_id, OLD.start_date, CURRENT_DATE);
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_cc_faculty_delete
    BEFORE DELETE ON cc_faculty
    FOR each ROW EXECUTE PROCEDURE log_cc_faculty_history();

-- +migrate StatementEnd

-- +migrate Down

DROP TRIGGER trigger_faculty_delete ON faculty;
DROP FUNCTION log_faculty_history;
DROP TRIGGER trigger_cc_faculty_delete ON cc_faculty;
DROP FUNCTION log_cc_faculty_history;
