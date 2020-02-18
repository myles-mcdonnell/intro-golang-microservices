===== LESSON 3 ======

DROP TABLE STUDENTS;
DROP SEQUENCE STUDENTIDSEQUENCE;
DROP TABLE CITY;
DROP SEQUENCE CITYIDSEQUENCE;




==========================



CREATE SEQUENCE STUDENTIDSEQUENCE START 100;
CREATE SEQUENCE CITYIDSEQUENCE START 100;



==========================


CREATE TABLE CITY (
                      ID INTEGER DEFAULT nextval('CITYIDSEQUENCE') PRIMARY KEY,
                      NAME VARCHAR(50)
);



CREATE TABLE STUDENTS (
                          ID INTEGER DEFAULT nextval('STUDENTIDSEQUENCE') PRIMARY KEY,
                          FIRSTNAME VARCHAR(50),
                          LASTNAME VARCHAR(100),
                          CITYID INT REFERENCES CITY(ID) NULL
);



==========================


INSERT INTO CITY (
    NAME
)
VALUES (
           'Barcelona');


INSERT INTO STUDENTS (
    FIRSTNAME,
    LASTNAME,
    CITYID
)
VALUES (
           'Cristina',
           'Fernandez',
           currval('CITYIDSEQUENCE')
       );



==========================




SELECT * FROM CITY;



SELECT * FROM STUDENTS;

SELECT
    S.FIRSTNAME, S.LASTNAME, C.NAME
FROM
    STUDENTS S
        JOIN
    CITY C
    ON
            S.CITYID = C.ID;




==========================




INSERT INTO STUDENTS (
    FIRSTNAME,
    LASTNAME
)
VALUES (
           'Aitor',
           'Riaz'
       );




==========================




SELECT
    S.FIRSTNAME, S.LASTNAME, C.NAME
FROM
    STUDENTS S
        JOIN
    CITY C
    ON
            S.CITYID = C.ID;




==========================




SELECT
    S.FIRSTNAME, S.LASTNAME, C.NAME
FROM
    STUDENTS S
        LEFT OUTER JOIN
    CITY C
    ON
            S.CITYID = C.ID;



===========================


INSERT INTO STUDENTS (
    FIRSTNAME,
    LASTNAME,
    CITYID
)
VALUES (
           'Myles',
           'McDonnell',
           5
       );


DELETE FROM CITY WHERE NAME = 'Barcelona';


=========================


ALTER TABLE STUDENTS DROP CONSTRAINT students_cityid_fkey;


ALTER TABLE STUDENTS
    ADD CONSTRAINT students_cityid_fkey
        FOREIGN KEY (CITYID)
            REFERENCES CITY(ID)
            ON DELETE CASCADE;





=========================





select kcu.table_schema || '.' || kcu.table_name as foreign_table,
       '>-' as rel,
       rel_kcu.table_schema || '.' || rel_kcu.table_name as primary_table,
       kcu.ordinal_position as no,
       kcu.column_name as fk_column,
       '=' as join,
       rel_kcu.column_name as pk_column,
       kcu.constraint_name
from information_schema.table_constraints tco
         join information_schema.key_column_usage kcu
              on tco.constraint_schema = kcu.constraint_schema
                  and tco.constraint_name = kcu.constraint_name
         join information_schema.referential_constraints rco
              on tco.constraint_schema = rco.constraint_schema
                  and tco.constraint_name = rco.constraint_name
         join information_schema.key_column_usage rel_kcu
              on rco.unique_constraint_schema = rel_kcu.constraint_schema
                  and rco.unique_constraint_name = rel_kcu.constraint_name
                  and kcu.ordinal_position = rel_kcu.ordinal_position
where tco.constraint_type = 'FOREIGN KEY'
order by kcu.table_schema,
         kcu.table_name,
         kcu.ordinal_position;



