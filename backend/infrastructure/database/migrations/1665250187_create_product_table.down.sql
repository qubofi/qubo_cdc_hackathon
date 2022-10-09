DROP FUNCTION IF EXISTS trigger_set_timestamp CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TRIGGER IF EXISTS update_products_updated_at ON products CASCADE;