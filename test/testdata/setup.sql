-- setup.sql

-- Create habits table if it doesn't exist
CREATE TABLE IF NOT EXISTS habits (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    frequency VARCHAR(10) NOT NULL,
    current_streak INT DEFAULT 0,
    last_completed_at TIMESTAMP NULL,
    total_completions INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create or replace the timestamp trigger function
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_proc WHERE proname = 'update_timestamp'
    ) THEN
        CREATE FUNCTION update_timestamp()
        RETURNS TRIGGER AS $$
        BEGIN
            NEW.updated_at = CURRENT_TIMESTAMP;
            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;
    END IF;
END$$;

-- Create the trigger if it doesn't already exist
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'set_timestamp'
    ) THEN
        CREATE TRIGGER set_timestamp
        BEFORE UPDATE ON habits
        FOR EACH ROW
        EXECUTE FUNCTION update_timestamp();
    END IF;
END$$;

-- Insert seed habit only if not already inserted
INSERT INTO habits (name, frequency, current_streak, last_completed_at, total_completions)
SELECT 'Test Habit', 'daily', 5, NOW(), 10
WHERE NOT EXISTS (
    SELECT 1 FROM habits WHERE name = 'Test Habit'
);
