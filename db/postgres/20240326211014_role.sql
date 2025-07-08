-- +goose Up
-- +goose StatementBegin
CREATE TYPE "user"."resource" AS ENUM(
	'R_asset',
	'R_group',
	'R_operation',
	'R_role',
	'R_user'
);
CREATE TYPE "user"."command" AS ENUM(
	'C_read',
	'C_edit',
	'C_write'
);
CREATE TABLE IF NOT EXISTS "user"."role" (
	"id" UUID NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"group_id" UUID NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("id"),
	UNIQUE ("group_id", "name"),
	CONSTRAINT "fk_role_group_id" FOREIGN KEY ("group_id") REFERENCES "user"."group" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "user"."permission" (
	"role_id" UUID NOT NULL,
	"resource" "user"."resource" NOT NULL,
	"command" "user"."command" NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("role_id", "resource", "command"),
	CONSTRAINT "fk_permission_role_id" FOREIGN KEY ("role_id") REFERENCES "user"."role" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX IF NOT EXISTS "idx_permission_role_id" ON "user"."permission" ("role_id");
CREATE TABLE IF NOT EXISTS "user"."role_user" (
	"role_id" UUID NOT NULL,
	"user_id" UUID NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("role_id", "user_id"),
	CONSTRAINT "fk_role_user_role_id" FOREIGN KEY ("role_id") REFERENCES "user"."role" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT "fk_role_user_user_id" FOREIGN KEY ("user_id") REFERENCES "user"."user" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user"."role_user";
DROP TABLE IF EXISTS "user"."permission";
DROP TABLE IF EXISTS "user"."role";
DROP TYPE IF EXISTS "user"."command";
DROP TYPE IF EXISTS "user"."resource";
-- +goose StatementEnd
