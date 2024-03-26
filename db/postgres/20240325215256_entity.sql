-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"."entity" (
	"id" UUID NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"avatar_url" VARCHAR(255),
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	PRIMARY KEY ("id")
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user"."entity";
-- +goose StatementEnd
