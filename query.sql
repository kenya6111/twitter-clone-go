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


-- name: CreateTweet :one
INSERT INTO tweet (
  user_id, content, img_url, reply_to_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetTweet :one
SELECT * FROM tweet
WHERE id = $1;

-- name: ListUserTweets :many
SELECT * FROM tweet
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListTimeline :many
SELECT t.*
FROM tweet t
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListReplies :many
SELECT * FROM tweet
WHERE reply_to_id = $1
ORDER BY created_at ASC;

-- name: UpdateTweetContent :one
UPDATE tweet
SET content = $2,
    img_url  = $3
WHERE id = $1
  AND user_id = $4
RETURNING *;

-- name: DeleteTweet :exec
DELETE FROM tweet
WHERE id = $1
  AND user_id = $2;

-- name: CountUserTweets :one
SELECT count(*) FROM tweet
WHERE user_id = $1;