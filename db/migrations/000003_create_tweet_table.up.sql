CREATE TABLE IF NOT EXISTS tweet (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(100) NOT NULL UNIQUE,
    content TEXT,
    img_url TEXT,
    reply_to_id VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT tweet_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT tweet_reply_id_fk FOREIGN KEY (reply_to_id) REFERENCES tweet (id)
);
