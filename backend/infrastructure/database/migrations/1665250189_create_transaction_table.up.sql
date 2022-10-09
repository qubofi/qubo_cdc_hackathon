CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE transactions (
   id BIGSERIAL primary key,
   subscription_id BIGINT references subscriptions (id),
   transaction_hash TEXT,
   created_at TIMESTAMP default now(),
   updated_at TIMESTAMP default now()
);

CREATE TRIGGER update_transactions_updated_at
BEFORE UPDATE ON transactions
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
