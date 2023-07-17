-- -------------------------------------------------------------
-- TablePlus 5.3.8(500)
--
-- https://tableplus.com/
--
-- Database: postgres
-- Generation Time: 2023-07-16 20:07:07.4200
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."breed_groups";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS breed_groups_id_seq;

-- Table Definition
CREATE TABLE "public"."breed_groups" (
    "id" int4 NOT NULL DEFAULT nextval('breed_groups_id_seq'::regclass),
    "name" varchar,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."breeds";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS breeds_id_seq;

-- Table Definition
CREATE TABLE "public"."breeds" (
    "id" int4 NOT NULL DEFAULT nextval('breeds_id_seq'::regclass),
    "name" varchar,
    "breed_group_id" int4,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."cats";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS cats_id_seq;

-- Table Definition
CREATE TABLE "public"."cats" (
    "id" int4 NOT NULL DEFAULT nextval('cats_id_seq'::regclass),
    "name" varchar,
    "date_of_birth" date,
    "breed_id" int4,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."cats_shows";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS cats_shows_id_seq;

-- Table Definition
CREATE TABLE "public"."cats_shows" (
    "id" int4 NOT NULL DEFAULT nextval('cats_shows_id_seq'::regclass),
    "cat_id" int4 NOT NULL,
    "show_id" int4 NOT NULL,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."organizations";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS organizations_id_seq;

-- Table Definition
CREATE TABLE "public"."organizations" (
    "id" int4 NOT NULL DEFAULT nextval('organizations_id_seq'::regclass),
    "name" varchar,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."shows";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS shows_id_seq;

-- Table Definition
CREATE TABLE "public"."shows" (
    "id" int4 NOT NULL DEFAULT nextval('shows_id_seq'::regclass),
    "name" varchar,
    "date" date,
    "organization_id" int4,
    PRIMARY KEY ("id")
);

-- Sample Data

ALTER TABLE "public"."breeds" ADD FOREIGN KEY ("breed_group_id") REFERENCES "public"."breed_groups"("id");
ALTER TABLE "public"."cats" ADD FOREIGN KEY ("breed_id") REFERENCES "public"."breeds"("id");
ALTER TABLE "public"."cats_shows" ADD FOREIGN KEY ("show_id") REFERENCES "public"."shows"("id") ON DELETE CASCADE;
ALTER TABLE "public"."cats_shows" ADD FOREIGN KEY ("cat_id") REFERENCES "public"."cats"("id") ON DELETE CASCADE;
ALTER TABLE "public"."shows" ADD FOREIGN KEY ("organization_id") REFERENCES "public"."organizations"("id") ON DELETE SET NULL;

INSERT INTO "public"."breed_groups" ("name", "id") VALUES ('Long-hair', 1);
INSERT INTO "public"."breed_groups" ("name", "id") VALUES ('Semi-long-hair', 2);
INSERT INTO "public"."breed_groups" ("name", "id") VALUES ('Foreign', 3);

INSERT INTO "public"."breeds" ("name", "breed_group_id", "id") VALUES ('Persian', 1, 1);
INSERT INTO "public"."breeds" ("name", "breed_group_id", "id") VALUES ('Himalayan', 1, 2);
INSERT INTO "public"."breeds" ("name", "breed_group_id", "id") VALUES ('Maine Coon', 2, 3);
INSERT INTO "public"."breeds" ("name", "breed_group_id", "id") VALUES ('Turkish Angora', 3, 4);

INSERT INTO "public"."cats" ("name", "date_of_birth", "breed_id", "id") VALUES ('Anna', '2021-01-01', 1, 1);
INSERT INTO "public"."cats" ("name", "date_of_birth", "breed_id", "id") VALUES ('Baxter', '2022-02-01', 1, 2);
INSERT INTO "public"."cats" ("name", "date_of_birth", "breed_id", "id") VALUES ('Charlie', '2020-03-01', 2, 3);
INSERT INTO "public"."cats" ("name", "date_of_birth", "breed_id", "id") VALUES ('Dot', '2019-04-01', 2, 4);
INSERT INTO "public"."cats" ("name", "date_of_birth", "breed_id", "id") VALUES ('Ellen', '2020-05-01', 3, 5);
INSERT INTO "public"."cats" ("name", "date_of_birth", "breed_id", "id") VALUES ('Freckles', '2022-06-01', 3, 6);
INSERT INTO "public"."cats" ("name", "date_of_birth", "breed_id", "id") VALUES ('George', '2021-07-01', 4, 7);
INSERT INTO "public"."cats" ("name", "date_of_birth", "breed_id", "id") VALUES ('Sir Henry', '2020-08-01', 4, 8);

INSERT INTO "public"."organizations" ("name", "id") VALUES ('Cat Fanciers'' Association', 1);
INSERT INTO "public"."organizations" ("name", "id") VALUES ('The International Cat Association', 2);

INSERT INTO "public"."shows" ("name", "date", "organization_id", "id") VALUES ('Cat Show 1', '2023-01-01', 1, 1);
INSERT INTO "public"."shows" ("name", "date", "organization_id", "id") VALUES ('Cat Show 2', '2023-02-01', 1, 2);
INSERT INTO "public"."shows" ("name", "date", "organization_id", "id") VALUES ('Cat Show 3', '2023-03-01', 1, 3);
INSERT INTO "public"."shows" ("name", "date", "organization_id", "id") VALUES ('Cat Show 4', '2023-04-01', 2, 4);
INSERT INTO "public"."shows" ("name", "date", "organization_id", "id") VALUES ('Cat Show 5', '2023-05-01', 2, 5);
INSERT INTO "public"."shows" ("name", "date", "organization_id", "id") VALUES ('Cat Show 6', '2023-06-01', 2, 6);

INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (1, 1, 1);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (1, 2, 2);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (1, 3, 3);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (2, 1, 4);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (2, 2, 5);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (2, 3, 6);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (3, 4, 7);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (3, 5, 8);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (3, 6, 9);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (4, 4, 10);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (4, 5, 11);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (4, 6, 12);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (5, 1, 13);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (5, 2, 14);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (5, 3, 15);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (6, 1, 16);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (6, 2, 17);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (6, 3, 18);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (7, 4, 19);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (7, 5, 20);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (7, 6, 21);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (8, 4, 22);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (8, 5, 23);
INSERT INTO "public"."cats_shows" ("cat_id", "show_id", "id") VALUES (8, 6, 24);