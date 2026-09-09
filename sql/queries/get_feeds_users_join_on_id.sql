-- name: GetFeedsWithUser :many
SELECT users.name AS user_name, 
feeds.name AS feed_name, 
feeds.url as feed_url 
FROM feeds 
INNER JOIN users 
ON feeds.user_id = users.id;

