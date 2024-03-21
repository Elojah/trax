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
-- Entity
CREATE TABLE IF NOT EXISTS "user"."entity" (
	"id" UUID NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("id")
);
CREATE TABLE IF NOT EXISTS "user"."entity_profile" (
	"entity_id" UUID NOT NULL,
	"description" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("entity_id"),
	CONSTRAINT "fk_entity_profile_entity_id" FOREIGN KEY ("entity_id") REFERENCES "user"."entity" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
-- Permission
CREATE TYPE "user"."permission_resource_type" AS ENUM ('user', 'entity', 'asset');
CREATE TABLE IF NOT EXISTS "user"."permission_resource" (
	-- Refers to external resource
	"id" UUID NOT NULL,
	-- External table name
	"type" "user"."permission_resource_type" NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("id")
);
CREATE TABLE IF NOT EXISTS "user"."permission_scope" (
	"id" UUID NOT NULL,
	-- Permission scope are always related to an entity
	"entity_id" UUID NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("id"),
	CONSTRAINT "fk_permission_scope_entity_id" FOREIGN KEY ("entity_id") REFERENCES "user"."entity" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "user"."permission" (
	"permission_scope_id" UUID NOT NULL,
	"permission_resource_id" UUID NOT NULL,
	"write" BOOLEAN NOT NULL,
	"read" BOOLEAN NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("permission_scope_id", "permission_resource_id"),
	CONSTRAINT "fk_permission_permission_scope_id" FOREIGN KEY ("permission_scope_id") REFERENCES "user"."permission_scope" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT "fk_permission_permission_resource_id" FOREIGN KEY ("permission_resource_id") REFERENCES "user"."permission_resource" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "user"."permission_user" (
	"user_id" UUID NOT NULL,
	"permission_scope_id" UUID NOT NULL,
	"created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	"updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	PRIMARY KEY ("user_id", "permission_scope_id"),
	CONSTRAINT "fk_permission_user_id" FOREIGN KEY ("user_id") REFERENCES "user"."user" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT "fk_permission_permission_scope_id" FOREIGN KEY ("permission_scope_id") REFERENCES "user"."permission_scope" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
