-- name: GetFeedFollowsForUser :many
SELECT feeds.name, users.name
FROM feed_follows
INNER JOIN users ON feed_follows.user_id = users.id
RIGHT JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE users.name = $1;