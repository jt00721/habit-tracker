-- teardown.sql

-- Drop trigger if it exists
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'set_timestamp'
    ) THEN
        DROP TRIGGER set_timestamp ON habits;
    END IF;
END$$;

-- Drop function if it exists
DROP FUNCTION IF EXISTS update_timestamp() CASCADE;

-- Drop table if it exists
DROP TABLE IF EXISTS habits CASCADE;

