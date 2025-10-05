INSERT INTO email_verify_token (user_id, token, expires_at, created_at)
VALUES
  ('1', 'abc123xyztoken', NOW() + INTERVAL '24 hours', NOW()),
  ('2', 'zzz555tokentest', NOW() + INTERVAL '24 hours', NOW());
