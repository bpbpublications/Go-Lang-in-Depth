ALTER USER postgres PASSWORD 'postgres';
CREATE USER newuser with PASSWORD 'newuser' CREATEDB;
select * from users;
\du
\l
createdb crm
CREATE TABLE "customers" (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" varchar(200) NOT NULL,
    "mobile_number" varchar(100),
    "address" varchar(400) DEFAULT NULL,
    "notes" text,
    UNIQUE (name)
);
createdb crm;
createdb crm
\l
\c crm
exit
\q
\l
\c crm
CREATE TABLE "customers" (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" varchar(200) NOT NULL,
    "mobile_number" varchar(100),
    "address" varchar(400) DEFAULT NULL,
    "notes" text,
    UNIQUE (name)
);
CREATE TABLE "customer" (
    "id"  SERIAL,
    "name" varchar(200) NOT NULL,
    "mobile_number" varchar(100),
    "address" varchar(400) DEFAULT NULL,
    "notes" text,
    UNIQUE (name)
)
\dt
show tables;
\dt
CREATE TABLE "customer" (
    "id"  SERIAL,
    "name" varchar(200) NOT NULL,
    "mobile_number" varchar(100),
    "address" varchar(400) DEFAULT NULL,
    "notes" text,
    UNIQUE (name)
);
\dt
\du
alter role newuser superuser;
\du
.quit
quit
\q
\du
\l
.quit
\q
\l
\c crm
CREATE TABLE "customer" (
    "id" serial,
    "name" varchar(200) NOT NULL,
    "mobile" varchar(100),
    "address" varchar(400) DEFAULT NULL,
    "notes" text,
    UNIQUE (name)
);
\dt
create user newuser with password 'newuser';
grant all privileges on database crm to newuser;
alter role newuser superuser;
\s database_commands.sql
