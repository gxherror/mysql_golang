#EXAMPLE 1
delimiter $$
CREATE TRIGGER upd_check BEFORE UPDATE ON account
FOR EACH ROW
BEGIN
IF NEW.amount < 0 THEN
SET NEW.amount = 0;
ELSEIF NEW.amount > 100 THEN
SET NEW.amount = 100;
END IF;
END$$
delimiter ;

#EXAMPLE 2
DELIMITER $$

CREATE TRIGGER after_members_insert
AFTER INSERT
ON members FOR EACH ROW
BEGIN
    IF NEW.birthDate IS NULL THEN
        INSERT INTO reminders(memberId, message)
        VALUES(new.id,CONCAT('Hi ', NEW.name, ', please update your date of birth.'));
    END IF;
END$$

DELIMITER ;


#5_7
DELIMITER $$
CREATE Trigger student_delete before delete ON student
for each row
BEGIN
    IF OLD.name in (SELECT `name` FROM student_takes )
    THEN SIGNAL 
    SQLSTATE '45000'
    SET MESSAGE_TEXT = 'An error occurred';
    END IF;
END$$

DELIMITER ;



