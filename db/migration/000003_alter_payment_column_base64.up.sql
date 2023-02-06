BEGIN;

  ALTER TABLE pix_payments RENAME COLUMN qr_code_base_64 TO qr_code_base64;

COMMIT;
