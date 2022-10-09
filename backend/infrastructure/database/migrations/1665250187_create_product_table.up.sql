CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE products (
   id BIGSERIAL primary key,
   merchant_id BIGINT references merchants (id),
   product_name TEXT,
   cost REAL,
   currency TEXT,
   interval INTEGER,
   interval_unit TEXT,
   created_at TIMESTAMP default now(),
   updated_at TIMESTAMP default now()
);

CREATE TRIGGER update_products_updated_at
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
