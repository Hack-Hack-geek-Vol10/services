CREATE TYPE "auth" AS ENUM (
  'read_only',
  'read_write',
  'owner'
);

CREATE TABLE IF NOT EXISTS "project_members" (
  "project_id" varchar NOT NULL,
  "user_id" varchar NOT NULL,
  "authority" auth NOT NULL
);

ALTER TABLE "project_members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
ALTER TABLE "project_members" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("project_id");
