CREATE TABLE IF NOT EXISTS users (
    id  VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    bio TEXT,
    header_image_url TEXT,
    profile_image_url TEXT,
    is_active boolean NOT NULL default FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);