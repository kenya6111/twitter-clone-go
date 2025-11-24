CREATE TABLE IF NOT EXISTS retweets (
    tweet_id INTEGER NOT NULL,
    user_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (tweet_id, user_id)
    -- FOREIGN KEY (tweet_id) REFERENCES tweets(id) ON DELETE CASCADE,
    -- FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);