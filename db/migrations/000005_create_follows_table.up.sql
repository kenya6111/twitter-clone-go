CREATE TABLE IF NOT EXISTS follows (
    follower_id TEXT NOT NULL,
    followed_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_id, followed_id)
    -- FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    -- FOREIGN KEY (followed_id) REFERENCES users(id) ON DELETE CASCADE
);