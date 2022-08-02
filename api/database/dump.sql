CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "uid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL UNIQUE,
  "password" VARCHAR NOT NULL,
  "is_personal" BOOLEAN DEFAULT false,
  "photo" VARCHAR,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
  "deleted_at" TIMESTAMP
);

CREATE TABLE "sessions" (
	"id" SERIAL PRIMARY KEY,
  "uid" uuid NOT NULL DEFAULT gen_random_uuid(),
	"user_id" INT NOT NULL,
	"refresh_token" CHARACTER varying NOT NULL UNIQUE,
	"created_at" TIMESTAMP NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMP NOT NULL DEFAULT now(),
	"deleted_at" TIMESTAMP
);

CREATE TABLE "user_student" (
  "user_id" INT,
  "student_id" INT,
  PRIMARY KEY ("user_id", "student_id")
);

CREATE TABLE "students" (
  "id" SERIAL PRIMARY KEY,
  "uid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "first_name" VARCHAR NOT NULL,
  "last_name" VARCHAR NOT NULL,
  "birth" DATE,
  "start_date" DATE NOT NULL DEFAULT now(),
  "color" VARCHAR,
  "code" VARCHAR NOT NULL UNIQUE,
  "weight" VARCHAR NOT NULL,
  "height" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
  "deleted_at" TIMESTAMP
);

CREATE TABLE "student_training_sheets" (
  "student_id" INT,
  "training_sheet_id" INT,
  PRIMARY KEY ("student_id", "training_sheet_id")
);

CREATE TABLE "training_sheets" (
  "id" SERIAL PRIMARY KEY,
  "uid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" VARCHAR NOT NULL,
  "disabled" BOOLEAN,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
  "deleted_at" TIMESTAMP
);

CREATE TABLE "training_sheets_workouts" (
  "training_sheet_id" INT,
  "workout_id" INT,
  PRIMARY KEY ("training_sheet_id", "workout_id")
);

CREATE TABLE "workouts" (
  "id" SERIAL PRIMARY KEY,
  "uid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "exercise" VARCHAR NOT NULL,
  "equipment" INT,
  "description" VARCHAR,
  "repetitions" VARCHAR,
  "series" INT,
  "rest_time" INT,
  "weight" float,
  "cadence" VARCHAR,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
  "deleted_at" TIMESTAMP
);

CREATE TABLE "equipment" (
  "id" SERIAL PRIMARY KEY,
  "uid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" VARCHAR
);

ALTER TABLE "user_student" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_student" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "student_training_sheets" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "student_training_sheets" ADD FOREIGN KEY ("training_sheet_id") REFERENCES "training_sheets" ("id");

ALTER TABLE "training_sheets_workouts" ADD FOREIGN KEY ("training_sheet_id") REFERENCES "training_sheets" ("id");

ALTER TABLE "training_sheets_workouts" ADD FOREIGN KEY ("workout_id") REFERENCES "workouts" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
