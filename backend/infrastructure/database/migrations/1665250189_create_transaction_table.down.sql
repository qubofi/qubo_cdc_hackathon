DROP FUNCTION IF EXISTS trigger_set_timestamp CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TRIGGER IF EXISTS update_transactions_updated_at ON transactions CASCADE;