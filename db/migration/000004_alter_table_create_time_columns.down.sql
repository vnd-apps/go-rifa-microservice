BEGIN;

  ALTER TABLE "orders" DROP COLUMN "created_at";
  ALTER TABLE "orders" DROP COLUMN "updated_at";
  ALTER TABLE "orders" DROP COLUMN "deleted_at";

  ALTER TABLE "order_items" DROP COLUMN "created_at";
  ALTER TABLE "order_items" DROP COLUMN "updated_at";
  ALTER TABLE "order_items" DROP COLUMN "deleted_at";

  ALTER TABLE "pix_payments" DROP COLUMN "created_at";
  ALTER TABLE "pix_payments" DROP COLUMN "updated_at";
  ALTER TABLE "pix_payments" DROP COLUMN "deleted_at";

COMMIT;
