BEGIN;

CREATE TABLE IF NOT EXISTS pix_payments
(
  id serial PRIMARY KEY,
  order_id integer REFERENCES "orders"(id),
  qr_code varchar(255) NOT NULL,
  qr_code_base_64 varchar(255) NOT NULL,
  status varchar(255) NOT NULL
);

ALTER TABLE "orders" DROP COLUMN pix_code;

CREATE INDEX orders_id_idx ON "pix_payments" (order_id);

COMMIT;
