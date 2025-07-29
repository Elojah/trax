-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"."invitation" (
	"id" UUID NOT NULL,
	"email" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("id"),
	UNIQUE ("email")
);
CREATE INDEX IF NOT EXISTS "idx_invitation_email" ON "user"."invitation" ("email");
CREATE INDEX IF NOT EXISTS "idx_invitation_created_at" ON "user"."invitation" ("created_at");

CREATE TABLE IF NOT EXISTS "user"."invitation_role" (
	"invitation_id" UUID NOT NULL,
	"role_id" UUID NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("invitation_id", "role_id"),
	CONSTRAINT "fk_invitation_role_invitation_id" FOREIGN KEY ("invitation_id") REFERENCES "user"."invitation" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT "fk_invitation_role_role_id" FOREIGN KEY ("role_id") REFERENCES "user"."role" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX IF NOT EXISTS "idx_invitation_role_invitation_id" ON "user"."invitation_role" ("invitation_id");
CREATE INDEX IF NOT EXISTS "idx_invitation_role_role_id" ON "user"."invitation_role" ("role_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user"."invitation_role";
DROP TABLE IF EXISTS "user"."invitation";
-- +goose StatementEnd
