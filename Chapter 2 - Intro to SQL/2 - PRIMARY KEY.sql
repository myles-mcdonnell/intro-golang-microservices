======= LESSON 2 =======


DROP TABLE STUDENTS;
DROP SEQUENCE STUDENTIDSEQUENCE;



==========================



CREATE SEQUENCE STUDENTIDSEQUENCE START 100;



==========================




CREATE TABLE STUDENTS (
	ID INTEGER DEFAULT nextval('STUDENTIDSEQUENCE') PRIMARY KEY,
	FIRSTNAME VARCHAR(50),
	LASTNAME VARCHAR(100),
	CITY VARCHAR(50)
);





==========================




SELECT * FROM STUDENTS;




==========================



INSERT INTO STUDENTS (
	FIRSTNAME,
	LASTNAME,
	CITY
)
VALUES (
	'Cristina',
	'Fernandez',
	'Barcelona'	
);



==========================




SELECT * FROM STUDENTS;



==========================




INSERT INTO STUDENTS (
	FIRSTNAME,
	LASTNAME,
	CITY
)
VALUES (
	'Cristina',
	'Fernandez',
	'Barcelona'	
);




==========================



INSERT INTO STUDENTS (
	FIRSTNAME,
	LASTNAME,
	CITY
)
VALUES (
	'Aitor',
	'Riaz',
	'Madrid'	
);



==========================




SELECT * FROM STUDENTS;


