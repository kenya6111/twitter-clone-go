-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CountUsersByEmail :one
SELECT count(*) FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  name,email,password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  set is_active = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateEmailVerifyToken :one
INSERT INTO email_verify_token (
  user_id,token,expires_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetEmailVerifyToken :one
SELECT * FROM email_verify_token
WHERE token = $1
AND user_id = $2;

-- name: DeleteEmailVerifyToken :exec
DELETE FROM email_verify_token
WHERE  token = $1;