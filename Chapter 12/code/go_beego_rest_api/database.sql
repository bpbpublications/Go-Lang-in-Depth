
CREATE TABLE "customer" (
    "id" serial,
    "name" varchar(200) NOT NULL,
    "mobile" varchar(100),
    "address" varchar(400) DEFAULT NULL,
    "notes" text,
    UNIQUE (name)
);


create user newuser with password 'newuser';
	
	
grant all privileges on database crm to newuser;