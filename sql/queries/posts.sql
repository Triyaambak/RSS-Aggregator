-- name: CreatePost :one
INSERT INTO POSTS (id, created_at, updated_at, title,description,published_at, url, feed_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts JOIN feed_follow ON posts.feed_id = feed_follow.feed_id WHERE feed_follow.user_id = $1 ORDER BY posts.published_at DESC LIMIT $2;
