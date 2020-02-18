# Chapter 2 - Introduction to SQL

To execute the SQL scripts first start a PostgresSQL database.  The easiest way to do this install [Docker](https://docs.docker.com/install/) then run `docker run -p 5432:5432 postgres`

Next install a PostgresSQL client tool such as [Postico](https://eggerapps.at/postico/) or [DBeaver](https://dbeaver.io/)

## Lesson 7 - Aggregates

Lesson 7 uses an imported dataset [feds1.csv](feds1.csv).  To import this data;
 
 - Execute the CREATE TABLE command to create FEDS table
 - Install the psql CLI; on OSX `brew install libpq` then add to path, typically `/usr/local/Cellar/libpq/12.1_1/bin/` but check your system as it may be a different version
 - Connect to your server (typically `psql -h localhost -U postgres`) to get to a Postgres terminal
 - Execute `\COPY feds FROM '{path to file on your system}/feds1.csv' WITH (FORMAT csv);`