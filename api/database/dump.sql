CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "is_personal" boolean DEFAULT false,
  "photo" varchar,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp
);

CREATE TABLE "user_student" (
  "user_id" int,
  "student_id" int,
  PRIMARY KEY ("user_id", "student_id")
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

CREATE TABLE "student_training_sheets" (
  "student_id" int,
  "training_sheet_id" int,
  PRIMARY KEY ("student_id", "training_sheet_id")
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

CREATE TABLE "training_sheets_workouts" (
  "training_sheet_id" int,
  "workout_id" int,
  PRIMARY KEY ("training_sheet_id", "workout_id")
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

ALTER TABLE "user_student" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_student" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "student_training_sheets" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "student_training_sheets" ADD FOREIGN KEY ("training_sheet_id") REFERENCES "training_sheets" ("id");

ALTER TABLE "training_sheets_workouts" ADD FOREIGN KEY ("training_sheet_id") REFERENCES "training_sheets" ("id");

ALTER TABLE "training_sheets_workouts" ADD FOREIGN KEY ("workout_id") REFERENCES "workouts" ("id");

ALTER TABLE "students" ADD FOREIGN KEY ("training_sheets_code") REFERENCES "training_sheets" ("id");

ALTER TABLE "training_sheets" ADD FOREIGN KEY ("workouts_code") REFERENCES "workouts" ("id");

ALTER TABLE "equipment" ADD FOREIGN KEY ("id") REFERENCES "workouts" ("equipment");
