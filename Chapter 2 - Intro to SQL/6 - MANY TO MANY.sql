===== LESSON 7 : MANY-2-MANY ======

DROP TABLE STUDENTS;
DROP TABLE COURSES;
DROP TABLE STUDENT_COURSES;

==========================


CREATE TABLE COURSES (
	ID INT PRIMARY KEY,
	NAME VARCHAR(50)
);


CREATE TABLE STUDENTS (
	ID INT PRIMARY KEY,
	FIRSTNAME VARCHAR(50) NOT NULL,
	LASTNAME VARCHAR(100) NOT NULL
);

CREATE TABLE STUDENT_COURSES (
	STUDENTID INT REFERENCES STUDENTS(ID),
	COURSEID INT REFERENCES COURSES(ID),
	PRIMARY KEY (STUDENTID, COURSEID)
);


==========================


INSERT INTO COURSES (
	ID, 
	NAME
)
VALUES (
	1,
	'Golang');
	
INSERT INTO COURSES (
	ID, 
	NAME
)
VALUES (
	2,
	'SQL');


INSERT INTO STUDENTS (
	ID,
	FIRSTNAME,
	LASTNAME
)
VALUES (
	1,
	'Cristina',
	'Fernandez'
);

INSERT INTO STUDENTS (
	ID,
	FIRSTNAME,
	LASTNAME
)
VALUES (
	2,
	'Aitor',
	'Riaz'
);


INSERT INTO STUDENT_COURSES (
	STUDENTID, COURSEID 
) 
VALUES (
	1, 1);
	
INSERT INTO STUDENT_COURSES (
	STUDENTID, COURSEID 
) 
VALUES (
	1, 2);
	
INSERT INTO STUDENT_COURSES (
	STUDENTID, COURSEID 
) 
VALUES (
	2, 2);



==========================



SELECT 
	S.FIRSTNAME, S.LASTNAME, C.NAME
FROM
	STUDENTS S
JOIN
	STUDENT_COURSES SC
ON 
	S.ID = SC.STUDENTID
JOIN
	COURSES C
ON
	SC.COURSEID = C.ID;

===== LESSON 7 : MANY-2-MANY ======

DROP TABLE STUDENTS;
DROP TABLE COURSES;
DROP TABLE STUDENT_COURSES;

==========================


CREATE TABLE COURSES (
	ID INT PRIMARY KEY,
	NAME VARCHAR(50)
);


CREATE TABLE STUDENTS (
	ID INT PRIMARY KEY,
	FIRSTNAME VARCHAR(50) NOT NULL,
	LASTNAME VARCHAR(100) NOT NULL
);

CREATE TABLE STUDENT_COURSES (
	STUDENTID INT REFERENCES STUDENTS(ID),
	COURSEID INT REFERENCES COURSES(ID),
	PRIMARY KEY (STUDENTID, COURSEID)
);


==========================


INSERT INTO COURSES (
	ID,
	NAME
)
VALUES (
	1,
	'Golang');

INSERT INTO COURSES (
	ID,
	NAME
)
VALUES (
	2,
	'SQL');


INSERT INTO STUDENTS (
	ID,
	FIRSTNAME,
	LASTNAME
)
VALUES (
	1,
	'Cristina',
	'Fernandez'
);

INSERT INTO STUDENTS (
	ID,
	FIRSTNAME,
	LASTNAME
)
VALUES (
	2,
	'Aitor',
	'Riaz'
);


INSERT INTO STUDENT_COURSES (
	STUDENTID, COURSEID
)
VALUES (
	1, 1);

INSERT INTO STUDENT_COURSES (
	STUDENTID, COURSEID
)
VALUES (
	1, 2);

INSERT INTO STUDENT_COURSES (
	STUDENTID, COURSEID
)
VALUES (
	2, 2);



==========================



SELECT
	S.FIRSTNAME, S.LASTNAME, C.NAME
FROM
	STUDENTS S
JOIN
	STUDENT_COURSES SC
ON
	S.ID = SC.STUDENTID
JOIN
	COURSES C
ON
	SC.COURSEID = C.ID;

