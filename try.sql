
CREATE Function dept_count(dept_name VARCHAR(20))
returns integer
READS SQL DATA
DETERMINISTIC
begin
declare d_count integer;
    SELECT count(*) into d_count
    FROM instructor
    where instructor.dept_name=dept_name;
return d_count;
end;