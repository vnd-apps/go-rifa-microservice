BEGIN;

  ALTER TABLE "raffle_numbers" DROP COLUMN "reserved_user_id"
  ALTER TABLE "raffle_numbers" DROP COLUMN "reserved_at"

COMMIT;
