CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE subscriptions (
   id BIGSERIAL primary key,
   product_id BIGINT references products (id),
   customer_wallet_address TEXT,
   started_date TIMESTAMP default now(),
   last_paid_date TIMESTAMP,
   created_at TIMESTAMP default now(),
   updated_at TIMESTAMP default now(),
   cancelled_at TIMESTAMP
);

CREATE TRIGGER update_subscriptions_updated_at
BEFORE UPDATE ON subscriptions
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
