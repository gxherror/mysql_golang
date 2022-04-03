DELIMITER $$
/* This is a complete statement, not part of the procedure, so use the custom delimiter $$ */
DROP PROCEDURE my_procedure$$

/* Now start the procedure code */
CREATE PROCEDURE my_procedure()
BEGIN    
  /* Inside the procedure, individual statements terminate with ; */
  CREATE TABLE tablea (
     col1 INT,
     col2 INT
  );

  INSERT INTO tablea
    SELECT * FROM table1;

  CREATE TABLE tableb (
     col1 INT,
     col2 INT
  );
  INSERT INTO tableb
    SELECT * FROM table2;
  
/* whole procedure ends with the custom delimiter */
END$$

/* Finally, reset the delimiter to the default ; */
DELIMITER ;