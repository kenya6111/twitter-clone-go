CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

INSERT INTO users (name, email,password)
VALUES
('Alice', 'alice@example.com', 'password'),
('Bob', 'bob@example.com','password'),
('Charlie', 'charlie@example.com','password');
