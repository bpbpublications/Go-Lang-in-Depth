CREATE TABLE "products" (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" varchar(200) NOT NULL,
    "category" varchar(100),
    "image" varchar(400) DEFAULT NULL,
    "notes" text,
    UNIQUE (name)
);