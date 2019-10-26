-- +migrate Up
-- +migrate StatementBegin
CREATE FUNCTION insert_update_faculty(uname VARCHAR, 
        startDate DATE, endDate DATE, deptn VARCHAR)
RETURNS text AS $$
DECLARE
        result faculty%ROWTYPE;
BEGIN
        SELECT * INTO result FROM faculty WHERE  emp_id = uname;
        IF NOT FOUND
        THEN
                INSERT INTO faculty (emp_id, start_date, end_date, dept)
                                        VALUES (uname, startDate, endDate, deptn);
                RETURN 'new';
        ELSE    
                IF startDate > result.end_date AND endDate > startDate
                THEN
                        INSERT INTO faculty (emp_id, start_date, end_date, dept)
                                        VALUES (uname, startDate, endDate, deptn);
                        RETURN 'new';
                ELSIF startDate = result.start_date AND endDate > startDate
                THEN
                        UPDATE faculty SET end_date = endDate, dept = deptn
                                                WHERE emp_id = uname AND start_date = startDate;
                        RETURN 'update';
                ELSE
                        RAISE EXCEPTION 'Faculty Exists!';   
                END IF;             
        END IF;
END;
$$ LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate Down
DROP FUNCTION insert_update_faculty(VARCHAR, DATE, DATE, VARCHAR);