-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS "user";
CREATE TABLE IF NOT EXISTS "user"."user" (
	"id" UUID NOT NULL,
	"email" VARCHAR(255) NOT NULL,
	"password_hash" BYTEA NOT NULL,
	"password_salt" BYTEA NOT NULL,
	"google_id" VARCHAR(255),
	"first_name" VARCHAR(255) NOT NULL,
	"last_name" VARCHAR(255) NOT NULL,
	"avatar_url" VARCHAR(255),
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("id"),
	UNIQUE ("email"),
	UNIQUE ("google_id")
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user"."user" CASCADE;
DROP SCHEMA IF EXISTS "user";
-- +goose StatementEnd
