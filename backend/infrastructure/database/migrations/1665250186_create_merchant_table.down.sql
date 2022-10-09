DROP FUNCTION IF EXISTS trigger_set_timestamp CASCADE;
DROP TABLE IF EXISTS merchants CASCADE;
DROP TRIGGER IF EXISTS update_merchants_updated_at ON merchants CASCADE;