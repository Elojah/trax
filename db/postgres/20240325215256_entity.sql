-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"."entity" (
	"id" UUID NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"avatar_url" VARCHAR(255),
	"description" TEXT,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX idx_entity_name ON "user"."entity" ("name");
CREATE INDEX idx_entity_created_at ON "user"."entity" ("created_at");
CREATE INDEX idx_entity_updated_at ON "user"."entity" ("updated_at");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user"."entity";
-- +goose StatementEnd
