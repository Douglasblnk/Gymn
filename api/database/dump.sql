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
	"deleted_at" TIMESTAMP,
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE TABLE "students" (
  "id" SERIAL PRIMARY KEY,
  "uid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" INT NOT NULL,
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
  "deleted_at" TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE "training_sheets" (
  "id" SERIAL PRIMARY KEY,
  "uid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" VARCHAR NOT NULL,
  "active" BOOLEAN DEFAULT true,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
  "deleted_at" TIMESTAMP
);

CREATE TABLE "student_training_sheets" (
  "student_id" INT,
  "training_sheet_id" INT,
  PRIMARY KEY ("student_id", "training_sheet_id"),
  FOREIGN KEY ("student_id") REFERENCES "students" ("id") ON UPDATE CASCADE,
  FOREIGN KEY ("training_sheet_id") REFERENCES "training_sheets" ("id") ON UPDATE CASCADE
);

CREATE TABLE "equipment" (
  "id" SERIAL PRIMARY KEY,
  "uid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" VARCHAR
);

CREATE TABLE "workouts" (
  "id" SERIAL PRIMARY KEY,
  "uid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "exercise" VARCHAR NOT NULL,
  "equipment_id" INT,
  "description" VARCHAR,
  "repetitions" VARCHAR,
  "series" INT,
  "rest_time" INT,
  "weight" float,
  "cadence" VARCHAR,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
  "deleted_at" TIMESTAMP,
  FOREIGN KEY ("equipment_id") REFERENCES "equipment" ("id") ON UPDATE CASCADE
);

CREATE TABLE "training_sheets_workouts" (
  "training_sheet_id" INT,
  "workout_id" INT,
  PRIMARY KEY ("training_sheet_id", "workout_id"),
  FOREIGN KEY ("training_sheet_id") REFERENCES "training_sheets" ("id") ON UPDATE CASCADE,
  FOREIGN KEY ("workout_id") REFERENCES "workouts" ("id") ON UPDATE CASCADE
);
