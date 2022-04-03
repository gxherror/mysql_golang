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

#5_15
DELIMITER $$
CREATE Function avg_salary(comp_name varchar(20))
returns integer
READS SQL DATA
DETERMINISTIC
begin
declare avg_salary integer;
    SELECT avg(salary) INTO avg_salary
    FROM works  
    WHERE company_name=comp_name
    GROUP BY company_name;
RETURN avg_salary;
end$$

DELIMITER ;