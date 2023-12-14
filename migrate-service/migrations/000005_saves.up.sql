CREATE TABLE "saves" (
  "save_id" varchar PRIMARY KEY,
  "project_id" varchar NOT NULL,
  "editor" varchar NOT NULL,
  "object" bytea NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);
ALTER TABLE "saves"
ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("project_id");