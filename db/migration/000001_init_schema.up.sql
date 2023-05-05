CREATE SCHEMA "autocopy";

-- statuses for runs and orders
CREATE TYPE "autocopy"."status" as ENUM (
    'initialized',
    'started',
    'pending',
    'complete',
    'failed'
);

-- statuses for files
CREATE TYPE "autocopy"."filesystem" as ENUM (
    'PFS',
    'HFS',
    'AFS',
    'PFSHFS',
    'PFSAFS',
    'HFSAFS',
    'PFSHFSAFS'
);

-- suite of analyses
DROP TABLE IF EXISTS "autocopy".analysis_suite;
create table "autocopy".analysis_suite (
    "id" bigserial PRIMARY KEY,
    "suite_name" varchar(255) NOT NULL,
    UNIQUE ("suite_name")
);

-- analysis type table
DROP TABLE IF EXISTS "autocopy".analysis_type;
create table "autocopy".analysis_type (
    "id" bigserial PRIMARY KEY,
    "anls_type_name" varchar NOT NULL,
    "anls_suite_id" int NOT NULL,
    FOREIGN KEY ("anls_suite_id") REFERENCES "autocopy".analysis_suite("id") ON DELETE CASCADE
);

-- user
DROP TABLE IF EXISTS "autocopy".user;
create table "autocopy".user (
    "id" bigserial PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "username" varchar(255) NOT NULL,
    "email" varchar(255) NOT NULL,
    UNIQUE ("username"),
    UNIQUE ("email")
);

-- sample table
DROP TABLE IF EXISTS "autocopy".sample;
create table "autocopy".sample (
    "id" bigserial PRIMARY KEY,
    "samplename" VARCHAR(255) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    UNIQUE ("samplename")
);

-- orders table
DROP TABLE IF EXISTS "autocopy".order;
create table "autocopy".order (
    "id" bigserial PRIMARY KEY,
    "user_id" int NOT NULL,
    "anls_suite_id" int NOT NULL,
    "status" autocopy.status NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY ("anls_suite_id") REFERENCES "autocopy".analysis_suite("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "autocopy".user("id") ON DELETE CASCADE
);

-- order sample many to many relation
DROP TABLE IF EXISTS "autocopy".order_sample;
create table "autocopy".order_sample (
    "id" bigserial PRIMARY KEY,
    "order_id" int NOT NULL,
    "sample_id" int NOT NULL,
    FOREIGN KEY ("order_id") REFERENCES "autocopy".order("id") ON DELETE CASCADE,
    FOREIGN KEY ("sample_id") REFERENCES "autocopy".sample("id") ON DELETE CASCADE
);

-- analysis runs
DROP TABLE IF EXISTS "autocopy".analysis_run;
create table "autocopy".analysis_run (
    "id" bigserial PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "anls_id" int NOT NULL,
    "order_sample_id" int NOT NULL,
    "status" autocopy.status NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY ("anls_id") REFERENCES "autocopy".analysis_type("id") ON DELETE CASCADE,
    FOREIGN KEY ("order_sample_id") REFERENCES "autocopy".order_sample("id") ON DELETE CASCADE
);

-- analysis run output files
DROP TABLE IF EXISTS "autocopy".analysis_run_outputs;
create table "autocopy".analysis_run_outputs (
    "id" bigserial PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "anls_run_id" int NOT NULL,
    "filesystem_path" VARCHAR(255) NOT NULL,
    "location_encoding" autocopy.filesystem NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY ("anls_run_id") REFERENCES "autocopy".analysis_run("id") ON DELETE CASCADE
)

