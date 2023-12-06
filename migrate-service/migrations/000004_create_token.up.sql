CREATE TABLE "tokens" (
  "token_id" VARCHAR PRIMARY KEY,
  "project_id" varchar NOT NULL,
  "authority" auth NOT NUll,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);
ALTER TABLE "tokens"
ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("project_id");