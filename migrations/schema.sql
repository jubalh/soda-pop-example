CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"firstname" text NOT NULL,
"lastname" text NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
, "age" integer NOT NULL DEFAULT '23');
