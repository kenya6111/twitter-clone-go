CREATE TABLE IF NOT EXISTS message_images (
    id SERIAL PRIMARY KEY,
    message_id INTEGER NOT NULL,
    image_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE
);