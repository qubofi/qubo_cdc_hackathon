DROP FUNCTION IF EXISTS trigger_set_timestamp CASCADE;
DROP TABLE IF EXISTS subscriptions CASCADE;
DROP TRIGGER IF EXISTS update_subscriptions_updated_at ON subscriptions CASCADE;