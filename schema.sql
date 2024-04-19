CREATE TABLE IF NOT EXISTS data_instance (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    value jsonb NOT NULL
)