BEGIN;

DROP INDEX orders_id_idx;
DROP TABLE pix_payments;
ALTER TABLE "orders" ADD pix_code varchar(255);

COMMIT;
