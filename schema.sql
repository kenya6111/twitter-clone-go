CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    is_active boolean default FALSE
);

CREATE TABLE IF NOT EXISTS email_verify_token (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(100) NOT NULL UNIQUE,
    token VARCHAR(100) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_team_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
);
