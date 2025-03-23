CREATE TABLE habits (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    frequency VARCHAR(10) NOT NULL,
    current_streak INT DEFAULT 0,
    last_completed_at TIMESTAMP NULL,
    total_completions INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Trigger to auto-update updated_at on row update
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON habits
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

INSERT INTO habits (name, frequency, current_streak, last_completed_at, total_completions)
VALUES ('Test Habit', 'daily', 5, NOW(), 10);

