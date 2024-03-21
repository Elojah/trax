CREATE SCHEMA IF NOT EXISTS "asset";
-- Commodity & Asset
CREATE TABLE IF NOT EXISTS "asset"."commodity" (
	"id" UUID NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	PRIMARY KEY ("id")
);
CREATE TYPE "asset"."unit" AS ENUM('dollar', 'euro', 'pound');
CREATE DOMAIN "asset"."amount_value" AS DECIMAL(10, 2) CHECK(VALUE IS NOT NULL);
CREATE DOMAIN "asset"."amount_unit" AS "asset"."unit" CHECK(VALUE IS NOT NULL);
CREATE TYPE "asset"."amount" AS (
	"value" "asset"."amount_value",
	"unit" "asset"."amount_unit"
);
CREATE TABLE IF NOT EXISTS "asset"."asset" (
	"id" UUID NOT NULL,
	"commodity_id" UUID NOT NULL,
	"amount" "asset"."amount" NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	PRIMARY KEY ("id"),
	CONSTRAINT "fk_asset_commodity_id" FOREIGN KEY ("commodity_id") REFERENCES "asset"."commodity" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
-- Transaction
CREATE TABLE IF NOT EXISTS "asset"."transaction" (
	"id" UUID NOT NULL,
	"at" TIMESTAMP NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	PRIMARY KEY ("id")
);
CREATE TYPE "asset"."status" AS ENUM('owner', 'loan', 'transport');
CREATE TABLE IF NOT EXISTS "asset"."holder" (
	"id" UUID NOT NULL,
	"asset_id" UUID NOT NULL,
	"entity_id" UUID NOT NULL,
	"status" "asset"."status" NOT NULL,
	"transaction_id_init" UUID NOT NULL,
	"transaction_id_term" UUID,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	PRIMARY KEY ("id"),
	CONSTRAINT "fk_holder_asset_id" FOREIGN KEY ("asset_id") REFERENCES "asset"."asset" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT "fk_holder_entity_id" FOREIGN KEY ("entity_id") REFERENCES "user"."entity" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT "fk_holder_transaction_id_init" FOREIGN KEY ("transaction_id_init") REFERENCES "asset"."transaction" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT "fk_holder_transaction_id_term" FOREIGN KEY ("transaction_id_term") REFERENCES "asset"."transaction" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS "asset"."process" (
	"parent_asset_id" UUID NOT NULL,
	"child_asset_id" UUID NOT NULL,
	"transaction_id" UUID NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	PRIMARY KEY ("parent_asset_id", "child_asset_id"),
	CONSTRAINT "fk_process_parent_asset_id" FOREIGN KEY ("parent_asset_id") REFERENCES "asset"."asset" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT "fk_process_child_asset_id" FOREIGN KEY ("child_asset_id") REFERENCES "asset"."asset" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
