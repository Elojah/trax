-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"."group" (
	"id" UUID NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"avatar_url" VARCHAR(255),
	"description" TEXT,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX idx_group_name ON "user"."group" ("name");
CREATE INDEX idx_group_created_at ON "user"."group" ("created_at");
CREATE INDEX idx_group_updated_at ON "user"."group" ("updated_at");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user"."group";
-- +goose StatementEnd
