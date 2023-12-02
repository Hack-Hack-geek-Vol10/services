CREATE TABLE IF NOT EXISTS "projects" (
  "project_id" varchar PRIMARY KEY,
  "title" varchar NOT NULL,
  "last_image" varchar NOT NULL,
  "is_personal" bool NOT NULL DEFAULT 'true',
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "is_delete" bool NOT NULL DEFAULT 'false'
);