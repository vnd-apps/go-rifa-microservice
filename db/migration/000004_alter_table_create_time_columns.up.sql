BEGIN;
  ALTER TABLE "orders" ADD COLUMN "created_at" DATE;
  ALTER TABLE "orders" ADD COLUMN "updated_at" DATE;
  ALTER TABLE "orders" ADD COLUMN "deleted_at" DATE;

  ALTER TABLE "order_items" ADD COLUMN "created_at" DATE;
  ALTER TABLE "order_items" ADD COLUMN "updated_at" DATE;
  ALTER TABLE "order_items" ADD COLUMN "deleted_at" DATE;

  ALTER TABLE "pix_payments" ADD COLUMN "created_at" DATE;
  ALTER TABLE "pix_payments" ADD COLUMN "updated_at" DATE;
  ALTER TABLE "pix_payments" ADD COLUMN "deleted_at" DATE;

COMMIT;
