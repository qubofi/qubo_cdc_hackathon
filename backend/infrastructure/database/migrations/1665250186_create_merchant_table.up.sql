CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE merchants (
   id BIGSERIAL primary key,
   merchant_name TEXT,
   wallet_address TEXT,
   email TEXT,
   created_at TIMESTAMP default now(),
   updated_at TIMESTAMP default now()
);

CREATE TRIGGER update_merchants_updated_at
BEFORE UPDATE ON merchants
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();