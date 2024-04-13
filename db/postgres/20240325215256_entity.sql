-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"."entity" (
	"id" UUID NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX idx_entity_name ON "user"."entity" ("name");
CREATE INDEX idx_entity_created_at ON "user"."entity" ("created_at");
CREATE INDEX idx_entity_updated_at ON "user"."entity" ("updated_at");
CREATE TABLE IF NOT EXISTS "user"."entity_profile" (
	"entity_id" UUID NOT NULL,
	"avatar_url" VARCHAR(255),
	"description" TEXT,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("entity_id"),
	CONSTRAINT "fk_profile_entity_id" FOREIGN KEY ("entity_id") REFERENCES "user"."entity" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user"."entity_profile";
DROP TABLE IF EXISTS "user"."entity";
-- +goose StatementEnd
