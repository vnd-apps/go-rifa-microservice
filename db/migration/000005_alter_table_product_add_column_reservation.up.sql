BEGIN;

ALTER TABLE "raffle_numbers" ADD COLUMN "reserved_user_id" varchar(255);
ALTER TABLE "raffle_numbers" ADD COLUMN "reserved_at" DATE;

COMMIT;
