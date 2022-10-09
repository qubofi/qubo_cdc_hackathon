CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE users (
   id BIGSERIAL primary key,
   first_name TEXT,
   last_name TEXT,
   email TEXT,
   password_hash TEXT,
   created_at TIMESTAMP default now(),
   updated_at TIMESTAMP default now()
);

CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();