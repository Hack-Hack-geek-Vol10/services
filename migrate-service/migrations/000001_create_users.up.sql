CREATE TABLE IF NOT EXISTS "users" (
  "user_id" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "is_delete" bool NOT NULL DEFAULT 'false'
);
