INSERT INTO users (id, name, email, password, is_active)
VALUES
  ('1', 'user1', 'user1@example.com', 'hashed_PW1!', TRUE),
  ('2', 'user2', 'user2@example.com', 'hashed_PW2!', FALSE);


INSERT INTO email_verify_token (user_id, token, expires_at, created_at)
VALUES
  ('1', 'abc123xyztoken', NOW() + INTERVAL '24 hours', NOW()),
  ('2', 'zzz555tokentest', NOW() + INTERVAL '24 hours', NOW());
