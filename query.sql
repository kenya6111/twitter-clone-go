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

-- name: UpdateUser :one
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
  user_id, content, reply_to_id
) VALUES (
  $1, $2, $3
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
SET content = $2
WHERE id = $1
  AND user_id = $3
RETURNING *;

-- name: DeleteTweet :exec
DELETE FROM tweet
WHERE id = $1
  AND user_id = $2;

-- name: CountUserTweets :one
SELECT count(*) FROM tweet
WHERE user_id = $1;

-- name: CreateTweetImage :one
INSERT INTO tweet_images (
  tweet_id, image_url
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListTweetImages :many
SELECT * FROM tweet_images
WHERE tweet_id = $1
ORDER BY created_at ASC;

-- name: DeleteTweetImagesByTweet :exec
DELETE FROM tweet_images
WHERE tweet_id = $1;

-- name: DeleteTweetImage :exec
DELETE FROM tweet_images
WHERE id = $1;



-- =======================
-- follows
-- =======================

-- name: CreateFollow :exec
INSERT INTO follows (
  follower_id, followed_id
) VALUES (
  $1, $2
);

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE follower_id = $1
  AND followed_id = $2;

-- name: ListFollowers :many
SELECT follower_id, followed_id, created_at
FROM follows
WHERE followed_id = $1
ORDER BY created_at DESC;

-- name: ListFollowings :many
SELECT follower_id, followed_id, created_at
FROM follows
WHERE follower_id = $1
ORDER BY created_at DESC;

-- name: CountFollowers :one
SELECT count(*) FROM follows
WHERE followed_id = $1;

-- =======================
-- likes
-- =======================

-- name: CreateLike :exec
INSERT INTO likes (
  tweet_id, user_id
) VALUES (
  $1, $2
);

-- name: DeleteLike :exec
DELETE FROM likes
WHERE tweet_id = $1
  AND user_id = $2;

-- name: CountTweetLikes :one
SELECT count(*) FROM likes
WHERE tweet_id = $1;

-- name: ListTweetLikes :many
SELECT tweet_id, user_id, created_at
FROM likes
WHERE tweet_id = $1
ORDER BY created_at DESC;

-- =======================
-- retweets
-- =======================
-- name: CreateRetweet :exec
INSERT INTO retweets (
  tweet_id, user_id
) VALUES (
  $1, $2
);

-- name: DeleteRetweet :exec
DELETE FROM retweets
WHERE tweet_id = $1
  AND user_id = $2;

-- name: CountTweetRetweets :one
SELECT count(*) FROM retweets
WHERE tweet_id = $1;

-- name: ListTweetRetweets :many
SELECT tweet_id, user_id, created_at
FROM retweets
WHERE tweet_id = $1
ORDER BY created_at DESC;

-- =======================
-- notices
-- =======================

-- name: CreateNotice :one
INSERT INTO notices (
  tweet_id, notice_type, sender_id, receiver_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: ListUserNotices :many
SELECT * FROM notices
WHERE receiver_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountUnreadNotices :one
SELECT count(*) FROM notices
WHERE receiver_id = $1
  AND is_read = FALSE;

-- name: MarkNoticeRead :exec
UPDATE notices
SET is_read = TRUE
WHERE id = $1;

-- name: DeleteNotice :exec
DELETE FROM notices
WHERE id = $1;

-- =======================
-- rooms
-- =======================

-- name: CreateRoom :one
INSERT INTO rooms (
  user1_id, user2_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetRoomByUsers :one
SELECT * FROM rooms
WHERE (user1_id = $1 AND user2_id = $2)
   OR (user1_id = $2 AND user2_id = $1)
LIMIT 1;

-- name: ListUserRooms :many
SELECT * FROM rooms
WHERE user1_id = $1 OR user2_id = $1
ORDER BY created_at DESC;

-- =======================
-- messages
-- =======================

-- name: CreateMessage :one
INSERT INTO messages (
  room_id, sender_id, receiver_id, sentence, image_url
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: ListRoomMessages :many
SELECT * FROM messages
WHERE room_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = $1;

-- =======================
-- message_images
-- =======================

-- name: CreateMessageImage :one
INSERT INTO message_images (
  message_id, image_url
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListMessageImages :many
SELECT * FROM message_images
WHERE message_id = $1
ORDER BY created_at ASC;

-- name: DeleteMessageImagesByMessage :exec
DELETE FROM message_images
WHERE message_id = $1;