CREATE TYPE "TransferStatus" AS ENUM (
  'created',
  'processed',
  'failure',
  'success'
);

CREATE TABLE "accounts" (
  "id" uuid PRIMARY KEY NOT NULL,
  "owner" text NOT NULL,
  "balance" float NOT NULL DEFAULT 0 CHECK balance >= 0,
  "currency" text NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "entries" (
  "id" uuid PRIMARY KEY NOT NULL,
  "account_id" uuid NOT NULL,
  "transfer_id" uuid NOT NULL,
  "amount" float NOT NULL CHECK amount <> 0,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "transfers" (
  "id" uuid PRIMARY KEY NOT NULL,
  "from_acc" uuid NOT NULL,
  "to_acc" uuid NOT NULL,
  "amount" float NOT NULL CHECK amount > 0,
  "status" TransferStatus NOT NULL DEFAULT ('created'),
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE INDEX "owner_idx" ON "accounts" ("owner");

CREATE INDEX "accs_trns_idx" ON "entries" ("account_id", "transfer_id");

CREATE INDEX "from_to_idx" ON "transfers" ("from_acc", "to_acc");

CREATE INDEX "status_idx" ON "transfers" ("status");

COMMENT ON COLUMN "entries"."amount" IS 'Do not allow 0 values on tranasfer either pos or neg.';

COMMENT ON COLUMN "transfers"."amount" IS 'Do not allow 0 values on transfer, must be a strict pos';

ALTER TABLE "accounts" ADD FOREIGN KEY ("id") REFERENCES "transfers" ("from_acc");

ALTER TABLE "accounts" ADD FOREIGN KEY ("id") REFERENCES "transfers" ("to_acc");

ALTER TABLE "accounts" ADD FOREIGN KEY ("id") REFERENCES "entries" ("account_id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("id") REFERENCES "entries" ("transfer_id");
