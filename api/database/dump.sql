CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "is_personal" boolean DEFAULT false,
  "photo" varchar,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp,
  "student_code" int
);

CREATE TABLE "students" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar NOT NULL,
  "weight" varchar NOT NULL,
  "height" float NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp,
  "training_sheets_code" int
);

CREATE TABLE "training_sheets" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "disabled" boolean,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp,
  "workouts_code" int
);

CREATE TABLE "workouts" (
  "id" SERIAL PRIMARY KEY,
  "exercise" varchar NOT NULL,
  "equipment" int,
  "description" varchar,
  "repetitions" varchar,
  "series" int,
  "rest_time" int,
  "weight" float,
  "cadence" varchar,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "equipment" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar
);

ALTER TABLE "users" ADD FOREIGN KEY ("student_code") REFERENCES "students" ("id");

ALTER TABLE "students" ADD FOREIGN KEY ("training_sheets_code") REFERENCES "training_sheets" ("id");

ALTER TABLE "training_sheets" ADD FOREIGN KEY ("workouts_code") REFERENCES "workouts" ("id");

ALTER TABLE "equipment" ADD FOREIGN KEY ("id") REFERENCES "workouts" ("equipment");
