CREATE TABLE IF NOT EXISTS notices (
    id SERIAL PRIMARY KEY,
    tweet_id INTEGER,
    notice_type SMALLINT NOT NULL,
    sender_id TEXT NOT NULL,
    receiver_id TEXT NOT NULL,
    is_read boolean DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- FOREIGN KEY (tweet_id) REFERENCES tweets(id) ON DELETE SET NULL,
    -- FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    -- FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
);