-- +goose Up
-- +goose StatementBegin
CREATE TYPE "user"."resource" AS ENUM('R_user', 'R_asset', 'R_operation');
CREATE TYPE "user"."command" AS ENUM('C_read', 'C_create', 'C_update', 'C_delete');
CREATE TABLE IF NOT EXISTS "user"."role" (
	"id" UUID NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"entity_id" UUID NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("id"),
	UNIQUE ("entity_id", "name"),
	CONSTRAINT "fk_role_entity_id" FOREIGN KEY ("entity_id") REFERENCES "user"."entity" ("id") ON DELETE CASCADE ON UPDATE CASCADE
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
