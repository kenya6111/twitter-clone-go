CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL,
    sender_id TEXT NOT NULL,
    receiver_id TEXT NOT NULL,
    sentence TEXT,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- FOREIGN KEY (room_id) REFERENCES dm_rooms(id) ON DELETE CASCADE,
    -- FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    -- FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
);
