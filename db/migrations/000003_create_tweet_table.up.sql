CREATE TABLE IF NOT EXISTS tweet (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(100),
    content TEXT,
    img_url TEXT,
    reply_to_id VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
