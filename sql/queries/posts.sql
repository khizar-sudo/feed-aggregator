-- name: CreatePost :one
INSERT INTO posts (
    id, 
    created_at, 
    updated_at, 
    title, 
    url, 
    description, 
    published_at, 
    feed_id
    )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
WITH user_feeds AS (
    SELECT id FROM feeds
    WHERE user_id = $1
)
SELECT posts.* FROM posts
INNER JOIN user_feeds ON posts.feed_id = user_feeds.id
ORDER BY published_at DESC
LIMIT $2;