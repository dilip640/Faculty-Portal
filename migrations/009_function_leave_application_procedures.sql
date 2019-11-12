-- +migrate Up

-- +migrate StatementBegin

CREATE FUNCTION get_remaining_leave(empID text, y smallint)
RETURNS int AS $$
DECLARE 
    leave_num int;
BEGIN
    SELECT no_of_leaves INTO leave_num FROM leave WHERE emp_id=empID AND year=y;
    IF NOT FOUND THEN
        INSERT INTO leave VALUES(empID, 10, y);
    ELSE
        RETURN leave_num;
    END IF;
    RETURN 10;
END;
$$ LANGUAGE plpgsql;

-- +migrate StatementEnd

-- +migrate Down
DROP FUNCTION get_remaining_leave;