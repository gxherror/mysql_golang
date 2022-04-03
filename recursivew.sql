
WITH RECURSIVE 
rec_prereq(course_id,prereq_id) AS
(
    SELECT course_id,prereq_id
    FROM prereq 
union ALL
    SELECT rec_prereq.course_id,prereq.prereq_id
    FROM rec_prereq,prereq
    WHERE rec_prereq.prereq_id=prereq.course_id
)
SELECT * 
FROM rec_prereq
WHERE course_id="CS-347";  


