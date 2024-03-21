-- +goose Up
-- +goose StatementBegin
-- User / Entity / Permission
CREATE SCHEMA IF NOT EXISTS "user";
-- User
CREATE TABLE IF NOT EXISTS "user"."user" (
	"id" UUID NOT NULL,
	"email" VARCHAR(255) NOT NULL,
	"password_hash" BYTEA NOT NULL,
	"password_salt" BYTEA NOT NULL,
	"google_id" VARCHAR(255),
	"twitch_id" VARCHAR(255),
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("id"),
	UNIQUE ("email"),
	UNIQUE ("google_id"),
	UNIQUE ("twitch_id")
);
CREATE TABLE IF NOT EXISTS "user"."user_profile" (
	"user_id" UUID NOT NULL,
	"first_name" VARCHAR(255) NOT NULL,
	"last_name" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("user_id"),
	CONSTRAINT "fk_user_profile_user_id" FOREIGN KEY ("user_id") REFERENCES "user"."user" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user"."user_profile";
DROP TABLE IF EXISTS "user"."user";
DROP SCHEMA "user";
-- +goose StatementEnd
