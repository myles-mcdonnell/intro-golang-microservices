===== LESSON 5 : UUID IDS ======

DROP TABLE STUDENTS;
DROP TABLE CITY;


==========================

create extension "uuid-ossp";

CREATE TABLE CITY (
                      ID UUID PRIMARY KEY,
                      NAME VARCHAR(50)
);



CREATE TABLE STUDENTS (
                          ID UUID PRIMARY KEY,
                          FIRSTNAME VARCHAR(50) NOT NULL,
                          LASTNAME VARCHAR(100) NOT NULL,
                          EMAIL VARCHAR(200) NOT NULL UNIQUE,
                          CITYID UUID REFERENCES CITY(ID)
);



==========================


INSERT INTO CITY (
    ID,
    NAME
)
VALUES (
           uuid_generate_v4(),
           'Barcelona');

SELECT * FROM CITY;

INSERT INTO STUDENTS (
    ID,
    FIRSTNAME,
    LASTNAME,
    EMAIL,
    CITYID
)
VALUES (
           uuid_generate_v4(),
           'Cristina',
           'Fernandez',
           'cris.fer@gmail.com',
           'fcf1f30c-8de7-4ba3-ba97-28ac997d84f4'
       );





==========================

SELECT * FROM STUDENTS;


INSERT INTO STUDENTS (
    ID,
    FIRSTNAME,
    LASTNAME,
    EMAIL,
    CITYID
)
VALUES (
           '4cef6475-3213-4682-8879-fb42e6aa2282',
           'Cristina',
           'Fernandez',
           'cris.fer@gmail.com',
           'fcf1f30c-8de7-4ba3-ba97-28ac997d84f4'
       )
on conflict (ID) do update
    set
        FIRSTNAME = 'Cristina',
        LASTNAME = 'Fernandez',
        EMAIL = 'cris.fer@gmail.com',
        CITYID = 'fcf1f30c-8de7-4ba3-ba97-28ac997d84f4';

DELETE FROM STUDENTS;


==========================




SELECT * FROM STUDENTS;

SELECT
    S.FIRSTNAME, S.LASTNAME, S.EMAIL, C.NAME
FROM
    STUDENTS S
        JOIN
    CITY C
    ON
            S.CITYID = C.ID;


