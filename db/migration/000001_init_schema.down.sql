BEGIN;

  DROP INDEX orders_raffle_slug_idx;
  DROP INDEX orders_user_id_idx;
  DROP TABLE order_items;
  DROP TABLE "orders";
  DROP TABLE raffle_numbers;
  DROP TABLE raffles;

  COMMIT;
