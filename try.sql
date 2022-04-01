ALTER TABLE course ADD PRIMARY KEY(course_id);

CREATE TABLE scores
(
    id MEDIUMINT,
    s_id MEDIUMINT,
    course_id VARCHAR(7),
    score NUMERIC(5,2),
    PRIMARY KEY(id),
    FOREIGN KEY(course_id) REFERENCES course(course_id)
);
ALTER
