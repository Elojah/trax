-- +goose Up
-- +goose StatementBegin
CREATE TYPE "user"."resource" AS ENUM('user', 'asset', 'operation');
CREATE TYPE "user"."command" AS ENUM('user', 'asset', 'operation');
CREATE TABLE IF NOT EXISTS "user"."permission" (
	"id" UUID NOT NULL,
	"resource" "user"."resource" NOT NULL,
	"command" "user"."command" NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("id"),
	UNIQUE ("resource", "command")
);
CREATE TABLE IF NOT EXISTS "user"."role" (
	"id" UUID NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"entity_id" UUID NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("id"),
	UNIQUE ("name")
);
CREATE TABLE IF NOT EXISTS "user"."role_permission" (
	"role_id" UUID NOT NULL,
	"permission_id" UUID NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("role_id", "permission_id"),
	CONSTRAINT "fk_role_permission_role_id" FOREIGN KEY ("role_id") REFERENCES "user"."role" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT "fk_role_permission_permission_id" FOREIGN KEY ("permission_id") REFERENCES "user"."permission" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user"."role_permission";
DROP TABLE IF EXISTS "user"."role";
DROP TABLE IF EXISTS "user"."permission";
DROP TYPE IF EXISTS "user"."command";
DROP TYPE IF EXISTS "user"."resource";
-- +goose StatementEnd
