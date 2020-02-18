====== LESSON 6 : AGGREGATES ======


DROP TABLE FEDS;



==================




CREATE TABLE FEDS (
	adshex VARCHAR(10),
	flight_id  VARCHAR(10),
	latitude VARCHAR(20),
	longitude VARCHAR(20),
	altitude INTEGER,
	speed VARCHAR(20),
	track VARCHAR(20),
	squawk  VARCHAR(10),
	type  VARCHAR(10),
	timestamp VARCHAR(20),
	name  VARCHAR(50),
	other_names1  VARCHAR(50),
	other_names2  VARCHAR(50),
	n_number VARCHAR(10),
	serial_number VARCHAR(20),
	mfr_mdl_code VARCHAR(20),
	mfr VARCHAR(50),
	model VARCHAR(20),
	year_mfr VARCHAR(10),
	type_aircraft VARCHAR(20),
	agency VARCHAR(10));
	
	
	
	
===================




SELECT
	COUNT(*)
FROM
	FEDS;
	
	
	
	
===================




SELECT
	*
FROM
	FEDS
WHERE
	altitude = 5500;
	
	
	
	
====================



SELECT 
	DISTINCT(NAME) 
FROM 
	FEDS;
	
	
	
	
	
====================



SELECT 
	COUNT(DISTINCT(NAME))
FROM 
	FEDS;
	
	

====================



SELECT 
	COUNT(*),
	NAME 
FROM 
	FEDS
GROUP BY
	NAME
HAVING 
	NAME LIKE '%INC';