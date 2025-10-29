-- Create agent table
CREATE TABLE IF NOT EXISTS agents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    subscription_id UUID NOT NULL,
    events_count BIGINT DEFAULT 0 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- -- Add foreign key constraint to subscription
-- ALTER TABLE agents 
-- ADD CONSTRAINT fk_agents_subscription_id 
-- FOREIGN KEY (subscription_id) REFERENCES subscriptions(id) ON DELETE CASCADE;

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_agents_subscription_id ON agents(subscription_id);
CREATE INDEX IF NOT EXISTS idx_agents_created_at ON agents(created_at);
