BEGIN;
CREATE TABLE IF NOT EXISTS raffles (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  description text NOT NULL,
  slug varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  image_url varchar(255),
  unit_price numeric(10,2) NOT NULL,
  user_limit integer NOT NULL,
  quantity integer NOT NULL,
  prize_draw_number integer NOT NULL,
  UNIQUE(slug)
);

CREATE TABLE IF NOT EXISTS raffle_numbers
(
  id serial PRIMARY KEY,
  raffle_slug varchar(255) REFERENCES raffles(slug),
  number integer NOT NULL,
  status varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "orders"
(
  id serial PRIMARY KEY,
  raffle_slug varchar(255) REFERENCES raffles(slug),
  user_id varchar(255) NOT NULL,
  total numeric(10,2) NOT NULL,
  payment_method integer NOT NULL,
  pix_code varchar(255),
  status varchar(255) NOT NULL
);

CREATE TABLE order_items (
  id serial PRIMARY KEY,
  order_id integer REFERENCES "orders"(id),
  number integer NOT NULL
);

CREATE INDEX raffle_slug_idx ON raffles (slug);
CREATE INDEX orders_raffle_slug_idx ON "orders" (raffle_slug);
CREATE INDEX orders_user_id_idx ON "orders" (user_id);

COMMIT;
