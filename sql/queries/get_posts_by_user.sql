-- name: GePostsForUser :many
SELECT posts.*
FROM posts
INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
LIMIT $2;