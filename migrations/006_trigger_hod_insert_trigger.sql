-- +migrate Up

-- +migrate StatementBegin

CREATE FUNCTION assign_hod()
RETURNS TRIGGER AS $$
DECLARE 
    faculty_rec faculty%ROWTYPE;
    hod_rec hod%ROWTYPE;
BEGIN
    SELECT * INTO faculty_rec FROM faculty WHERE emp_id=NEW.emp_id;
    IF NOT FOUND THEN
        RAISE EXCEPTION 'No faculty found with emp_id!';
    ELSIF faculty_rec.dept_id <> NEW.dept_id THEN
        RAISE EXCEPTION 'Faculty is not of the same Department!';
    ELSE
        NEW.end_date = NEW.start_date + interval '1 year';
        SELECT * INTO hod_rec FROM hod WHERE dept_id=NEW.dept_id;
        IF FOUND THEN
            INSERT INTO hod_history VALUES (hod_rec.emp_id,
                                    hod_rec.dept_id, hod_rec.start_date, hod_rec.end_date);
            DELETE FROM hod WHERE dept_id=NEW.dept_id;
        END IF; 
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_hod 
    BEFORE INSERT ON hod
    FOR EACH ROW EXECUTE PROCEDURE assign_hod();
-- +migrate StatementEnd

-- +migrate Down
DROP TRIGGER trigger_hod ON hod;
DROP FUNCTION assign_hod;